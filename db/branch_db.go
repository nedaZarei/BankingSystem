package db

import (
	"database/sql"
	"fmt"

	"github.com/nedaZarei/BankingSystem/model"
)

func CreateBranch(branch *model.Branch) error {
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	query := "INSERT INTO branch (bank_id, address) VALUES ($1, $2)"
	_, err := db.Exec(query, branch.BankID, branch.Address)
	if err != nil {
		return fmt.Errorf("failed to create branch: %w", err)
	}

	fmt.Println("successfully created branch")
	return nil
}

func GetBranch(branchID int) (*model.Branch, error) {
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	var branch model.Branch
	query := "SELECT branch_id, bank_id, address FROM branch WHERE branch_id = $1"
	err := db.QueryRow(query, branchID).Scan(&branch.BranchID, &branch.BankID, &branch.Address)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("branch with ID %d not found", branchID)
	} else if err != nil {
		return nil, fmt.Errorf("failed to fetch branch: %w", err)
	}

	return &branch, nil
}

func UpdateBranch(branch *model.Branch) error {
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	query := "UPDATE branch SET bank_id = $1, address = $2 WHERE branch_id = $3"
	_, err := db.Exec(query, branch.BankID, branch.Address, branch.BranchID)
	if err != nil {
		return fmt.Errorf("failed to update branch: %w", err)
	}

	fmt.Println("successfully updated branch")
	return nil
}

func DeleteBranch(branchID int) error {
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	query := "DELETE FROM branch WHERE branch_id = $1"
	_, err := db.Exec(query, branchID)
	if err != nil {
		return fmt.Errorf("failed to delete branch: %w", err)
	}

	fmt.Println("successfully deleted branch")
	return nil
}
