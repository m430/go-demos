package main

import "fmt"

// 实现功能选项模式
type FinishedHouse struct {
	style                  int
	centralAirConditioning bool
	floorMaterial          string
	wallMaterial           string
}

type Options struct {
	Style                  int
	CentralAirConditioning bool
	FloorMaterial          string
	WallMaterial           string
}

func NewFinishedHouse(options *Options) *FinishedHouse {
	var style int = 0
	var centralAirConditioning = true
	var floorMaterial = "wood"
	var wallMaterial = "paper"

	if options != nil {
		style = options.Style
		centralAirConditioning = options.CentralAirConditioning
		floorMaterial = options.FloorMaterial
		wallMaterial = options.WallMaterial
	}

	return &FinishedHouse{
		style:                  style,
		centralAirConditioning: centralAirConditioning,
		floorMaterial:          floorMaterial,
		wallMaterial:           wallMaterial,
	}
}

func main() {
	fmt.Printf("%+v\n", NewFinishedHouse(nil))
	fmt.Printf("%+v\n", NewFinishedHouse(&Options{
		Style:                  1,
		CentralAirConditioning: false,
		FloorMaterial:          "ground-tile",
		WallMaterial:           "paper",
	}))
}
