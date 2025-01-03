package db

import (
	"fmt"
	"log"

	"github.com/nedaZarei/BankingSystem/model"
)

func RegisterCustomer(login *model.CustomerLogin, email *model.CustomerEmail, details *model.CustomerDetails) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	err = tx.QueryRow(
		"INSERT INTO customer_login (username, password) VALUES ($1, $2) RETURNING customer_id",
		login.Username, login.Password).Scan(&details.CustomerID)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	email.CustomerID = details.CustomerID

	_, err = tx.Exec(
		"INSERT INTO customer_email (email, customer_id) VALUES ($1, $2)",
		email.Email, email.CustomerID)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	_, err = tx.Exec(
		"INSERT INTO customer_details (customer_id, first_name, last_name, birth_date, phone_number, address, customer_type, bank_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		details.CustomerID, details.FirstName, details.LastName, details.BirthDate, details.PhoneNumber, details.Address, details.CustomerType, details.BankID)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("successfully registered customer")
}

func LoginCustomer(username, password string) (*model.CustomerDetails, error) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	var customerID int
	err := db.QueryRow(
		"SELECT customer_id FROM customer_login WHERE username = $1 AND password = $2",
		username, password).Scan(&customerID)
	if err != nil {
		return nil, err
	}

	var details model.CustomerDetails
	err = db.QueryRow(
		"SELECT customer_id, first_name, last_name, birth_date, phone_number, address, customer_type, bank_id FROM customer_details WHERE customer_id = $1",
		customerID).Scan(
		&details.CustomerID, &details.FirstName, &details.LastName, &details.BirthDate,
		&details.PhoneNumber, &details.Address, &details.CustomerType, &details.BankID)
	if err != nil {
		return nil, err
	}

	return &details, nil
}

func UpdateCustomer(details *model.CustomerDetails, email string) error {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.Exec(
		"UPDATE customer_details SET first_name = $1, last_name = $2, birth_date = $3, phone_number = $4, address = $5, customer_type = $6, bank_id = $7 WHERE customer_id = $8",
		details.FirstName, details.LastName, details.BirthDate, details.PhoneNumber, details.Address, details.CustomerType, details.BankID, details.CustomerID)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	_, err = tx.Exec(
		"UPDATE customer_email SET email = $1 WHERE customer_id = $2",
		email, details.CustomerID)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("successfully updated customer")
	return nil
}

func DeleteCustomer(customerID int) error {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.Exec("DELETE FROM customer_email WHERE customer_id = $1", customerID)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	_, err = tx.Exec("DELETE FROM customer_details WHERE customer_id = $1", customerID)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	_, err = tx.Exec("DELETE FROM customer_login WHERE customer_id = $1", customerID)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("successfully deleted customer")
	return nil
}
