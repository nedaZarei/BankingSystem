package model

type Bank struct {
	BankID             int    `db:"bank_id"`
	Name               string `db:"name"`
	HeadquarterAddress string `db:"headquarter_address"`
}
