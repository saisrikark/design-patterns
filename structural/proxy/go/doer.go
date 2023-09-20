package main

import "fmt"

type DockerI interface {
	Do()
}

type basicDoer struct {
}

func (bd basicDoer) Do() {
	fmt.Println("doing")
}

type proxyDoer struct {
}

func (pd proxyDoer) Do() {
	fmt.Println("about to do")
	bd := basicDoer{}
	bd.Do()
	fmt.Println("done doing")
}
