package main

import "fmt"

// 实现功能选项模式
type FinishedHouse struct {
	style                  int
	centralAirConditioning bool
	floorMaterial          string
	wallMaterial           string
}

type Option func(*FinishedHouse)

func NewFinishedHouse(options ...Option) *FinishedHouse {
	h := &FinishedHouse{
		style:                  0,
		centralAirConditioning: true,
		floorMaterial:          "wood",
		wallMaterial:           "paper",
	}

	for _, option := range options {
		option(h)
	}

	return h
}

func WithStyle(style int) Option {
	return func(h *FinishedHouse) {
		h.style = style
	}
}

func WithFloorMaterial(material string) Option {
	return func(h *FinishedHouse) {
		h.floorMaterial = material
	}
}

func WithWallMaterial(material string) Option {
	return func(h *FinishedHouse) {
		h.wallMaterial = material
	}
}

func WithCentralAirConditioning(centralAirConditioning bool) Option {
	return func(h *FinishedHouse) {
		h.centralAirConditioning = centralAirConditioning
	}
}

func main() {
	fmt.Printf("%+v\n", NewFinishedHouse())
	fmt.Printf("%+v\n", NewFinishedHouse(WithStyle(1), WithFloorMaterial("ground-tile"), WithCentralAirConditioning(false)))
}

// 功能选项模式让我们获得如下好处
// 1. 更漂亮、不随时间变化的公共API
// 2. 参数可读性更好
// 3. 配置选项高度可拓展
// 4. 使用更安全
