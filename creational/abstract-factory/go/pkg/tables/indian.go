package tables

type IndianTable struct {
}

func (at IndianTable) Seats() int {
	return 12
}

func (at IndianTable) Legs() int {
	return 12
}
