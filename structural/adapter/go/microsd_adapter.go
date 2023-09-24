package main

import "fmt"

type SDCard interface {
	Store(b []byte)
}

type SDCardAdapter struct {
	microSD MicroSD
}

func (sda SDCardAdapter) Store(b []byte) {
	fmt.Println("sd card adapter doing adapting ops")
	// blah blah blah
	sda.microSD.Store(b)
}
