package handle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	zebrash "github.com/lroentgenoil/zebrashMod"
	"github.com/lroentgenoil/zebrashMod/drawers"
	zebrashElements "github.com/lroentgenoil/zebrashMod/elements"
	"github.com/lroentgenoil/zpl-builder/elements"
	"github.com/lroentgenoil/zpl-builder/functions"
)

func PDFOutput(res []zebrashElements.LabelInfo, params elements.FormattedParams, opts drawers.DrawerOptions) error {
	lote := 10000
	Mnumber := params.Filas * params.Columnas
	switch Mnumber {
	case 1:
		lote = 10000
	case 3:
		lote = 10002
	case 4:
		lote = 10004
	case 5:
		lote = 10005
	case 14:
		lote = 10010
	case 44:
		lote = 10032
	}
	loteIndex := 0

	drawer := zebrash.NewDrawer()

	var result interface{}
	switch params.Output {
	case "file":
		result = []string{}
	default:
		result = [][]byte{}
	}

	pdf := functions.NewPDF(params)

	for idx, label := range res {
		buf := &bytes.Buffer{}
		err := drawer.DrawLabelAsPng(label, buf, opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Etiqueta %d falló: %v\n", idx, err)
			continue
		}

		if params.Mosaico {
			functions.AddImageToMosaic(pdf, buf, idx, params.Filas, params.Columnas, params.MarginX, params.MarginY)
		} else {
			functions.AddPage(pdf, buf, idx, params.LabelWidth, params.LabelHeight)
		}

		if (idx+1)%lote == 0 && idx != len(res)-1 {
			result, err = functions.AddOutput(pdf, result, params, loteIndex)
			if err != nil {
				return err
			}
			loteIndex++
			pdf = functions.NewPDF(params)
		}
	}

	result, err := functions.AddOutput(pdf, result, params, loteIndex)
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
