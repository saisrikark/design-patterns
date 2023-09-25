package main

import "fmt"

type ShapeFactoryI interface {
	Get() Shape
}

type ShapeFactory struct {
	shapeType         ShapeType
	dimensions        Dimensions
	combinationPreset CombinationPreset
}

func NewShapeFactory(shapeType ShapeType, dimensions Dimensions, combinationPreset CombinationPreset) ShapeFactory {
	return ShapeFactory{
		shapeType:         shapeType,
		dimensions:        dimensions,
		combinationPreset: combinationPreset,
	}
}

func (sf ShapeFactory) Get() Shape {
	switch {
	case sf.shapeType == rectangle:
		r := NewRectangle(sf.dimensions.dimension1, sf.dimensions.dimension2, sf.combinationPreset)
		return r
	case sf.shapeType == square:
		return NewSquare(sf.dimensions.dimension1, sf.combinationPreset)
	case sf.shapeType == circle:
		return NewCircle(sf.dimensions.dimension1, sf.combinationPreset)
	default:
		panic(fmt.Sprintf("unknown shape type %s", sf.shapeType))
	}
}
