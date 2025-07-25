package handle

import (
	"bytes"
	"fmt"
	"os"

	zebrash "github.com/lroentgenoil/zebrashMod"
	"github.com/lroentgenoil/zebrashMod/drawers"
	zebrashElements "github.com/lroentgenoil/zebrashMod/elements"
	"github.com/lroentgenoil/zpl-builder/elements"
	"github.com/lroentgenoil/zpl-builder/functions"
	"github.com/phpdave11/gofpdf"
)

func PDFOutput(res []zebrashElements.LabelInfo, params elements.FormattedParams, opts drawers.DrawerOptions) error {
	var pdf *gofpdf.Fpdf
	if params.Mosaico {
		pdf = gofpdf.New(params.Orientacion, "mm", params.TipoPapel, "")
	} else {
		pdf = gofpdf.NewCustom(&gofpdf.InitType{
			UnitStr: "mm",
			Size: gofpdf.SizeType{
				Wd: params.LabelWidth,
				Ht: params.LabelHeight,
			},
		})
	}

	drawer := zebrash.NewDrawer()
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
	}

	buf := &bytes.Buffer{}
	err := pdf.Output(buf)
	if err != nil {
		return fmt.Errorf("error generando PDF: %w", err)
	}

	_, err = os.Stdout.Write(buf.Bytes())
	return err
}
