package model

type CustomerLogin struct {
	Username   string `db:"username"`
	Password   string `db:"password"`
	CustomerID int    `db:"customer_id"`
}

type CustomerEmail struct {
	Email      string `db:"email"`
	CustomerID int    `db:"customer_id"`
}

type CustomerDetails struct {
	CustomerID   int    `db:"customer_id"`
	FirstName    string `db:"first_name"`
	LastName     string `db:"last_name"`
	BirthDate    string `db:"birth_date"`
	PhoneNumber  string `db:"phone_number"`
	Address      string `db:"address"`
	CustomerType string `db:"customer_type"`
	BankID       int    `db:"bank_id"`
}
