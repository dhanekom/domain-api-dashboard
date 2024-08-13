package models

import (
	"slices"
)

var DomainActiveStatusses []string = []string{"Ready", "Termination Period"}

type Domain struct {
	DomainName string
	AccountNo  string
	ClientNo   string
	BillForReg bool
	Status     string
}

func (d *Domain) Active() bool {
	return slices.Contains(DomainActiveStatusses, d.Status)
}
