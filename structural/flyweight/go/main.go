package main

import "fmt"

func main() {
	shapeTypes := []ShapeType{
		rectangle,
		square,
		circle}
	dimensions := []Dimensions{
		{dimension1: 2, dimension2: 3},
		{dimension1: 4},
		{dimension1: 6},
	}
	combinationPresets := []CombinationPreset{
		SharpCornerMediumLine,
		SharpCornerThickLine,
		SharpCornerThinLine,
	}

	for i := 0; i < len(shapeTypes); i++ {
		shape := NewShapeFactory(shapeTypes[i], dimensions[i], combinationPresets[i])
		fmt.Printf(
			"shape type is %s area is %f\n",
			shape.Get().Show(),
			shape.Get().Area())
	}
}
