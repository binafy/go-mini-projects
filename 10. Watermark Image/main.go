package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	// Open original image
	srcFile, err := os.Open("original.jpg")
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()

	src, err := jpeg.Decode(srcFile)
	if err != nil {
		panic(err)
	}

	// Open watermark
	wmFile, err := os.Open("watermark.png")
	if err != nil {
		panic(err)
	}
	defer wmFile.Close()

	watermark, err := png.Decode(wmFile)
	if err != nil {
		panic(err)
	}

	// Create a writable RGBA image
	result := image.NewRGBA(src.Bounds())

	// Copy original image
	draw.Draw(result, result.Bounds(), src, image.Point{}, draw.Src)

	// Position watermark in bottom-right corner
	margin := 20

	x := src.Bounds().Dx() - watermark.Bounds().Dx() - margin
	y := src.Bounds().Dy() - watermark.Bounds().Dy() - margin

	position := image.Rect(
		x,
		y,
		x+watermark.Bounds().Dx(),
		y+watermark.Bounds().Dy(),
	)

	// Overlay watermark
	draw.Draw(result, position, watermark, image.Point{}, draw.Over)

	// Save result
	outFile, err := os.Create("watermarked.jpg")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	err = jpeg.Encode(outFile, result, &jpeg.Options{
		Quality: 90,
	})
	if err != nil {
		panic(err)
	}
}
