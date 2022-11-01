package main

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
