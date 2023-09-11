package sofas

type AmericanSofa struct {
}

func (as AmericanSofa) ShowCloth() string {
	return "leather"
}

func (as AmericanSofa) HasMotion() bool {
	return true
}
