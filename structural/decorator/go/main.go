package main

import "fmt"

func main() {
	rawBill := RawBill{cost: 123.1}
	taxedBill := TaxBill{rawBill: rawBill}

	fmt.Println("without tax", rawBill.Cost())
	fmt.Println("with tax", taxedBill.Cost())
}
