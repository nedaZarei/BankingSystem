package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/nedaZarei/BankingSystem/model"
)

var db *sqlx.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "neda.z"
	password = "nz2003nz"
	dbname   = "banking_system"
)

func Connect() error {
	var err error
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	return err
}

func Register(account *model.Account) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	query := `INSERT INTO accounts (username, password, first_name, last_name, date_of_birth, national_id, email, phone, account_type, balance) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	if _, err := db.Exec(query, account.Username, account.Password, account.First_name, account.Last_name, account.Date_of_birth, account.National_id, account.Email, account.Phone, account.AccountType, account.Balance); err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfully registered")
}

func Login(username string, password string) *model.Account {
	account := &model.Account{}
	query := `SELECT * FROM accounts WHERE username=$1 AND password=$2`
	if err := db.Get(account, query, username, password); err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("successfully logged-in")
	return nil
}
