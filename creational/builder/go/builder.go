package main

import "fmt"

type Pizza struct {
	topping1  string
	topping2  string
	topping3  string
	crustType string
}

func (pizza Pizza) Description() string {
	return fmt.Sprintf("baked a %s crust pizza with %s, %s and %s", pizza.crustType, pizza.topping1, pizza.topping2, pizza.topping3)
}

type PizzaMaker struct {
	topping1  string
	topping2  string
	topping3  string
	crustType string
}

func (pm PizzaMaker) WithFirstTopping(topping string) PizzaMaker {
	pm.topping1 = topping
	return pm
}

func (pm PizzaMaker) WithSecondTopping(topping string) PizzaMaker {
	pm.topping2 = topping
	return pm
}

func (pm PizzaMaker) WithThirdTopping(topping string) PizzaMaker {
	pm.topping3 = topping
	return pm
}

func (pm PizzaMaker) WithCrust(crustType string) PizzaMaker {
	pm.crustType = crustType
	return pm
}

func NewPizzaMaker() PizzaMaker {
	return PizzaMaker{}
}

func (pm PizzaMaker) Bake() Pizza {
	if pm.topping1 == "" {
		pm.topping1 = "pineapple"
	}
	if pm.topping2 == "" {
		pm.topping2 = "mozzarella"
	}
	if pm.topping3 == "" {
		pm.topping3 = "avocado"
	}
	if pm.crustType == "" {
		pm.crustType = "thick"
	}

	pizza := Pizza{
		topping1:  pm.topping1,
		topping2:  pm.topping2,
		topping3:  pm.topping3,
		crustType: pm.crustType,
	}

	return pizza
}
