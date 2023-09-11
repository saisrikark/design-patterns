package main

import (
	"abstract-factory/pkg/chairs"
	"abstract-factory/pkg/sofas"
	"abstract-factory/pkg/tables"
)

type AmericanDecor struct {
}

func (ad AmericanDecor) NewChair() chairs.ChairFactory {
	return chairs.AmericanChair{}
}

func (ad AmericanDecor) NewSofa() sofas.SofaFactory {
	return sofas.AmericanSofa{}
}

func (ad AmericanDecor) NewTable() tables.TableFactory {
	return tables.AmericanTable{}
}
