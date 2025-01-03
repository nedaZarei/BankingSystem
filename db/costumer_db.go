package db

import (
	"database/sql"
	"fmt"

	"github.com/nedaZarei/BankingSystem/model"
	"golang.org/x/crypto/bcrypt"
)

func RegisterCustomer(login *model.CustomerLogin, email *model.CustomerEmail, details *model.CustomerDetails) error {
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	//hashing password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(login.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	err = tx.QueryRow(
		"INSERT INTO customer_login (username, password) VALUES ($1, $2) RETURNING customer_id",
		login.Username, hashedPassword).Scan(&details.CustomerID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert customer login: %w", err)
	}

	email.CustomerID = details.CustomerID

	_, err = tx.Exec(
		"INSERT INTO customer_email (email, customer_id) VALUES ($1, $2)",
		email.Email, email.CustomerID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert customer email: %w", err)
	}

	_, err = tx.Exec(
		"INSERT INTO customer_details (customer_id, first_name, last_name, birth_date, phone_number, address, customer_type, bank_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		details.CustomerID, details.FirstName, details.LastName, details.BirthDate, details.PhoneNumber, details.Address, details.CustomerType, details.BankID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert customer details: %w", err)
	}

	//committing transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	fmt.Println("Successfully registered customer")
	return nil
}

func LoginCustomer(username, password string) (*model.CustomerDetails, error) {
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	var hashedPassword string
	var customerID int

	//retrieving hashed password and customer id
	err := db.QueryRow(
		"SELECT password, customer_id FROM customer_login WHERE username = $1",
		username).Scan(&hashedPassword, &customerID)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("invalid username or password")
	} else if err != nil {
		return nil, fmt.Errorf("failed to retrieve customer login: %w", err)
	}

	//compering given password with hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return nil, fmt.Errorf("invalid username or password")
	}

	var details model.CustomerDetails

	err = db.QueryRow(
		"SELECT customer_id, first_name, last_name, birth_date, phone_number, address, customer_type, bank_id FROM customer_details WHERE customer_id = $1",
		customerID).Scan(
		&details.CustomerID, &details.FirstName, &details.LastName, &details.BirthDate,
		&details.PhoneNumber, &details.Address, &details.CustomerType, &details.BankID)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("customer details not found")
	} else if err != nil {
		return nil, fmt.Errorf("failed to retrieve customer details: %w", err)
	}

	return &details, nil
}

func UpdateCustomer(details *model.CustomerDetails, email string) error {
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	_, err = tx.Exec(
		"UPDATE customer_details SET first_name = $1, last_name = $2, birth_date = $3, phone_number = $4, address = $5, customer_type = $6, bank_id = $7 WHERE customer_id = $8",
		details.FirstName, details.LastName, details.BirthDate, details.PhoneNumber, details.Address, details.CustomerType, details.BankID, details.CustomerID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update customer details: %w", err)
	}

	_, err = tx.Exec(
		"UPDATE customer_email SET email = $1 WHERE customer_id = $2",
		email, details.CustomerID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update customer email: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	fmt.Println("Successfully updated customer")
	return nil
}

func DeleteCustomer(customerID int) error {
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	_, err = tx.Exec("DELETE FROM customer_email WHERE customer_id = $1", customerID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete customer email: %w", err)
	}

	_, err = tx.Exec("DELETE FROM customer_details WHERE customer_id = $1", customerID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete customer details: %w", err)
	}

	_, err = tx.Exec("DELETE FROM customer_login WHERE customer_id = $1", customerID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete customer login: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	fmt.Println("Successfully deleted customer")
	return nil
}
