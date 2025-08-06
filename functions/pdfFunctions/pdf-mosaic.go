package pdfFunctions

import (
	"bytes"
	"fmt"

	"github.com/lroentgenoil/zpl-builder/elements"
	"github.com/phpdave11/gofpdf"
)

func AddImageToMosaic(pdf *gofpdf.Fpdf, buf *bytes.Buffer, index int, params elements.FormattedParams) {
	pageWidth, pageHeight := pdf.GetPageSize()
	cellWidth := pageWidth / float64(params.Columnas)
	cellHeight := pageHeight / float64(params.Filas)

	currentRow := (index / params.Columnas) % params.Filas
	currentCol := index % params.Columnas

	// Agrega una nueva página si es el primero de la hoja
	if currentRow == 0 && currentCol == 0 {
		pdf.AddPage()
	}

	imageID := fmt.Sprintf("img_%d", index)
	_ = pdf.RegisterImageOptionsReader(imageID, gofpdf.ImageOptions{ImageType: "PNG"}, buf)

	x := float64(currentCol)*cellWidth + params.MarginX
	y := float64(currentRow)*cellHeight + params.MarginY
	w := cellWidth - 2*params.MarginX
	h := cellHeight - 2*params.MarginY

	pdf.ImageOptions(imageID, x, y, w, h, false, gofpdf.ImageOptions{ImageType: "PNG"}, 0, "")
}
