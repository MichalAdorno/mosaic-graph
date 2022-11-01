package main

import (
	"flag"
	"strings"
)

func main() {
	var inputFileName, outputFileName string
	var x, y, width, height int

	flag.StringVar(&inputFileName, "input", "data.csv", "Input CSV file")
	flag.StringVar(&outputFileName, "output", getDefaultOutputFileName(inputFileName), "Output file")
	flag.IntVar(&x, "x", 0, "Starting point (X)")
	flag.IntVar(&y, "y", 0, "Starting point (Y)")
	flag.IntVar(&width, "width", 1000, "Rectangle width")
	flag.IntVar(&height, "height", 500, "Rectangle height")

	flag.Parse()

	mosaicImage := createMosaicRepresentation(&Rectangle{x, y, width, height, false})
	saveMosaicImageAsPng(outputFileName, mosaicImage)
}

func getDefaultOutputFileName(inputFileName string) string {
	return "out_" + strings.Split(inputFileName, ".")[0] + ".png"
}
