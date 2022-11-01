package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func createMosaicRepresentation(rectangle *Rectangle) image.Image {
	background := createBackground(rectangle)
	anotherRectangle := &Rectangle{60, 80, 120, 160, false}
	drawRectangleOnBackround(background, anotherRectangle)

	return background
}

func drawRectangleOnBackround(background *image.RGBA, rectangle *Rectangle) {
	rectangleImage := image.Rect(
		rectangle.x,
		rectangle.y,
		rectangle.x+rectangle.width,
		rectangle.y+rectangle.height,
	)
	colour := color.RGBA{200, 0, 0, 255}
	draw.Draw(background, rectangleImage, &image.Uniform{C: colour}, image.Point{}, draw.Src)
}

func createBackground(rectangle *Rectangle) *image.RGBA {
	background := image.NewRGBA(image.Rect(rectangle.x, rectangle.y, rectangle.x+rectangle.width, rectangle.y+rectangle.height))
	colour := color.RGBA{G: 100, A: 255}
	draw.Draw(background, background.Bounds(), &image.Uniform{C: colour}, image.Point{}, draw.Src)
	return background
}

func saveMosaicImageAsPng(fileName string, mosaicImage image.Image) {
	fileHandler, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer func(fileHandler *os.File) {
		err := fileHandler.Close()
		if err != nil {

		}
	}(fileHandler)
	err = png.Encode(fileHandler, mosaicImage)
	if err != nil {
		return
	}
}