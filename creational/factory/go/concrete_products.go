package main

import "fmt"

type Circle struct {
}

func (c Circle) Display() {
	fmt.Println("CIRCLE I AM")
}

type Square struct {
}

func (s Square) Display() {
	fmt.Println("SQUARE I AM")
}

type Rectangle struct {
}

func (r Rectangle) Display() {
	fmt.Println("RECTANGLE I AM")
}

type Unknown struct {
}

func (u Unknown) Display() {
	fmt.Println("WHAT IN THE WORLD DID YOU GIVE?")
}
