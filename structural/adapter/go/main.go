package main

func main() {
	microSD := MicroSD{}
	phone := Phone{microSD: microSD}
	phone.TakePhoto()

	camera := Camera{sdCard: SDCardAdapter{microSD: microSD}}
	camera.TakePhoto()
}
