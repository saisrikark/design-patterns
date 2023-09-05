package main

type Factory interface {
	Create(shapeType string) Shape
}
