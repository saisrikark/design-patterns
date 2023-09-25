package main

type RemoteI interface {
	Button1()
	Button2()
	Button3()
	Button4()
	Button5()
}

type LGRemote struct {
	signaler SignalerI
}

func (lr LGRemote) Button1() {
	lr.signaler.DecreaseChannel()
}

func (lr LGRemote) Button2() {
	lr.signaler.IncreaseChannel()
}

func (lr LGRemote) Button3() {
	lr.signaler.IncreaseVolume()
}

func (lr LGRemote) Button4() {
	lr.signaler.DecreaseVolume()
}

func (lr LGRemote) Button5() {
	lr.signaler.SetVolume(0)
}

type PhoneBluetoothRemote struct {
	signaler Signaler
}

func (pbr PhoneBluetoothRemote) Button1() {
	pbr.signaler.IncreaseChannel()
}

func (pbr PhoneBluetoothRemote) Button2() {
	pbr.signaler.DecreaseChannel()
}

func (pbr PhoneBluetoothRemote) Button3() {
	pbr.signaler.SetVolume(0)
}

func (pbr PhoneBluetoothRemote) Button4() {
	pbr.signaler.IncreaseVolume()
}

func (pbr PhoneBluetoothRemote) Button5() {
	pbr.signaler.DecreaseVolume()
}
