package main

import "fmt"

type ComponentI interface {
	Cost() int
}

type Desktop struct {
	subComponents []ComponentI
}

func (d Desktop) Cost() int {
	cost := 0
	for _, component := range d.subComponents {
		cost += component.Cost()
	}
	return cost
}

type Case struct {
	cost          int
	subComponents []ComponentI
}

func (c Case) Cost() int {
	cost := c.cost
	for _, component := range c.subComponents {
		cost += component.Cost()
	}
	return cost
}

type PowerThing struct {
	cost int
}

func (pt PowerThing) Cost() int {
	return pt.cost
}

type CPU struct {
	cost int
}

func (cpu CPU) Cost() int {
	return cpu.cost
}

type Motherboard struct {
	cost          int
	subComponents []ComponentI
}

func (m Motherboard) Cost() int {
	cost := m.cost
	for _, component := range m.subComponents {
		cost += component.Cost()
	}
	return cost
}

type Fan struct {
	cost int
}

func (f Fan) Cost() int {
	return f.cost
}

func main() {
	subComponents := []ComponentI{
		PowerThing{
			cost: 5,
		},
		Fan{
			cost: 10,
		},
		Case{
			cost: 10,
			subComponents: []ComponentI{
				Motherboard{
					cost: 20,
					subComponents: []ComponentI{
						CPU{
							cost: 100,
						},
					},
				},
			},
		},
	}
	desktop := Desktop{
		subComponents: subComponents,
	}
	fmt.Println(desktop.Cost())
}
