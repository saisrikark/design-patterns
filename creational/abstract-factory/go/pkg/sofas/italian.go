package sofas

type ItalianSofa struct {
}

func (as ItalianSofa) ShowCloth() string {
	return "leather"
}

func (as ItalianSofa) HasMotion() bool {
	return false
}
