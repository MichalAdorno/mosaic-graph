package main

import (
	"errors"
	"fmt"
	"math"
)

type Rectangle struct {
	x         int
	y         int
	width     int
	height    int
	direction bool
}

type MosaicImageInput struct {
	list []Rectangle
}

func CreateMosaicImageInput(payload []DataPair, canvas Rectangle) *MosaicImageInput {
	acc := MosaicImageInput{list: []Rectangle{}}
	createMosaicImageInput(payload, canvas, &acc)
	fmt.Println(acc)
	return &acc
}

func createMosaicImageInput(payload []DataPair, canvas Rectangle, acc *MosaicImageInput) {
	if payload == nil {
		return
	}
	if len(payload) == 1 {
		return
	}
	leftPayload, rightPayload, _ := splitIntoTwo(payload)

	leftSubCanvas, rightSubCanvas := divideCanvas(payload, canvas)

	if len(leftPayload) == 1 {
		newList := append(acc.list, leftSubCanvas)
		acc.list = newList
		createMosaicImageInput(leftPayload, leftSubCanvas, acc)
	} else {
		createMosaicImageInput(leftPayload, leftSubCanvas, acc)
	}

	if rightPayload != nil || len(rightPayload) > 0 {
		if len(rightPayload) == 1 {
			newList := append(acc.list, rightSubCanvas)
			acc.list = newList
			createMosaicImageInput(rightPayload, rightSubCanvas, acc)
		} else {
			createMosaicImageInput(rightPayload, rightSubCanvas, acc)
		}

	}
}

func divideCanvas(payload []DataPair, canvas Rectangle) (Rectangle, Rectangle) {
	payloadTotal := 0.0
	payloadCutoff := 0.0
	for _, item := range payload {
		payloadTotal += item.value
	}
	for i, item := range payload {
		if i == 0 || payloadCutoff+item.value <= payloadTotal/2 {
			payloadCutoff += item.value
		}
	}
	var leftSubCanvas, rightSubCanvas Rectangle
	if canvas.direction { //vertical
		alfa := int(math.Floor((float64(canvas.width) * payloadCutoff / payloadTotal)))
		leftSubCanvas = Rectangle{
			x:         canvas.x,
			y:         canvas.y,
			width:     alfa,
			height:    canvas.height,
			direction: !canvas.direction,
		}
		rightSubCanvas = Rectangle{
			x:         canvas.x + alfa,
			y:         canvas.y,
			width:     canvas.width - alfa,
			height:    canvas.height,
			direction: !canvas.direction,
		}
	} else { //horizontal
		beta := int(math.Floor(math.Floor(float64(canvas.height) * payloadCutoff / payloadTotal)))
		leftSubCanvas = Rectangle{
			x:         canvas.x,
			y:         canvas.y,
			width:     canvas.width,
			height:    beta,
			direction: !canvas.direction,
		}
		rightSubCanvas = Rectangle{
			x:         canvas.x,
			y:         canvas.y + beta,
			width:     canvas.width,
			height:    canvas.height - beta,
			direction: !canvas.direction,
		}
	}
	return leftSubCanvas, rightSubCanvas
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
