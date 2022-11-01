package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type DataPair struct {
	category string
	value    float64
}

func readData(inputFileName string) []DataPair {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	dataList := buildDataList(data)
	return dataList
}

func buildDataList(data [][]string) []DataPair {
	var list []DataPair
	for _, line := range data {
		var pair DataPair
		var category string
		var value float64
		println(&line)
		for j, field := range line {
			if j == 0 {
				category = field
			} else if j == 1 {
				parsedValue, err := strconv.ParseFloat(field, 64)
				if err == nil {
					fmt.Println(err)
				}
				value = parsedValue
			}
		}
		pair = DataPair{category: category, value: value}
		list = append(list, pair)
	}
	return list
}

func printData(inputFileName string) {
	dataList := readData(inputFileName)
	for i, data := range dataList {
		fmt.Printf("[%v]=<[%v],[%v]>\n", i, data.category, data.value)
	}
}
