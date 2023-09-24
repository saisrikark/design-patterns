package main

import "fmt"

type MicroSDI interface {
	Store(b []byte)
}

type MicroSD struct {
}

func (msd MicroSD) Store(b []byte) {
	fmt.Println("storing")
}
