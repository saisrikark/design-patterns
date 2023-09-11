package chairs

type IndianChair struct {
}

func (ac IndianChair) ShowWood() string {
	return "sandalwood"
}

func (ac IndianChair) ShowPolishing() string {
	return "polished"
}
