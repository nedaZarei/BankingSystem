package model

type Loan struct {
	LoanID       int     `db:"loan_id"`
	CustomerID   int     `db:"customer_id"`
	LoanType     string  `db:"loan_type"`
	Amount       float64 `db:"amount"`
	InterestRate float64 `db:"interest_rate"`
	Duration     int     `db:"duration"`
	StartDate    string  `db:"start_date"`
	EndDate      string  `db:"end_date"`
	LoanStatus   string  `db:"loan_status"`
}
