package model

type Account struct {
	AccountID       int     `db:"account_id"`
	AccountNumber   string  `db:"account_number"`
	AccountType     string  `db:"account_type"`
	AccountPassword string  `db:"account_password"`
	Balance         float64 `db:"balance"`
	AccountStatus   string  `db:"account_status"`
	OpenDate        string  `db:"open_date"`
	CloseDate       *string `db:"close_date"`
	CustomerID      int     `db:"customer_id"`
}
