package jobs

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"

	zebrash "github.com/lroentgenoil/zebrashMod"
	"github.com/lroentgenoil/zebrashMod/drawers"
	zebrashElements "github.com/lroentgenoil/zebrashMod/elements"
)

type pngResult struct {
	Index int
	Buf   *bytes.Buffer
	Err   error
}

func GeneratePNGsParallel(res []zebrashElements.LabelInfo, opts drawers.DrawerOptions) ([]*bytes.Buffer, error) {
	numJobs := len(res)
	numWorkers := runtime.NumCPU()

	jobs := make(chan int, numJobs)
	results := make(chan pngResult, numJobs)
	var wg sync.WaitGroup

	drawer := zebrash.NewDrawer()

	// Workers
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for idx := range jobs {
				buf := &bytes.Buffer{}
				err := drawer.DrawLabelAsPng(res[idx], buf, opts)
				results <- pngResult{Index: idx, Buf: buf, Err: err}
			}
		}()
	}

	// Enviar trabajos
	for i := 0; i < numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	// Esperar que terminen
	go func() {
		wg.Wait()
		close(results)
	}()

	// Recibir resultados
	outputBuffers := make([]*bytes.Buffer, numJobs)
	for res := range results {
		if res.Err != nil {
			return nil, fmt.Errorf("error drawer image: %w", res.Err)
		} else {
			outputBuffers[res.Index] = res.Buf
		}
	}

	return outputBuffers, nil
}
