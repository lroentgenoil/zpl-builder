package picturesFunctions

import "image"

func NewMosaicCanvas(labelWidthMm float64, labelHeightMm float64, filas int, columnas int) *image.RGBA {
	labelWidthPx := int((labelWidthMm / 25.4) * float64(300))
	labelHeightPx := int((labelHeightMm / 25.4) * float64(300))

	return image.NewRGBA(image.Rect(0, 0, labelWidthPx, labelHeightPx))
}
