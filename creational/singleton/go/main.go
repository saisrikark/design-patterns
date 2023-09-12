package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var accessor *FileAccessor
var lock *sync.Mutex

type FileAccessor struct {
}

func GetAccessor() *FileAccessor {
	if accessor == nil {
		accessor = &FileAccessor{}
		lock = &sync.Mutex{}
	}
	return accessor
}

func (fa *FileAccessor) Read() {
	lock.Lock()
	defer lock.Unlock()

	dat, _ := os.ReadFile("file.txt")
	fmt.Println("data is ", string(dat))
}

func main() {
	for i := 0; i < 100; i++ {
		go GetAccessor().Read()
	}
	time.Sleep(time.Second * 10)
}
