package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
	"time"
)

//func createMosaicRepresentation(rectangle *Rectangle) image.Image {
//	background := createBackground(rectangle)
//	anotherRectangle := &Rectangle{60, 80, 120, 160, false}
//	drawRectangleOnBackround(background, anotherRectangle)
//
//	return background
//}

func drawRectangleOnBackround(background *image.RGBA, rectangle *Rectangle) {
	rectangleImage := image.Rect(
		rectangle.x,
		rectangle.y,
		rectangle.x+rectangle.width,
		rectangle.y+rectangle.height,
	)
	colour := color.RGBA{randomNumber(), randomNumber(), randomNumber(), randomNumber()}
	draw.Draw(background, rectangleImage, &image.Uniform{C: colour}, image.Point{}, draw.Src)
}

func createBackground(rectangle *Rectangle) *image.RGBA {
	background := image.NewRGBA(image.Rect(rectangle.x, rectangle.y, rectangle.x+rectangle.width, rectangle.y+rectangle.height))
	white := color.RGBA{R: 255, G: 255, A: 255}
	draw.Draw(background, background.Bounds(), &image.Uniform{C: white}, image.Point{}, draw.Src)
	return background
}

func saveMosaicImageAsPng(fileName string, mosaicImage image.RGBA) {
	fileHandler, err := os.Create(fileName)
	fmt.Println(fileName)
	if err != nil {
		panic(err)
	}
	defer func(fileHandler *os.File) {
		err := fileHandler.Close()
		if err != nil {

		}
	}(fileHandler)
	err = png.Encode(fileHandler, &mosaicImage)
	if err != nil {
		return
	}
}

func randomNumber() uint8 {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 255
	return uint8(rand.Intn(max-min+1) + min)
}
