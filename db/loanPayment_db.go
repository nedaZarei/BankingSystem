package db

import (
	"fmt"
	"log"
	"strings"

	"github.com/nedaZarei/BankingSystem/model"
)

func CreateLoanPayment(payment *model.LoanPayment) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err := db.Exec(
		"INSERT INTO loan_payment (loan_id, payment_amount, due_date, payment_date, payment_status) VALUES ($1, $2, $3, $4, $5)",
		payment.LoanID, payment.PaymentAmount, payment.DueDate, payment.PaymentDate, payment.PaymentStatus)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("successfully created loan payment")
}

func GetLoanPayment(paymentID int) (*model.LoanPayment, error) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	var payment model.LoanPayment
	err := db.QueryRow(
		"SELECT payment_id, loan_id, payment_amount, due_date, payment_date, payment_status FROM loan_payment WHERE payment_id = $1",
		paymentID).Scan(&payment.PaymentID, &payment.LoanID, &payment.PaymentAmount, &payment.DueDate, &payment.PaymentDate, &payment.PaymentStatus)
	if err != nil {
		return nil, err
	}

	return &payment, nil
}

func UpdateLoanPayment(paymentID int, updates map[string]interface{}) error {
	if err := db.Ping(); err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	paymentUpdates := make(map[string]interface{})

	//sorting updates into appropriate maps
	for key, value := range updates {
		switch key {
		case "loan_id", "payment_amount", "due_date", "payment_date", "payment_status":
			paymentUpdates[key] = value
		}
	}

	if len(paymentUpdates) > 0 {
		query := "UPDATE loan_payment SET "
		values := []interface{}{paymentID}
		paramCount := 2
		updates_arr := []string{}

		for key, value := range paymentUpdates {
			updates_arr = append(updates_arr, fmt.Sprintf("%s = $%d", key, paramCount))
			values = append(values, value)
			paramCount++
		}

		query += strings.Join(updates_arr, ", ")
		query += " WHERE payment_id = $1"

		_, err = tx.Exec(query, values...)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update loan payment: %v", err)
		}
	}

	return tx.Commit()
}

func DeleteLoanPayment(paymentID int) error {
	if err := db.Ping(); err != nil {
		return err
	}

	_, err := db.Exec("DELETE FROM loan_payment WHERE payment_id = $1", paymentID)
	if err != nil {
		return err
	}

	return nil
}
