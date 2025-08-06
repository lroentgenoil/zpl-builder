package picturesFunctions

import (
	"image"
	"image/draw"

	"github.com/disintegration/imaging"
)

func AddImageToCanvas(canvas draw.Image, img image.Image, index, filas, columnas int, marginX, marginY float64) {
	const dpi = 300.0
	pxPerMM := dpi / 25.4

	marginXPx := int(marginX * pxPerMM)
	marginYPx := int(marginY * pxPerMM)

	canvasWidth := canvas.Bounds().Dx()
	canvasHeight := canvas.Bounds().Dy()

	totalMarginX := (columnas + 1) * marginXPx
	totalMarginY := (filas + 1) * marginYPx

	cellWidth := (canvasWidth - totalMarginX) / columnas
	cellHeight := (canvasHeight - totalMarginY) / filas

	currentRow := (index / columnas) % filas
	currentCol := index % columnas

	x := marginXPx + currentCol*(cellWidth+marginXPx)
	y := marginYPx + currentRow*(cellHeight+marginYPx)

	resized := imaging.Resize(img, cellWidth, cellHeight, imaging.Lanczos)

	draw.Draw(canvas,
		image.Rect(x, y, x+cellWidth, y+cellHeight),
		resized,
		image.Point{},
		draw.Over,
	)
}
