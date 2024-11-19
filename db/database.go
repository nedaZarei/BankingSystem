package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/nedaZarei/BankingSystem/model"
)

var db *sql.DB

const (
	host     = "postgres"
	port     = 5432
	user     = "neda.z"
	password = "nz2003nz"
	dbname   = "banking_system"
)

func Connect() error {
	var err error
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	return db.Ping()
}

func Register(account *model.Account) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	query := `INSERT INTO account (username, password, firstName, lastName, dateOfBirth, nationalID, email, phoneNumber, accountType, balance) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := db.Exec(query, account.Username, account.Password, account.First_name, account.Last_name, account.Date_of_birth, account.National_id, account.Email, account.Phone, account.AccountType, account.Balance)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfully registered")
}
