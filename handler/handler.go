package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	db "github.com/nedaZarei/BankingSystem/db"
	"github.com/nedaZarei/BankingSystem/model"
)

func Register() {
	reader := bufio.NewReader(os.Stdin) //for multi-word inputs

	var account model.Account

	fmt.Print("Username: ")
	account.Username, _ = reader.ReadString('\n')
	account.Username = trimInput(account.Username)

	fmt.Print("Password: ")
	account.Password, _ = reader.ReadString('\n')
	account.Password = trimInput(account.Password)

	fmt.Print("First Name: ")
	account.First_name, _ = reader.ReadString('\n')
	account.First_name = trimInput(account.First_name)

	fmt.Print("Last Name: ")
	account.Last_name, _ = reader.ReadString('\n')
	account.Last_name = trimInput(account.Last_name)

	fmt.Print("Date of Birth (YYYY-MM-DD): ")
	account.Date_of_birth, _ = reader.ReadString('\n')
	account.Date_of_birth = trimInput(account.Date_of_birth)

	fmt.Print("National ID: ")
	account.National_id, _ = reader.ReadString('\n')
	account.National_id = trimInput(account.National_id)

	fmt.Print("Email: ")
	account.Email, _ = reader.ReadString('\n')
	account.Email = trimInput(account.Email)

	fmt.Print("Phone: ")
	account.Phone, _ = reader.ReadString('\n')
	account.Phone = trimInput(account.Phone)

	fmt.Print("Account Type: ")
	account.AccountType, _ = reader.ReadString('\n')
	account.AccountType = trimInput(account.AccountType)

	fmt.Print("Initial Balance: ")
	fmt.Scanf("%f", &account.Balance)

	//passing account struct to db for registration
	db.Register(&account)
	fmt.Println("registered successfully!")
}

func Login() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	username = trimInput(username)

	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')
	password = trimInput(password)

	if err := db.Login(username, password); err != nil {
		fmt.Println("login failed:", err)
	} else {
		fmt.Println("logged-in successfully!")
	}
}

// to trim input strings
func trimInput(input string) string {
	return strings.TrimSpace(input)
}
