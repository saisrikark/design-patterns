package chairs

type ItalianChair struct {
}

func (ac ItalianChair) ShowWood() string {
	return "oak"
}

func (ac ItalianChair) ShowPolishing() string {
	return "polished"
}
