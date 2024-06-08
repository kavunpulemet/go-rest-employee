package repository

import (
	_ "embed"
	"fmt"
	"github.com/jmoiron/sqlx"
	repository "go-rest-employee/pkg/repository/models"
	"strings"
)

type EmployeeRepository interface {
	Create(employee repository.CreateEmployee) (int, error)
	GetByCompany(companyId int) ([]repository.EmployeeResponse, error)
	GetByDepartment(departmentName string) ([]repository.EmployeeResponse, error)
	Update(employeeId int, input repository.UpdateEmployee) error
	Delete(employeeId int) error
}

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

//go:embed sql/CreatePassport.sql
var createPassport string

//go:embed sql/CreateDepartment.sql
var createDepartment string

//go:embed sql/CreateEmployee.sql
var createEmployee string

func (r *Repository) Create(employee repository.CreateEmployee) (int, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	var passportId int
	err = tx.Get(&passportId, createPassport, employee.Passport.Type, employee.Passport.Number)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var departmentId int
	err = tx.Get(&departmentId, createDepartment, employee.Department.Name, employee.Department.Phone)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var id int
	err = tx.Get(&id, createEmployee, employee.Name, employee.Surname, employee.Phone, employee.CompanyId, passportId, departmentId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

//go:embed sql/GetByCompany.sql
var getByCompany string

func (r *Repository) GetByCompany(companyId int) ([]repository.EmployeeResponse, error) {

	var employees []repository.EmployeeResponse

	err := r.db.Select(&employees, getByCompany, companyId)

	return employees, err
}

//go:embed sql/GetByDepartment.sql
var getByDepartment string

func (r *Repository) GetByDepartment(departmentName string) ([]repository.EmployeeResponse, error) {
	var employees []repository.EmployeeResponse

	err := r.db.Select(&employees, getByDepartment, departmentName)

	return employees, err
}

func (r *Repository) Update(employeeId int, input repository.UpdateEmployee) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	fmt.Println(input.Name)
	fmt.Println(&input.Name)
	fmt.Println(input.CompanyId)
	fmt.Println(&input.CompanyId)
	fmt.Println(input.Passport)
	fmt.Println(&input.Passport)

	var employeeUpdates []string
	var args []interface{}
	argID := 1

	if input.Name != "" {
		employeeUpdates = append(employeeUpdates, fmt.Sprintf("name = $%d", argID))
		args = append(args, input.Name)
		argID++
	}
	if input.Surname != "" {
		employeeUpdates = append(employeeUpdates, fmt.Sprintf("surname = $%d", argID))
		args = append(args, input.Surname)
		argID++
	}
	if input.Phone != "" {
		employeeUpdates = append(employeeUpdates, fmt.Sprintf("phone = $%d", argID))
		args = append(args, input.Phone)
		argID++
	}
	if input.CompanyId != 0 {
		employeeUpdates = append(employeeUpdates, fmt.Sprintf("company_id = $%d", argID))
		args = append(args, input.CompanyId)
		argID++
	}
	if input.DepartmentId != 0 {
		employeeUpdates = append(employeeUpdates, fmt.Sprintf("department_id = $%d", argID))
		args = append(args, input.DepartmentId)
		argID++
	}

	if len(employeeUpdates) > 0 {
		query := fmt.Sprintf("UPDATE employees SET %s WHERE id = $%d",
			strings.Join(employeeUpdates, ", "), argID)
		args = append(args, employeeId)
		if _, err = tx.Exec(query, args...); err != nil {
			return err
		}
	}

	if input.Passport.Type != "" && input.Passport.Number != "" {
		passport := input.Passport
		var passportUpdates []string
		args = []interface{}{}
		argID = 1

		if passport.Type != "" {
			passportUpdates = append(passportUpdates, fmt.Sprintf("type = $%d", argID))
			args = append(args, passport.Type)
			argID++
		}
		if passport.Number != "" {
			passportUpdates = append(passportUpdates, fmt.Sprintf("number = $%d", argID))
			args = append(args, passport.Number)
			argID++
		}

		if len(passportUpdates) > 0 {
			query := fmt.Sprintf("UPDATE passports SET %s WHERE id = (SELECT passport_id FROM employees WHERE id = $%d)",
				strings.Join(passportUpdates, ", "), argID)
			args = append(args, employeeId)
			if _, err = tx.Exec(query, args...); err != nil {
				return err
			}
		}
	}

	if input.Department.Name != "" && input.Department.Phone != "" {
		department := input.Department
		var departmentUpdates []string
		args = []interface{}{}
		argID = 1

		if department.Name != "" {
			departmentUpdates = append(departmentUpdates, fmt.Sprintf("name = $%d", argID))
			args = append(args, department.Name)
			argID++
		}
		if department.Phone != "" {
			departmentUpdates = append(departmentUpdates, fmt.Sprintf("phone = $%d", argID))
			args = append(args, department.Phone)
			argID++
		}

		if len(departmentUpdates) > 0 {
			query := fmt.Sprintf("UPDATE departments SET %s WHERE id = (SELECT department_id FROM employees WHERE id = $%d)",
				strings.Join(departmentUpdates, ", "), argID)
			args = append(args, employeeId)
			if _, err = tx.Exec(query, args...); err != nil {
				return err
			}
		}
	}

	return nil
}

//go:embed sql/GetPassportId.sql
var getPassportId string

//go:embed sql/DeleteEmployee.sql
var deleteEmployee string

//go:embed sql/DeletePassport.sql
var deletePassport string

func (r *Repository) Delete(employeeId int) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	var passportId int
	err = tx.Get(&passportId, getPassportId, employeeId)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(deleteEmployee, employeeId)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(deletePassport, passportId)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
