package db

import (
	"fmt"
	"log"
	"strings"

	"github.com/nedaZarei/BankingSystem/model"
)

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

func UpdateAccount(accountID int, updates map[string]interface{}) error {
	if err := db.Ping(); err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	detailsUpdates := make(map[string]interface{})

	//sorting updates into appropriate maps
	for key, value := range updates {
		switch key {
		case "account_type", "account_password", "balance", "account_status", "open_date", "close_date", "customer_id":
			detailsUpdates[key] = value
		}
	}

	if len(detailsUpdates) > 0 {
		query := "UPDATE account_details SET "
		values := []interface{}{accountID}
		paramCount := 2
		updates_arr := []string{}

		for key, value := range detailsUpdates {
			updates_arr = append(updates_arr, fmt.Sprintf("%s = $%d", key, paramCount))
			values = append(values, value)
			paramCount++
		}

		query += strings.Join(updates_arr, ", ")
		query += " WHERE account_id = $1"

		_, err = tx.Exec(query, values...)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update account details: %v", err)
		}
	}

	return tx.Commit()
}

func DeleteAccount(accountID int) error {
	if err := db.Ping(); err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM account_details WHERE account_id = $1", accountID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete account details: %v", err)
	}

	_, err = tx.Exec("DELETE FROM account_numbers WHERE account_id = $1", accountID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete account number: %v", err)
	}

	return tx.Commit()
}
