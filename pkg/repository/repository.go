package repository

import (
	_ "embed"
	"fmt"
	"github.com/jmoiron/sqlx"
	models "go-rest-employee"
	"strings"
)

type EmployeeRepository interface {
	Create(employee models.CreateEmployeeRequest) (int, error)
	GetByCompany(companyId int) ([]models.EmployeeResponse, error)
	GetByDepartment(departmentName string) ([]models.EmployeeResponse, error)
	Update(employeeId int, input models.UpdateEmployeeInput) error
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

func (r *Repository) Create(employee models.CreateEmployeeRequest) (int, error) {
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

func (r *Repository) GetByCompany(companyId int) ([]models.EmployeeResponse, error) {

	var employees []models.EmployeeResponse

	err := r.db.Select(&employees, getByCompany, companyId)

	return employees, err
}

//go:embed sql/GetByDepartment.sql
var getByDepartment string

func (r *Repository) GetByDepartment(departmentName string) ([]models.EmployeeResponse, error) {
	var employees []models.EmployeeResponse

	err := r.db.Select(&employees, getByDepartment, departmentName)

	return employees, err
}

//go:embed sql/UpdateEmployee.sql
var updateEmployee string

func (r *Repository) Update(employeeId int, input models.UpdateEmployeeInput) error {
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

	// Build the query for updating employees
	var employeeUpdates []string
	var args []interface{}
	argID := 1

	if input.Name != nil {
		employeeUpdates = append(employeeUpdates, fmt.Sprintf("name = $%d", argID))
		args = append(args, *input.Name)
		argID++
	}
	if input.Surname != nil {
		employeeUpdates = append(employeeUpdates, fmt.Sprintf("surname = $%d", argID))
		args = append(args, *input.Surname)
		argID++
	}
	if input.Phone != nil {
		employeeUpdates = append(employeeUpdates, fmt.Sprintf("phone = $%d", argID))
		args = append(args, *input.Phone)
		argID++
	}
	if input.CompanyId != nil {
		employeeUpdates = append(employeeUpdates, fmt.Sprintf("company_id = $%d", argID))
		args = append(args, *input.CompanyId)
		argID++
	}
	if input.DepartmentId != nil {
		employeeUpdates = append(employeeUpdates, fmt.Sprintf("department_id = $%d", argID))
		args = append(args, *input.DepartmentId)
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

	// Update passport if provided
	if input.Passport != nil {
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

	// Update department if provided
	if input.Department != nil {
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

	/*updateFields := []string{}
	args := []interface{}{}
	argId := 1

	if input.Name != nil {
		updateFields = append(updateFields, fmt.Sprintf("name = $%d", argId))
		args = append(args, *input.Name)
		argId++
	}
	if input.Surname != nil {
		updateFields = append(updateFields, fmt.Sprintf("surname = $%d", argId))
		args = append(args, *input.Surname)
		argId++
	}
	if input.Phone != nil {
		updateFields = append(updateFields, fmt.Sprintf("phone = $%d", argId))
		args = append(args, *input.Phone)
		argId++
	}
	if input.CompanyId != nil {
		updateFields = append(updateFields, fmt.Sprintf("company_id = $%d", argId))
		args = append(args, *input.CompanyId)
		argId++
	}

	if len(updateFields) > 0 {
		args = append(args, employeeId)
		query := fmt.Sprintf("UPDATE employees SET %s WHERE id = $%d", strings.Join(updateFields, ", "), argId)
		_, err = tx.Exec(query, args...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if input.Passport != nil {
		passportUpdateFields := []string{}
		passportArgs := []interface{}{}
		passportArgId := 1

		if input.Passport.Type != nil {
			passportUpdateFields = append(passportUpdateFields, fmt.Sprintf("type = $%d", passportArgId))
			passportArgs = append(passportArgs, *input.Passport.Type)
			passportArgId++
		}
		if input.Passport.Number != nil {
			passportUpdateFields = append(passportUpdateFields, fmt.Sprintf("number = $%d", passportArgId))
			passportArgs = append(passportArgs, *input.Passport.Number)
			passportArgId++
		}

		if len(passportUpdateFields) > 0 {
			passportArgs = append(passportArgs, employeeId)
			passportQuery := fmt.Sprintf("UPDATE passports SET %s WHERE id = (SELECT passport_id FROM employees WHERE id = $%d)", strings.Join(passportUpdateFields, ", "), passportArgId)
			_, err = tx.Exec(passportQuery, passportArgs...)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	if input.Department != nil {
		var departmentId int
		err := tx.Get(&departmentId, `SELECT department_id FROM employees_departments WHERE employee_id = $1`, employeeId)
		if err != nil {
			tx.Rollback()
			return err
		}

		departmentUpdateFields := []string{}
		departmentArgs := []interface{}{}
		departmentArgId := 1

		if input.Department.Name != nil {
			departmentUpdateFields = append(departmentUpdateFields, fmt.Sprintf("name = $%d", departmentArgId))
			departmentArgs = append(departmentArgs, *input.Department.Name)
			departmentArgId++
		}
		if input.Department.Phone != nil {
			departmentUpdateFields = append(departmentUpdateFields, fmt.Sprintf("phone = $%d", departmentArgId))
			departmentArgs = append(departmentArgs, *input.Department.Phone)
			departmentArgId
		}
		if len(departmentUpdateFields) > 0 {
			departmentArgs = append(departmentArgs, departmentId)
			departmentQuery := fmt.Sprintf("UPDATE departments SET %s WHERE id = $%d", strings.Join(departmentUpdateFields, ", "), departmentArgId)
			_, err = tx.Exec(departmentQuery, departmentArgs...)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}*/

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
