package models

type Employee struct {
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	Surname    string     `json:"surname"`
	Phone      string     `json:"phone"`
	CompanyId  int        `json:"company_id"`
	Passport   Passport   `json:"passport"`
	Department Department `json:"department"`
}

type Passport struct {
	Id     int    `json:"passport_id"`
	Type   string `json:"passport_type"`
	Number string `json:"passport_number"`
}

type Department struct {
	Id    int    `json:"department_id"`
	Name  string `json:"department_name"`
	Phone string `json:"department_phone"`
}
