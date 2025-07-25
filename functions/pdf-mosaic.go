package functions

import (
	"bytes"
	"fmt"

	"github.com/phpdave11/gofpdf"
)

func AddImageToMosaic(pdf *gofpdf.Fpdf, buf *bytes.Buffer, index, filas int, columnas int, marginX float64, marginY float64) {
	pageWidth, pageHeight := pdf.GetPageSize()
	cellWidth := pageWidth / float64(columnas)
	cellHeight := pageHeight / float64(filas)

	currentRow := (index / columnas) % filas
	currentCol := index % columnas

	// Agrega una nueva página si es el primero de la hoja
	if currentRow == 0 && currentCol == 0 {
		pdf.AddPage()
	}

	imageID := fmt.Sprintf("img_%d", index)
	_ = pdf.RegisterImageOptionsReader(imageID, gofpdf.ImageOptions{ImageType: "PNG"}, buf)

	x := float64(currentCol)*cellWidth + marginX
	y := float64(currentRow)*cellHeight + marginY
	w := cellWidth - 2*marginX
	h := cellHeight - 2*marginY

	pdf.ImageOptions(imageID, x, y, w, h, false, gofpdf.ImageOptions{ImageType: "PNG"}, 0, "")
}
