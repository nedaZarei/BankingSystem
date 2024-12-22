package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/nedaZarei/BankingSystem/model"
)

var db *sql.DB

const (
	host     = "postgres"
	port     = 5432
	user     = "neda.z"
	password = "nz2003nz"
	dbname   = "banking_system"
)

func Connect() error {
	var err error
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	return db.Ping()
}

func CreateBank(bank *model.Bank) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	_, err := db.Exec("INSERT INTO bank (name, headquarter_address) VALUES ($1, $2)",
		bank.Name, bank.HeadquarterAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfully created bank")
}

func RegisterEmployee(login *model.EmployeeLogin, details *model.EmployeeDetails) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	err = tx.QueryRow(
		"INSERT INTO employee_login (username, password) VALUES ($1, $2) RETURNING employee_id",
		login.Username, login.Password).Scan(&details.EmployeeID)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	_, err = tx.Exec(
		"INSERT INTO employee_details (employee_id, first_name, last_name, position, department, salary, branch_id) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		details.EmployeeID, details.FirstName, details.LastName, details.Position, details.Department, details.Salary, details.BranchID)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("successfully registered employee")
}

func CreateBranch(branch *model.Branch) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	_, err := db.Exec("INSERT INTO branch (bank_id, address) VALUES ($1, $2)",
		branch.BankID, branch.Address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfully created branch")
}

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

func CreateAccount(number *model.AccountNumber, details *model.AccountDetails) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	err = tx.QueryRow(
		"INSERT INTO account_numbers (account_number) VALUES ($1) RETURNING account_id",
		number.AccountNumber).Scan(&details.AccountID)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	_, err = tx.Exec(
		"INSERT INTO account_details (account_id, account_type, account_password, balance, account_status, open_date, close_date, customer_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		details.AccountID, details.AccountType, details.AccountPassword, details.Balance,
		details.AccountStatus, details.OpenDate, details.CloseDate, details.CustomerID)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("successfully created account")
}

func GetAccount(accountNumber, accountPassword string) (*model.AccountDetails, error) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	var accountID int
	err := db.QueryRow(
		"SELECT account_id FROM account_numbers WHERE account_number = $1",
		accountNumber).Scan(&accountID)
	if err != nil {
		return nil, err
	}

	var details model.AccountDetails
	err = db.QueryRow(
		"SELECT account_id, account_type, account_password, balance, account_status, open_date, close_date, customer_id FROM account_details WHERE account_id = $1",
		accountID).Scan(
		&details.AccountID, &details.AccountType, &details.AccountPassword,
		&details.Balance, &details.AccountStatus, &details.OpenDate,
		&details.CloseDate, &details.CustomerID)
	if err != nil {
		return nil, err
	}

	if details.AccountPassword != accountPassword {
		return nil, fmt.Errorf("invalid account password")
	}

	return &details, nil
}

func CreateTransaction(transaction *model.Transaction) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err := db.Exec(
		"INSERT INTO transaction (source_account_id, destination_account_id, amount, transaction_type, transaction_date) VALUES ($1, $2, $3, $4, $5)",
		transaction.SourceAccountID, transaction.DestinationAccountID,
		transaction.Amount, transaction.TransactionType, transaction.TransactionDate)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("successfully created transaction")
}
