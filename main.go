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

	//example
	newAccount := &model.Account{
		Username:      "roryG",
		Password:      "securepassword",
		First_name:    "rory",
		Last_name:     "geller",
		Date_of_birth: "2000-04-06",
		National_id:   "1234567890",
		Email:         "rorygeller@gmail.com",
		Phone:         "09135673249",
		AccountType:   "savings",
		Balance:       1000.00,
	}
	db.Register(newAccount)

	// a simple HTTP server to keep the application running
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Banking System Server")
	})

	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
