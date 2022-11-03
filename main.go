package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var inputFileName, outputFileName string
	var x, y, width, height int

	flag.StringVar(&inputFileName, "input", "input/data.csv", "Input CSV file")
	flag.StringVar(&outputFileName, "output", "output/"+getDefaultOutputFileName(inputFileName), "Output file")
	flag.IntVar(&x, "x", 0, "Starting point (X)")
	flag.IntVar(&y, "y", 0, "Starting point (Y)")
	flag.IntVar(&width, "width", 1000, "Rectangle width")
	flag.IntVar(&height, "height", 500, "Rectangle height")
	flag.Parse()

	backgroundCanvas := Rectangle{x, y, width, height, false}
	dataList := readCsvData(inputFileName)

	mosaicImageInput := CreateMosaicImageInput(dataList, backgroundCanvas)
	backgroundImage := createBackground(&backgroundCanvas)
	fmt.Println(mosaicImageInput)
	for _, rectangle := range mosaicImageInput.list {
		drawRectangleOnBackround(backgroundImage, &rectangle)
	}
	saveMosaicImageAsPng(outputFileName, *backgroundImage)
}

func getDefaultOutputFileName(inputFilePath string) string {
	re := regexp.MustCompile("[a-zA-Z0-9]+\\.")
	match := re.FindStringSubmatch(inputFilePath)
	return "out_" + strings.Split(match[0], ".")[0] + ".png"
}
