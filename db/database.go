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
	_, err := db.Exec("INSERT INTO bank (name, headquarter_address) VALUES ($1, $2)", bank.Name, bank.HeadquarterAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfully created bank")
}

func RegisterEmployee(employee *model.Employee) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	_, err := db.Exec("INSERT INTO employee (username, password, first_name, last_name, position, department, salary, branch_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", employee.Username, employee.Password, employee.FirstName, employee.LastName, employee.Position, employee.Department, employee.Salary, employee.BranchID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("successfully registered employee")
}

func CreateBranch(branch *model.Branch) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	_, err := db.Exec("INSERT INTO branch (bank_id, address) VALUES ($1, $2)", branch.BankID, branch.Address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfully created branch")
}

func RegisterCostumer(costumer *model.Customer) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	_, err := db.Exec("INSERT INTO customer (username, password, first_name, last_name, birth_date, phone_number, email, address, customer_type, bank_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", costumer.Username, costumer.Password, costumer.FirstName, costumer.LastName, costumer.BirthDate, costumer.PhoneNumber, costumer.Email, costumer.Address, costumer.CustomerType, costumer.BankID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfully registered")
}

func LoginCustomer(username, password string) (*model.Customer, error) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	var customer model.Customer
	err := db.QueryRow("SELECT * FROM customer WHERE username = $1 AND password = $2", username, password).Scan(&customer.CustomerID, &customer.Username, &customer.Password, &customer.FirstName, &customer.LastName, &customer.BirthDate, &customer.PhoneNumber, &customer.Email, &customer.Address, &customer.CustomerType, &customer.BankID)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func CreateAccount(account *model.Account) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	_, err := db.Exec("INSERT INTO account (account_number, account_type, account_password, balance, account_status, open_date, close_date, customer_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", account.AccountNumber, account.AccountType, account.AccountPassword, account.Balance, account.AccountStatus, account.OpenDate, account.CloseDate, account.CustomerID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfully created account")
}

func GetAccount(accountNumber, accountPassword string) (*model.Account, error) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	var account model.Account
	err := db.QueryRow("SELECT * FROM account WHERE account_number = $1", accountNumber).Scan(&account.AccountID, &account.AccountNumber, &account.AccountType, &account.AccountPassword, &account.Balance, &account.AccountStatus, &account.OpenDate, &account.CloseDate, &account.CustomerID)
	if err != nil {
		return nil, err
	}
	if account.AccountPassword != accountPassword {
		return nil, fmt.Errorf("invalid account password")
	}
	return &account, nil
}

func CreateTransaction(transaction *model.Transaction) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err := db.Exec("INSERT INTO transaction (source_account_id, destination_account_id, amount, transaction_type, transaction_date) VALUES ($1, $2, $3, $4, $5)", transaction.SourceAccountID, transaction.DestinationAccountID, transaction.Amount, transaction.TransactionType, transaction.TransactionDate)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("successfully created transaction")
}
