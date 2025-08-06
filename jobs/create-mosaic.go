package jobs

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"math"
	"runtime"
	"sync"

	"github.com/lroentgenoil/zpl-builder/elements"
	"github.com/lroentgenoil/zpl-builder/functions/picturesFunctions"
)

type mosaicJob struct {
	Index   int
	Buffers []*bytes.Buffer
}

func GenerateMosaicParallel(outputBuffers []*bytes.Buffer, params elements.FormattedParams, start int, PDFwidth, PDFheight float64) ([]*bytes.Buffer, error) {

	Mnumber := params.Filas * params.Columnas
	numMosaics := int(math.Ceil(float64(len(outputBuffers)) / float64(Mnumber)))
	numWorkers := runtime.NumCPU()

	jobs := make(chan mosaicJob, numMosaics)
	results := make(chan struct {
		Index int
		Buf   *bytes.Buffer
		Err   error
	}, numMosaics)

	var wg sync.WaitGroup

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {
				canvas := picturesFunctions.NewMosaicCanvas(PDFwidth, PDFheight, params.Filas, params.Columnas)

				for i, buf := range job.Buffers {
					if buf == nil {
						continue
					}

					img, _, err := image.Decode(buf)
					if err != nil {
						results <- struct {
							Index int
							Buf   *bytes.Buffer
							Err   error
						}{job.Index, nil, fmt.Errorf("error decoding image %d: %w", job.Index*Mnumber+i, err)}
						return
					}

					globalIndex := start + job.Index*Mnumber + i
					picturesFunctions.AddImageToCanvas(canvas, img, globalIndex, params.Filas, params.Columnas, params.MarginX, params.MarginY)
				}

				buf := &bytes.Buffer{}
				if err := png.Encode(buf, canvas); err != nil {
					results <- struct {
						Index int
						Buf   *bytes.Buffer
						Err   error
					}{job.Index, nil, fmt.Errorf("error encoding canvas: %w", err)}
					return
				}

				results <- struct {
					Index int
					Buf   *bytes.Buffer
					Err   error
				}{job.Index, buf, nil}
			}
		}()
	}

	// Enviar trabajos
	for i := 0; i < numMosaics; i++ {
		startIdx := i * Mnumber
		endIdx := startIdx + Mnumber
		if endIdx > len(outputBuffers) {
			endIdx = len(outputBuffers)
		}

		jobs <- mosaicJob{
			Index:   i,
			Buffers: outputBuffers[startIdx:endIdx],
		}
	}
	close(jobs)

	// Esperar a que terminen
	go func() {
		wg.Wait()
		close(results)
	}()

	// Recolectar resultados en orden
	mosaicBuffers := make([]*bytes.Buffer, numMosaics)
	for res := range results {
		if res.Err != nil {
			return nil, res.Err
		}
		mosaicBuffers[res.Index] = res.Buf
	}

	return mosaicBuffers, nil
}
