package main

import (
	"flag"
	"fmt"
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

	backgroundCanvas := Rectangle{x, y, width, height, false}
	//printData(inputFileName)
	dataList := readCsvData(inputFileName)

	mosaicImageInput := CreateMosaicImageInput(dataList, backgroundCanvas)
	//_ = CreateMosaicImageInput(dataList, backgroundCanvas)
	backgroundImage := createBackground(&backgroundCanvas)
	//_ = createBackground(&backgroundCanvas)
	fmt.Println(mosaicImageInput)
	for _, rectangle := range mosaicImageInput.list {
		drawRectangleOnBackround(backgroundImage, &rectangle)
		//fmt.Println(rectangle)
	}
	saveMosaicImageAsPng(outputFileName, *backgroundImage)
	////////////

}

func getDefaultOutputFileName(inputFileName string) string {
	return "out_" + strings.Split(inputFileName, ".")[0] + ".png"
}
