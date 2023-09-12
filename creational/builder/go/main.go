package main

import "fmt"

func main() {
	pizzaBuilder1 := NewPizzaMaker()
	fmt.Println(
		pizzaBuilder1.
			Bake().
			Description())

	pizzaBuilder2 := NewPizzaMaker()
	fmt.Println(
		pizzaBuilder2.
			WithCrust("thin").
			WithFirstTopping("mushroom").
			WithSecondTopping("paneer").
			WithThirdTopping("mozzarella").
			Bake().
			Description())

	pizzaBuilder3 := NewPizzaMaker()
	fmt.Println(NewNeapolitanMaker(pizzaBuilder3).Bake().Description())
}
