package pdfFunctions

import (
	"bytes"
	"fmt"

	"github.com/phpdave11/gofpdf"
)

func AddPage(pdf *gofpdf.Fpdf, buf *bytes.Buffer, index int, labelWidth float64, labelHeight float64) {
	pdf.AddPage()
	imageID := fmt.Sprintf("label_%d.png", index)

	pdf.RegisterImageOptionsReader(imageID, gofpdf.ImageOptions{ImageType: "PNG"}, buf)
	pdf.ImageOptions(imageID, 0, 0, labelWidth, labelHeight, false, gofpdf.ImageOptions{ImageType: "PNG"}, 0, "")
}
