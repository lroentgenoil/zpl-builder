package functions

import (
	"github.com/lroentgenoil/zpl-builder/elements"
	"github.com/phpdave11/gofpdf"
)

func NewPDF(params elements.FormattedParams) *gofpdf.Fpdf {
	if params.Mosaico {
		return gofpdf.New(params.Orientacion, "mm", params.TipoPapel, "")
	}
	return gofpdf.NewCustom(&gofpdf.InitType{
		UnitStr: "mm",
		Size: gofpdf.SizeType{
			Wd: params.LabelWidth,
			Ht: params.LabelHeight,
		},
	})
}
