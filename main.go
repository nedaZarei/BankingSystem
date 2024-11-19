package main

import (
	"fmt"
	"log"

	db "github.com/nedaZarei/BankingSystem/db"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to database")
}
