package main

type ConcreteFactory struct {
}

func (cf ConcreteFactory) Create(shapeType string) Shape {
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
