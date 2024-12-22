package model

type EmployeeLogin struct {
	Username   string `db:"username"`
	Password   string `db:"password"`
	EmployeeID int    `db:"employee_id"`
}

type EmployeeDetails struct {
	EmployeeID int     `db:"employee_id"`
	FirstName  string  `db:"first_name"`
	LastName   string  `db:"last_name"`
	Position   string  `db:"position"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
	BranchID   int     `db:"branch_id"`
}
