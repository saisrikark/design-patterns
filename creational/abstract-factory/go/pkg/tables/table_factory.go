package tables

type TableFactory interface {
	Seats() int
	Legs() int
}
