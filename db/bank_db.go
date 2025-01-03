package db

import (
	"database/sql"
	"fmt"

	"github.com/nedaZarei/BankingSystem/model"
)

func CreateBank(bank *model.Bank) error {
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	query := "INSERT INTO bank (name, headquarter_address) VALUES ($1, $2)"
	_, err := db.Exec(query, bank.Name, bank.HeadquarterAddress)
	if err != nil {
		return fmt.Errorf("failed to insert bank: %w", err)
	}

	fmt.Println("Successfully created bank")
	return nil
}

func GetBank(bankID int) (*model.Bank, error) {
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	var bank model.Bank
	query := "SELECT bank_id, name, headquarter_address FROM bank WHERE bank_id = $1"
	err := db.QueryRow(query, bankID).Scan(&bank.BankID, &bank.Name, &bank.HeadquarterAddress)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("bank with ID %d not found", bankID)
	} else if err != nil {
		return nil, fmt.Errorf("failed to fetch bank: %w", err)
	}

	return &bank, nil
}

func UpdateBank(bank *model.Bank) error {
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	query := "UPDATE bank SET name = $1, headquarter_address = $2 WHERE bank_id = $3"
	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(bank.Name, bank.HeadquarterAddress, bank.BankID)
	if err != nil {
		return fmt.Errorf("failed to update bank: %w", err)
	}

	fmt.Println("Successfully updated bank")
	return nil
}

func DeleteBank(bankID int) error {
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	query := "DELETE FROM bank WHERE bank_id = $1"
	_, err := db.Exec(query, bankID)
	if err != nil {
		return fmt.Errorf("failed to delete bank: %w", err)
	}

	fmt.Println("Successfully deleted bank")
	return nil
}
