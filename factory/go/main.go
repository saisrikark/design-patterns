package main

func main() {
	cf := ConcreteFactory{}
	shapeTypes := [4]string{"CIRCLE", "SQUARE", "RECTANGLE", "JINKAMARINA"}

	for _, shapeType := range shapeTypes {
		shape := cf.Create(shapeType)
		shape.Display()
	}
}
