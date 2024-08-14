package models

import (
	"slices"
)

type DomainSearchByType int

const (
	DomainSearchByDomainName DomainSearchByType = iota + 1
	DomainSearchByAccountNo
	DomainSearchByClientNo
)

func DomainSearchByColumn(searchByType DomainSearchByType) string {
	switch searchByType {
	case DomainSearchByAccountNo:
		return "Account_Number"
	case DomainSearchByClientNo:
		return "Client_Number"
	default:
		return "Domain_Name"
	}
}

func DomainSearchByTypeFromCode(code string) DomainSearchByType {
	switch code {
	case "account":
		return DomainSearchByAccountNo
	case "client":
		return DomainSearchByClientNo
	default:
		return DomainSearchByDomainName
	}
}

var DomainActiveStatusses []string = []string{"Ready", "Termination Period"}

type Domain struct {
	ID           int    `db:"ID"`
	DomainNumber string `db:"Domain_Number"`
	DomainName   string `db:"Domain_Name"`
	AccountNo    string `db:"Account_Number"`
	ClientNo     string `db:"Client_Number"`
	BillForReg   bool   `db:"Bill_For_Registration"`
	Status       string `db:"Status"`
}

func (d *Domain) Active() bool {
	return slices.Contains(DomainActiveStatusses, d.Status)
}
