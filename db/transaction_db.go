package db

import (
	"fmt"
	"log"
	"strings"

	"github.com/nedaZarei/BankingSystem/model"
)

func CreateTransaction(transaction *model.Transaction) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err := db.Exec(
		"INSERT INTO transaction (source_account_id, destination_account_id, amount, transaction_type, transaction_date) VALUES ($1, $2, $3, $4, $5)",
		transaction.SourceAccountID, transaction.DestinationAccountID,
		transaction.Amount, transaction.TransactionType, transaction.TransactionDate)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully created transaction")
}

func GetTransaction(transactionID int) (*model.Transaction, error) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	var transaction model.Transaction
	err := db.QueryRow(
		"SELECT transaction_id, source_account_id, destination_account_id, amount, transaction_type, transaction_date FROM transaction WHERE transaction_id = $1",
		transactionID).Scan(&transaction.TransactionID, &transaction.SourceAccountID,
		&transaction.DestinationAccountID, &transaction.Amount, &transaction.TransactionType, &transaction.TransactionDate)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func UpdateTransaction(transactionID int, updates map[string]interface{}) error {
	if err := db.Ping(); err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	transactionUpdates := make(map[string]interface{})

	for key, value := range updates {
		switch key {
		case "source_account_id", "destination_account_id", "amount", "transaction_type", "transaction_date":
			transactionUpdates[key] = value
		}
	}

	if len(transactionUpdates) > 0 {
		query := "UPDATE transaction SET "
		values := []interface{}{transactionID}
		paramCount := 2
		updatesArr := []string{}

		//building query with parameterized updates
		for key, value := range transactionUpdates {
			updatesArr = append(updatesArr, fmt.Sprintf("%s = $%d", key, paramCount))
			values = append(values, value)
			paramCount++
		}

		query += strings.Join(updatesArr, ", ")
		query += " WHERE transaction_id = $1"

		_, err = tx.Exec(query, values...)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update transaction: %v", err)
		}
	}

	return tx.Commit()
}

func DeleteTransaction(transactionID int) error {
	if err := db.Ping(); err != nil {
		return err
	}

	_, err := db.Exec("DELETE FROM transaction WHERE transaction_id = $1", transactionID)
	if err != nil {
		return err
	}

	return nil
}
