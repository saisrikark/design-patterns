package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Show() string
	Area() float64
}

type Rectangle struct {
	width       int
	breadth     int
	combination *IntrinsicCombinations
}

func NewRectangle(width int, breadth int, combinationPreset CombinationPreset) Rectangle {
	r := Rectangle{
		width:       width,
		breadth:     breadth,
		combination: presets[combinationPreset],
	}
	return r
}

func (r Rectangle) Show() string {
	return fmt.Sprintf("%s cornered %s lined %s", r.combination.corner, r.combination.lineThickness, "Rectangle")
}

func (r Rectangle) Area() float64 {
	return float64(r.breadth * r.width)
}

type Square struct {
	sideSize    int
	combination *IntrinsicCombinations
}

func NewSquare(sideSize int, combinationPreset CombinationPreset) Square {
	return Square{
		sideSize:    sideSize,
		combination: presets[combinationPreset],
	}
}

func (s Square) Show() string {
	return fmt.Sprintf("%s cornered %s lined %s", s.combination.corner, s.combination.lineThickness, "Square")
}

func (s Square) Area() float64 {
	return float64(s.sideSize * s.sideSize)
}

type Circle struct {
	radius      int
	combination *IntrinsicCombinations
}

func NewCircle(radius int, combinationPreset CombinationPreset) Circle {
	return Circle{
		radius:      radius,
		combination: presets[combinationPreset],
	}
}

func (c Circle) Show() string {
	return fmt.Sprintf("%s cornered %s lined %s", c.combination.corner, c.combination.lineThickness, "Circle")
}

func (c Circle) Area() float64 {
	rsquare := c.radius * c.radius
	return math.Pi * float64(rsquare)
}
