package main

import (
	"fmt"
	"time"
)

type IBulky interface {
	print(string)
	clone() IBulky
}

type Bulky struct {
	message string
}

func NewBulky() IBulky {
	time.Sleep(time.Second * 3)
	return Bulky{
		message: "random",
	}
}

func (b Bulky) print(_ string) {
	print(fmt.Sprintf("%d", (time.Now().Second()))+" "+b.message, "\n")
}

func (b Bulky) clone() IBulky {
	return Bulky{message: b.message}
}

func main() {
	bulky1 := NewBulky()
	bulky1.print("")
	bulky2 := bulky1.clone()
	bulky2.print("")
	bulky3 := bulky1.clone()
	bulky3.print("")
}
