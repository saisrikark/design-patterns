package main

type ShapeFactory interface {
	Create(shapeType string) Shape
}
