package main

import (
	"abstract-factory/pkg/chairs"
	"abstract-factory/pkg/sofas"
	"abstract-factory/pkg/tables"
)

type ItalianDecor struct {
}

func (ad ItalianDecor) NewChair() chairs.ChairFactory {
	return chairs.ItalianChair{}
}

func (ad ItalianDecor) NewSofa() sofas.SofaFactory {
	return sofas.ItalianSofa{}
}

func (ad ItalianDecor) NewTable() tables.TableFactory {
	return tables.ItalianTable{}
}
