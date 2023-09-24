package main

import "fmt"

type SimpleI interface {
	op()
}

type SimpleFunctionality struct {
}

func (sf SimpleFunctionality) op() {
	cf := ComplexFunctionality{}
	fmt.Println("calling op1 through op without exposing the complexity")
	cf.op1()
	fmt.Println("calling op2 through op without exposing the complexity")
	cf.op2()
	fmt.Println("calling op3 through op without exposing the complexity")
	cf.op3()
	fmt.Println("calling op4 through op without exposing the complexity")
	cf.op4()
	fmt.Println("calling op5 through op without exposing the complexity")
	cf.op5()
	fmt.Println("calling op6 through op without exposing the complexity")
	cf.op6()
}
