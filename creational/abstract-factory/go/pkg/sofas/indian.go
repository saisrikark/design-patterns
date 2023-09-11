package sofas

type IndianSofa struct {
}

func (as IndianSofa) ShowCloth() string {
	return "cloth"
}

func (as IndianSofa) HasMotion() bool {
	return false
}
