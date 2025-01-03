package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/nedaZarei/BankingSystem/model"
)

func CreateAccount(number *model.AccountNumber, details *model.AccountDetails) error {
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	err = tx.QueryRow(
		"INSERT INTO account_numbers (account_number) VALUES ($1) RETURNING account_id",
		number.AccountNumber).Scan(&details.AccountID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert account number: %w", err)
	}

	_, err = tx.Exec(
		"INSERT INTO account_details (account_id, account_type, account_password, balance, account_status, open_date, close_date, customer_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		details.AccountID, details.AccountType, details.AccountPassword, details.Balance,
		details.AccountStatus, details.OpenDate, details.CloseDate, details.CustomerID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert account details: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	fmt.Println("successfully created account")
	return nil
}

func GetAccount(accountNumber, accountPassword string) (*model.AccountDetails, error) {
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	var accountID int
	err := db.QueryRow(
		"SELECT account_id FROM account_numbers WHERE account_number = $1",
		accountNumber).Scan(&accountID)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("account number not found")
	} else if err != nil {
		return nil, fmt.Errorf("failed to fetch account ID: %w", err)
	}

	var details model.AccountDetails
	err = db.QueryRow(
		"SELECT account_id, account_type, account_password, balance, account_status, open_date, close_date, customer_id FROM account_details WHERE account_id = $1",
		accountID).Scan(
		&details.AccountID, &details.AccountType, &details.AccountPassword,
		&details.Balance, &details.AccountStatus, &details.OpenDate,
		&details.CloseDate, &details.CustomerID)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("account details not found")
	} else if err != nil {
		return nil, fmt.Errorf("failed to fetch account details: %w", err)
	}

	if details.AccountPassword != accountPassword {
		return nil, fmt.Errorf("invalid account password")
	}

	return &details, nil
}

func UpdateAccount(accountID int, updates map[string]interface{}) error {
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	detailsUpdates := make(map[string]interface{})

	// Categorize updates for account_details table
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
		updateClauses := []string{}

		for key, value := range detailsUpdates {
			updateClauses = append(updateClauses, fmt.Sprintf("%s = $%d", key, paramCount))
			values = append(values, value)
			paramCount++
		}

		query += strings.Join(updateClauses, ", ")
		query += " WHERE account_id = $1"

		_, err = tx.Exec(query, values...)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update account details: %w", err)
		}
	}

	return tx.Commit()
}

func DeleteAccount(accountID int) error {
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	_, err = tx.Exec("DELETE FROM account_details WHERE account_id = $1", accountID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete account details: %w", err)
	}

	_, err = tx.Exec("DELETE FROM account_numbers WHERE account_id = $1", accountID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete account number: %w", err)
	}

	return tx.Commit()
}
