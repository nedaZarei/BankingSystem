package db

import (
	"fmt"
	"log"
	"strings"

	"github.com/nedaZarei/BankingSystem/model"
)

func CreateLoan(loan *model.Loan) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err := db.Exec(
		"INSERT INTO loan (customer_id, loan_type, amount, interest_rate, duration, start_date, end_date, loan_status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		loan.CustomerID, loan.LoanType, loan.Amount, loan.InterestRate, loan.Duration, loan.StartDate, loan.EndDate, loan.LoanStatus)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("successfully created loan")
}

func GetLoan(loanID int) (*model.Loan, error) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	var loan model.Loan
	err := db.QueryRow(
		"SELECT loan_id, customer_id, loan_type, amount, interest_rate, duration, start_date, end_date, loan_status FROM loan WHERE loan_id = $1",
		loanID).Scan(&loan.LoanID, &loan.CustomerID, &loan.LoanType, &loan.Amount, &loan.InterestRate, &loan.Duration, &loan.StartDate, &loan.EndDate, &loan.LoanStatus)
	if err != nil {
		return nil, err
	}

	return &loan, nil
}

func UpdateLoan(loanID int, updates map[string]interface{}) error {
	if err := db.Ping(); err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	loanUpdates := make(map[string]interface{})

	//sorting updates into appropriate maps
	for key, value := range updates {
		switch key {
		case "customer_id", "loan_type", "amount", "interest_rate", "duration", "start_date", "end_date", "loan_status":
			loanUpdates[key] = value
		}
	}

	if len(loanUpdates) > 0 {
		query := "UPDATE loan SET "
		values := []interface{}{loanID}
		paramCount := 2
		updates_arr := []string{}

		for key, value := range loanUpdates {
			updates_arr = append(updates_arr, fmt.Sprintf("%s = $%d", key, paramCount))
			values = append(values, value)
			paramCount++
		}

		query += strings.Join(updates_arr, ", ")
		query += " WHERE loan_id = $1"

		_, err = tx.Exec(query, values...)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update loan: %v", err)
		}
	}

	return tx.Commit()
}

func DeleteLoan(loanID int) error {
	if err := db.Ping(); err != nil {
		return err
	}

	_, err := db.Exec("DELETE FROM loan WHERE loan_id = $1", loanID)
	if err != nil {
		return err
	}

	return nil
}
