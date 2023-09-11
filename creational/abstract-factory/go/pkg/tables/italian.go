package tables

type ItalianTable struct {
}

func (at ItalianTable) Seats() int {
	return 12
}

func (at ItalianTable) Legs() int {
	return 12
}
