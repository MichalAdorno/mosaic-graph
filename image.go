package main

import (
	"errors"
)

type Rectangle struct {
	x         int
	y         int
	width     int
	height    int
	direction bool
}

type MosaicImage struct {
	list []Rectangle
}

func Mosaic(payload []DataPair, canvas Rectangle) *MosaicImage {
	return split(payload, canvas, &MosaicImage{list: []Rectangle{}})
}

func split(payload []DataPair, canvas Rectangle, acc *MosaicImage) *MosaicImage {
	if payload == nil {
		return nil
	}
	if len(payload) == 1 {
		return &MosaicImage{list: append(acc.list, canvas)}
	}
	return nil
}

func splitIntoTwo(data []DataPair) ([]DataPair, []DataPair, error) {
	if data == nil && len(data) == 0 {
		return nil, nil, errors.New("No data")
	}
	if len(data) == 1 {
		return data, nil, nil
	}
	sum := 0.0
	for _, pair := range data {
		sum += pair.value
	}
	acc := 0.0
	var left, right []DataPair

	for i, pair := range data {
		if i == 0 || acc+pair.value <= sum/2 {
			acc += pair.value
			left = append(left, pair)
		} else {
			acc += pair.value
			right = append(right, pair)
		}
	}
	return left, right, nil
}
