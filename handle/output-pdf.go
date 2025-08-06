package handle

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"runtime"

	"github.com/lroentgenoil/zebrashMod/drawers"
	zebrashElements "github.com/lroentgenoil/zebrashMod/elements"
	"github.com/lroentgenoil/zpl-builder/elements"
	"github.com/lroentgenoil/zpl-builder/functions"
	"github.com/lroentgenoil/zpl-builder/functions/pdfFunctions"
	"github.com/lroentgenoil/zpl-builder/jobs"
)

// Estructura para resultados

func PDFOutput(res []zebrashElements.LabelInfo, params elements.FormattedParams, opts drawers.DrawerOptions) error {
	loteIndex := 0
	Mnumber := params.Filas * params.Columnas
	lote := int(math.Ceil(float64(params.Chunk) / float64(Mnumber)))
	loteSize := lote * Mnumber
	pdf := pdfFunctions.NewPDF(params)
	PDFwidth, PDFheight := pdf.GetPageSize()
	var err error

	var result interface{}
	switch params.Output {
	case "file":
		result = []string{}
	default:
		result = [][]byte{}
	}

	for start := 0; start < len(res); start += loteSize {
		end := start + loteSize
		if end > len(res) {
			end = len(res)
		}

		currentLote := res[start:end]

		outputBuffers, err := jobs.GeneratePNGsParallel(currentLote, opts)
		if err != nil {
			return err
		}

		if params.Mosaico {
			if params.Comprimir {
				mosaics, err := jobs.GenerateMosaicParallel(outputBuffers, params, start, PDFwidth, PDFheight)
				if err != nil {
					return err
				}

				for idx, buf := range mosaics {
					index := start + idx

					pdfFunctions.AddPage(pdf, buf, index, PDFwidth, PDFheight)
				}
			} else {
				for i, buf := range outputBuffers {
					if buf == nil {
						continue
					}

					index := start + i

					pdfFunctions.AddImageToMosaic(pdf, buf, index, params)
				}
			}
		} else {
			if params.Comprimir {
				mosaics, err := jobs.GenerateMosaicParallel(outputBuffers, params, start, PDFwidth, PDFheight)
				if err != nil {
					return err
				}

				for idx, buf := range mosaics {
					index := start + idx

					pdfFunctions.AddPage(pdf, buf, index, PDFwidth, PDFheight)
				}
			} else {
				for idx, buf := range outputBuffers {
					if buf == nil {
						continue
					}

					index := start + idx

					pdfFunctions.AddPage(pdf, buf, index, params.LabelWidth, params.LabelHeight)
				}
			}
		}

		if end != len(res) {
			result, err = functions.AddOutput(pdf, result, params, loteIndex)
			if err != nil {
				return err
			}

			loteIndex++

			pdf.Close()
			runtime.GC()
			pdf = pdfFunctions.NewPDF(params)
		}
	}

	result, err = functions.AddOutput(pdf, result, params, loteIndex)
	if err != nil {
		return err
	}

	data, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("error generando Json: %w", err)
	}

	_, err = os.Stdout.Write(data)

	return err
}
