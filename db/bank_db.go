package db

import (
	"fmt"
	"log"

	"github.com/nedaZarei/BankingSystem/model"
)

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

func UpdateBank(bank *model.Bank) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	_, err := db.Exec("UPDATE bank SET name = $1, headquarter_address = $2 WHERE bank_id = $3",
		bank.Name, bank.HeadquarterAddress, bank.BankID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfully updated bank")
}

func GetBank(bankID int) (*model.Bank, error) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	var bank model.Bank
	err := db.QueryRow("SELECT bank_id, name, headquarter_address FROM bank WHERE bank_id = $1",
		bankID).Scan(&bank.BankID, &bank.Name, &bank.HeadquarterAddress)
	if err != nil {
		return nil, err
	}

	return &bank, nil
}

func DeleteBank(bankID int) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	_, err := db.Exec("DELETE FROM bank WHERE bank_id = $1", bankID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfully deleted bank")
}
