package chairs

type AmericanChair struct {
}

func (ac AmericanChair) ShowWood() string {
	return "cherry"
}

func (ac AmericanChair) ShowPolishing() string {
	return "minimal"
}
