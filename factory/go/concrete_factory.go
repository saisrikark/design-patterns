package main

type ShapeConcreteFactory struct {
}

func (cf ShapeConcreteFactory) Create(shapeType string) Shape {
	switch {
	case shapeType == "CIRCLE":
		return &Circle{}
	case shapeType == "SQUARE":
		return &Square{}
	case shapeType == "RECTANGLE":
		return &Rectangle{}
	default:
		return &Unknown{}
	}
}
