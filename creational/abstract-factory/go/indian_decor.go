package main

import (
	"abstract-factory/pkg/chairs"
	"abstract-factory/pkg/sofas"
	"abstract-factory/pkg/tables"
)

type IndianDecor struct {
}

func (ad IndianDecor) NewChair() chairs.ChairFactory {
	return chairs.IndianChair{}
}

func (ad IndianDecor) NewSofa() sofas.SofaFactory {
	return sofas.IndianSofa{}
}

func (ad IndianDecor) NewTable() tables.TableFactory {
	return tables.IndianTable{}
}
