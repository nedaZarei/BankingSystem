package db

import (
	"fmt"
	"log"

	"github.com/nedaZarei/BankingSystem/model"
)

func CreateBranch(branch *model.Branch) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	_, err := db.Exec("INSERT INTO branch (bank_id, address) VALUES ($1, $2)",
		branch.BankID, branch.Address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfully created branch")
}

func GetBranch(branchID int) (*model.Branch, error) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	var branch model.Branch
	err := db.QueryRow("SELECT branch_id, bank_id, address FROM branch WHERE branch_id = $1",
		branchID).Scan(&branch.BranchID, &branch.BankID, &branch.Address)
	if err != nil {
		return nil, err
	}

	return &branch, nil
}

func UpdateBranch(branch *model.Branch) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	_, err := db.Exec("UPDATE branch SET bank_id = $1, address = $2 WHERE branch_id = $3",
		branch.BankID, branch.Address, branch.BranchID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfully updated branch")
}

func DeleteBranch(branchID int) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	_, err := db.Exec("DELETE FROM branch WHERE branch_id = $1", branchID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfully deleted branch")
}
