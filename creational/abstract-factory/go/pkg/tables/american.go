package tables

type AmericanTable struct {
}

func (at AmericanTable) Seats() int {
	return 6
}

func (at AmericanTable) Legs() int {
	return 6
}
