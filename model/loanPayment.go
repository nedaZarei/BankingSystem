package model

type LoanPayment struct {
	PaymentID     int     `db:"payment_id"`
	LoanID        int     `db:"loan_id"`
	PaymentAmount float64 `db:"payment_amount"`
	DueDate       string  `db:"due_date"`
	PaymentDate   string  `db:"payment_date"`
	PaymentStatus string  `db:"payment_status"`
}
