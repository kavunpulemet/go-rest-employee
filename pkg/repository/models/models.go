package repository

type CreateEmployee struct {
	Name       string
	Surname    string
	Phone      string
	CompanyId  int
	Passport   Passport
	Department Department
}

type Passport struct {
	Id     int    `db:"passport_id"`
	Type   string `db:"passport_type"`
	Number string `db:"passport_number"`
}

type Department struct {
	Id    int    `db:"department_id"`
	Name  string `db:"department_name"`
	Phone string `db:"department_phone"`
}

type EmployeeResponse struct {
	Id        int    `db:"id"`
	Name      string `db:"name"`
	Surname   string `db:"surname"`
	Phone     string `db:"phone"`
	CompanyId int    `db:"company_id"`
	Passport
	Department
}

type UpdateEmployee struct {
	Name         string
	Surname      string
	Phone        string
	CompanyId    int
	DepartmentId int
	Passport     Passport
	Department   Department
}
