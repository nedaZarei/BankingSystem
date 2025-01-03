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

	customerLogin := &model.CustomerLogin{
		Username:   "davidmills",
		Password:   "password2627",
		CustomerID: 0, //will be set by database
	}

	customerEmail := &model.CustomerEmail{
		Email: "davidmillss8@gmail.com",
	}

	customerDetails := &model.CustomerDetails{
		FirstName:    "david",
		LastName:     "mills",
		BirthDate:    "1985-11-02",
		PhoneNumber:  "214555123",
		Address:      "12 main st",
		CustomerType: "Natural", //haghighi
		BankID:       1,
	}
	db.RegisterCustomer(customerLogin, customerEmail, customerDetails)

	customer, err := db.LoginCustomer("davidmills", "password2627")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("logged in customer: %+v\n", customer)

	accountNumber := &model.AccountNumber{
		AccountNumber: "1921072918",
		AccountID:     0, //will be set by database
	}

	accountDetails := &model.AccountDetails{
		AccountType:     "Deposit",
		AccountPassword: "1111",
		Balance:         5000,
		AccountStatus:   "Active",
		OpenDate:        "2025-01-01",
		CustomerID:      1,
	}
	db.CreateAccount(accountNumber, accountDetails)

	customerLogin2 := &model.CustomerLogin{
		Username:   "sallysmith",
		Password:   "salsal",
		CustomerID: 0, //will be set by database
	}

	customerEmail2 := &model.CustomerEmail{
		Email: "sallysmith345@gmail.com",
	}

	customerDetails2 := &model.CustomerDetails{
		FirstName:    "sally",
		LastName:     "smith",
		BirthDate:    "1982-10-13",
		PhoneNumber:  "3425534657",
		Address:      "18 main st",
		CustomerType: "Legal", //hoghoghi
		BankID:       1,
	}
	db.RegisterCustomer(customerLogin2, customerEmail2, customerDetails2)

	accountNumber2 := &model.AccountNumber{
		AccountNumber: "3424235456",
		AccountID:     0, //will be set by database
	}

	accountDetails2 := &model.AccountDetails{
		AccountType:     "Deposit",
		AccountPassword: "1234",
		Balance:         1000,
		AccountStatus:   "Closed",
		OpenDate:        "2014-03-01",
		CustomerID:      2,
	}
	db.CreateAccount(accountNumber2, accountDetails2)

	db.GetAccount("1921072918", "1111")
	db.GetAccount("3424235456", "1234")

	//to keep the app running
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Banking System Server")
	})
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
