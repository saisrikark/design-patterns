package main

import (
	"abstract-factory/pkg/chairs"
	"abstract-factory/pkg/sofas"
	"abstract-factory/pkg/tables"
)

type DecorAbstractFactory interface {
	NewChair() chairs.ChairFactory
	NewSofa() sofas.SofaFactory
	NewTable() tables.TableFactory
}
