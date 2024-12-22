package main

import (
	"fmt"
	"log"
	"net/http"

	db "github.com/nedaZarei/BankingSystem/db"
	"github.com/nedaZarei/BankingSystem/model"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to database")

	bank := &model.Bank{
		Name:               "Sepah Bank",
		HeadquarterAddress: "123 Main St",
	}
	db.CreateBank(bank)

	branch := &model.Branch{
		BankID:  1,
		Address: "456 Elm St",
	}
	db.CreateBranch(branch)

	employeeLogin := &model.EmployeeLogin{
		Username:   "rossbenson",
		Password:   "password",
		EmployeeID: 0, //will be set by database
	}

	employeeDetails := &model.EmployeeDetails{
		FirstName:  "ross",
		LastName:   "benson",
		Position:   "teller",
		Department: "operations",
		Salary:     50000,
		BranchID:   1,
	}
	db.RegisterEmployee(employeeLogin, employeeDetails)

	// Create customer records
	customerLogin := &model.CustomerLogin{
		Username:   "eddiebooker",
		Password:   "password2627",
		CustomerID: 0, //will be set by database
	}

	customerEmail := &model.CustomerEmail{
		Email: "eddiebooker@gmail.com",
	}

	customerDetails := &model.CustomerDetails{
		FirstName:    "eddie",
		LastName:     "booker",
		BirthDate:    "1990-01-01",
		PhoneNumber:  "1234567890",
		Address:      "3202 Trails end road",
		CustomerType: "Individual",
		BankID:       1,
	}
	db.RegisterCustomer(customerLogin, customerEmail, customerDetails)

	customer, err := db.LoginCustomer("eddiebooker", "password2627")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("logged in customer: %+v\n", customer)

	accountNumber := &model.AccountNumber{
		AccountNumber: "1234567890",
		AccountID:     0, //will be set by database
	}

	accountDetails := &model.AccountDetails{
		AccountType:     "Savings",
		AccountPassword: "passpass",
		Balance:         1000,
		AccountStatus:   "Active",
		OpenDate:        "2025-01-01",
		CustomerID:      1,
	}
	db.CreateAccount(accountNumber, accountDetails)

	fmt.Println("getting the customer 1 account:")
	db.GetAccount("1234567890", "passpass")

	transaction := &model.Transaction{
		SourceAccountID:      1,
		DestinationAccountID: nil,
		Amount:               1000,
		TransactionType:      "Deposit",
		TransactionDate:      "2025-01-01",
	}
	db.CreateTransaction(transaction)

	//to keep the app running
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Banking System Server")
	})
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
