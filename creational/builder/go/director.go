package main

func NewNeapolitanMaker(pizzaMaker PizzaMaker) PizzaMaker {
	return pizzaMaker.WithCrust("thin").WithFirstTopping("mozzarella").WithSecondTopping("marinara").WithThirdTopping("basil")
}
