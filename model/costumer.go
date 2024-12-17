package model

type Customer struct {
	CustomerID   int    `db:"customer_id"`
	Username     string `db:"username"`
	Password     string `db:"password"`
	FirstName    string `db:"first_name"`
	LastName     string `db:"last_name"`
	BirthDate    string `db:"birth_date"`
	PhoneNumber  string `db:"phone_number"`
	Email        string `db:"email"`
	Address      string `db:"address"`
	CustomerType string `db:"customer_type"`
	BankID       int    `db:"bank_id"`
}
