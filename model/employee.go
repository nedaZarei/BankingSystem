package model

type Employee struct {
	EmployeeID int     `db:"employee_id"`
	Username   string  `db:"username"`
	Password   string  `db:"password"`
	FirstName  string  `db:"first_name"`
	LastName   string  `db:"last_name"`
	Position   string  `db:"position"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
	BranchID   int     `db:"branch_id"`
}
