package main

import "fmt"

func NewDecor(decorType string) DecorAbstractFactory {
	var decor DecorAbstractFactory

	switch {
	case decorType == "indian":
		decor = IndianDecor{}
	case decorType == "american":
		decor = AmericanDecor{}
	default:
		decor = ItalianDecor{}
	}

	return decor
}

func main() {
	decorType := "indian"
	decor := NewDecor(decorType)

	fmt.Printf("\nchair uses %s wood", decor.NewChair().ShowWood())
	fmt.Printf("\ntable has %d legs", decor.NewTable().Legs())
	fmt.Printf("\nsofa uses %s material", decor.NewSofa().ShowCloth())
	fmt.Println()
}
