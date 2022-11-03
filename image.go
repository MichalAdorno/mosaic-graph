package main

import (
	"errors"
	"fmt"
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
	fmt.Println("*****************")
	fmt.Printf("acc = %v\n", acc)
	if payload == nil {
		return
	}
	if len(payload) == 1 {
		fmt.Printf("x/payload = %v\n", payload)
		fmt.Printf("x/acc = %v\n", acc)
		//acc = &MosaicImageInput{list: append(acc.list, canvas)}
		fmt.Printf("x/newAcc = %v\n", acc)
		return
	}
	leftPayload, rightPayload, _ := splitIntoTwo(payload)
	//fmt.Printf("leftPayload=%v, rightPayload=%v, err=%v\n", leftPayload, rightPayload, err)

	leftSubCanvas, rightSubCanvas := divideCanvas(payload, canvas)

	//if leftPayload != nil || len(leftPayload) > 0 {
	//var newAcc MosaicImageInput
	if len(leftPayload) == 1 {
		newList := append(acc.list, leftSubCanvas)
		acc.list = newList
		fmt.Printf("leftSubCanvas/payload = %v\n", payload)
		fmt.Printf("leftSubCanvas/leftPayload = %v\n", leftPayload)
		fmt.Printf("leftSubCanvas = %v\n", leftSubCanvas)
		fmt.Printf("leftSubCanvas/newAcc = %v\n", acc)
		createMosaicImageInput(leftPayload, leftSubCanvas, acc)
	} else {
		createMosaicImageInput(leftPayload, leftSubCanvas, acc)
	}

	if rightPayload != nil || len(rightPayload) > 0 {
		if len(rightPayload) == 1 {
			newList := append(acc.list, rightSubCanvas)
			acc.list = newList
			fmt.Printf("rightSubCanvas/payload = %v\n", payload)
			fmt.Printf("rightSubCanvas/rightPayload = %v\n", rightPayload)
			fmt.Printf("rightSubCanvas = %v\n", rightSubCanvas)
			fmt.Printf("rightSubCanvas/newAcc = %v\n", acc)
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
		leftSubCanvas = Rectangle{
			x:         canvas.x,
			y:         canvas.y,
			width:     int(float64(canvas.width) * payloadCutoff / payloadTotal),
			height:    canvas.height,
			direction: !canvas.direction,
		}
		rightSubCanvas = Rectangle{
			x:         canvas.x + int(float64(canvas.width)*payloadCutoff/payloadTotal),
			y:         canvas.y,
			width:     int((1 - payloadCutoff/payloadTotal) * float64(canvas.width)),
			height:    canvas.height,
			direction: !canvas.direction,
		}
	} else { //horizontal
		leftSubCanvas = Rectangle{
			x:         canvas.x,
			y:         canvas.y,
			width:     canvas.width,
			height:    int(float64(canvas.height) * payloadCutoff / payloadTotal),
			direction: !canvas.direction,
		}
		rightSubCanvas = Rectangle{
			x:         canvas.x,
			y:         canvas.y + int(float64(canvas.height)*payloadCutoff/payloadTotal),
			width:     canvas.width,
			height:    int((1 - payloadCutoff/payloadTotal) * float64(canvas.height)),
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
