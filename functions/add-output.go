package functions

import (
	"bytes"
	"fmt"

	"github.com/lroentgenoil/zpl-builder/elements"
	"github.com/phpdave11/gofpdf"
)

func AddOutput(pdf *gofpdf.Fpdf, result interface{}, params elements.FormattedParams, loteIndex int) (interface{}, error) {
	switch params.Output {
	case "file":
		fileName := fmt.Sprintf(params.UrlOutput+"lote_%05d.pdf", loteIndex)
		if err := pdf.OutputFileAndClose(fileName); err != nil {
			return result, fmt.Errorf("error guardando archivo %s: %w", fileName, err)
		}
		return append(result.([]string), fileName), nil
	default:
		buf := &bytes.Buffer{}
		if err := pdf.Output(buf); err != nil {
			return result, fmt.Errorf("error generando PDF: %w", err)
		}
		pdf.Close()
		return append(result.([][]byte), buf.Bytes()), nil
	}
}
