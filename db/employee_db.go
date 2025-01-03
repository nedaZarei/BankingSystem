package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/nedaZarei/BankingSystem/model"
	"golang.org/x/crypto/bcrypt"
)

func RegisterEmployee(login *model.EmployeeLogin, details *model.EmployeeDetails) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(login.Password), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		log.Fatal("failed to hash password: ", err)
	}

	err = tx.QueryRow(
		"INSERT INTO employee_login (username, password) VALUES ($1, $2) RETURNING employee_id",
		login.Username, hashedPassword).Scan(&details.EmployeeID)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	_, err = tx.Exec(
		"INSERT INTO employee_details (employee_id, first_name, last_name, position, department, salary, branch_id) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		details.EmployeeID, details.FirstName, details.LastName, details.Position, details.Department, details.Salary, details.BranchID)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully registered employee")
}

func LoginEmployee(username, password string) (*model.EmployeeDetails, error) {
	if err := db.Ping(); err != nil {
		return nil, err
	}
	var hashedPassword string
	var employeeID int

	//retrieving the hashed password and employee ID
	err := db.QueryRow(
		"SELECT password, employee_id FROM employee_login WHERE username = $1",
		username).Scan(&hashedPassword, &employeeID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid username or password")
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return nil, errors.New("invalid username or password")
	}

	var details model.EmployeeDetails

	err = db.QueryRow(
		"SELECT employee_id, first_name, last_name, position, department, salary, branch_id FROM employee_details WHERE employee_id = $1",
		employeeID).Scan(
		&details.EmployeeID, &details.FirstName, &details.LastName,
		&details.Position, &details.Department, &details.Salary, &details.BranchID)
	if err != nil {
		return nil, err
	}

	return &details, nil
}

func UpdateEmployee(employeeID int, updates map[string]interface{}) error {
	if err := db.Ping(); err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	loginUpdates := make(map[string]interface{})
	detailsUpdates := make(map[string]interface{})

	for key, value := range updates {
		switch key {
		case "username":
			loginUpdates[key] = value
		case "password":
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(value.(string)), bcrypt.DefaultCost)
			if err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to hash password: %v", err)
			}
			loginUpdates[key] = hashedPassword
		case "first_name", "last_name", "position", "department", "salary", "branch_id":
			detailsUpdates[key] = value
		}
	}

	if len(loginUpdates) > 0 {
		query := "UPDATE employee_login SET "
		values := []interface{}{employeeID}
		paramCount := 2
		updatesArr := []string{}

		for key, value := range loginUpdates {
			updatesArr = append(updatesArr, fmt.Sprintf("%s = $%d", key, paramCount))
			values = append(values, value)
			paramCount++
		}

		query += strings.Join(updatesArr, ", ")
		query += " WHERE employee_id = $1"

		_, err = tx.Exec(query, values...)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update employee login: %v", err)
		}
	}

	if len(detailsUpdates) > 0 {
		query := "UPDATE employee_details SET "
		values := []interface{}{employeeID}
		paramCount := 2
		updatesArr := []string{}

		for key, value := range detailsUpdates {
			updatesArr = append(updatesArr, fmt.Sprintf("%s = $%d", key, paramCount))
			values = append(values, value)
			paramCount++
		}

		query += strings.Join(updatesArr, ", ")
		query += " WHERE employee_id = $1"

		_, err = tx.Exec(query, values...)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update employee details: %v", err)
		}
	}

	return tx.Commit()
}

func GetEmployeeByID(employeeID int) (*model.EmployeeLogin, *model.EmployeeDetails, error) {
	if err := db.Ping(); err != nil {
		return nil, nil, err
	}

	login := &model.EmployeeLogin{}
	details := &model.EmployeeDetails{}

	err := db.QueryRow(
		`SELECT username, password, employee_id 
         FROM employee_login 
         WHERE employee_id = $1`,
		employeeID).Scan(&login.Username, &login.Password, &login.EmployeeID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, errors.New("employee not found")
		}
		return nil, nil, err
	}

	err = db.QueryRow(
		`SELECT employee_id, first_name, last_name, position, 
                department, salary, branch_id 
         FROM employee_details 
         WHERE employee_id = $1`,
		employeeID).Scan(
		&details.EmployeeID, &details.FirstName, &details.LastName,
		&details.Position, &details.Department, &details.Salary, &details.BranchID)
	if err != nil {
		return nil, nil, err
	}

	return login, details, nil
}

func GetEmployeesByBranch(branchID int) ([]model.EmployeeDetails, error) {
	if err := db.Ping(); err != nil {
		return nil, err
	}

	rows, err := db.Query(
		`SELECT employee_id, first_name, last_name, position, 
                department, salary, branch_id 
         FROM employee_details 
         WHERE branch_id = $1`,
		branchID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []model.EmployeeDetails
	for rows.Next() {
		var emp model.EmployeeDetails
		err := rows.Scan(
			&emp.EmployeeID, &emp.FirstName, &emp.LastName,
			&emp.Position, &emp.Department, &emp.Salary, &emp.BranchID)
		if err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}

	return employees, nil
}

func DeleteEmployee(employeeID int) error {
	if err := db.Ping(); err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	//deleting from employee_details first due to foreign key constraint
	_, err = tx.Exec("DELETE FROM employee_details WHERE employee_id = $1", employeeID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete employee details: %v", err)
	}

	_, err = tx.Exec("DELETE FROM employee_login WHERE employee_id = $1", employeeID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete employee login: %v", err)
	}

	return tx.Commit()
}
