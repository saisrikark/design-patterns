package main

type BillI interface {
	Cost() float64
}

type RawBill struct {
	cost float64
}

func (r RawBill) Cost() float64 {
	return r.cost
}

type TaxBill struct {
	rawBill RawBill
}

func (tb TaxBill) Cost() float64 {
	return tb.rawBill.Cost() * 1.3
}
