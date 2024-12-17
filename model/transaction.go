package model

type Transaction struct {
	TransactionID        int     `db:"transaction_id"`
	SourceAccountID      int     `db:"source_account_id"`
	DestinationAccountID *int    `db:"destination_account_id"`
	Amount               float64 `db:"amount"`
	TransactionType      string  `db:"transaction_type"`
	TransactionDate      string  `db:"transaction_date"`
}
