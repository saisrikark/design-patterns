package main

type SignalerI interface {
	SetVolume(int)
	IncreaseVolume()
	DecreaseVolume()
	IncreaseChannel()
	DecreaseChannel()
}

type Signaler struct {
}

func (s Signaler) SetVolume(int) {
}

func (s Signaler) IncreaseVolume() {
}

func (s Signaler) DecreaseVolume() {
}

func (s Signaler) IncreaseChannel() {
}

func (s Signaler) DecreaseChannel() {
}
