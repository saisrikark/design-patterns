package main

import (
	"fmt"
)

type ShapeType string

const (
	rectangle = ShapeType("rectangle")
	square    = ShapeType("square")
	circle    = ShapeType("circle")
)

type Corner string

const (
	sharp  = Corner("Sharp")
	curved = Corner("Curved")
)

type LineThickness int

const (
	thin   = LineThickness(1)
	medium = LineThickness(2)
	thick  = LineThickness(3)
)

func (lt LineThickness) String() string {
	switch {
	case lt == thin:
		return "thin"
	case lt == medium:
		return "medium"
	case lt == thick:
		return "thick"
	default:
		panic(fmt.Sprintf("unknown thickness %d", lt))
	}
}

type Dimensions struct {
	dimension1 int
	dimension2 int
}

type IntrinsicCombinations struct {
	corner        Corner
	lineThickness LineThickness
}

type CombinationPreset int

const (
	SharpCornerThinLine    = CombinationPreset(1)
	SharpCornerMediumLine  = CombinationPreset(2)
	SharpCornerThickLine   = CombinationPreset(3)
	CurvedCornerThinLine   = CombinationPreset(4)
	CurvedCornerMediumLine = CombinationPreset(5)
	CurvedCornerThickLine  = CombinationPreset(6)
)

var (
	presets = map[CombinationPreset]*IntrinsicCombinations{
		SharpCornerThinLine: {
			corner:        sharp,
			lineThickness: thin,
		},
		SharpCornerMediumLine: {
			corner:        sharp,
			lineThickness: medium,
		},
		SharpCornerThickLine: {
			corner:        sharp,
			lineThickness: thick,
		},
		CurvedCornerThinLine: {
			corner:        curved,
			lineThickness: thin,
		},
		CurvedCornerMediumLine: {
			corner:        curved,
			lineThickness: medium,
		},
		CurvedCornerThickLine: {
			corner:        curved,
			lineThickness: thick,
		},
	}
)
