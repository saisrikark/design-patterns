package main

type PhotoTaker interface {
	TakePhoto()
}

type Phone struct {
	microSD MicroSDI
}

func (p Phone) TakePhoto() {
	photo := []byte("some photo data")
	p.microSD.Store(photo)
}

type Camera struct {
	sdCard SDCard
}

func (c Camera) TakePhoto() {
	photo := []byte("camera photo")
	c.sdCard.Store(photo)
}
