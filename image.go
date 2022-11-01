package main

type Data struct {
	category string
	value    float64
}

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

func Mosaic(payload []Data, canvas Rectangle) *MosaicImage {
	return split(payload, canvas, &MosaicImage{list: []Rectangle{}})
}

func split(payload []Data, canvas Rectangle, acc *MosaicImage) *MosaicImage {
	if payload == nil {
		return nil
	}
	if len(payload) == 1 {
		return &MosaicImage{list: append(acc.list, canvas)}
	}
	return nil
}
