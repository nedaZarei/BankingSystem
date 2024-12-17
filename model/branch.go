package model

type Branch struct {
	BranchID int    `db:"branch_id"`
	BankID   int    `db:"bank_id"`
	Address  string `db:"address"`
}
