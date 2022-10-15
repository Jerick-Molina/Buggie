package db

type Account struct {
	Id        int    `json:"Id"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
	Role      string `json:"Role"`
	CompanyId int    `json:"companyId"`
}

type Company struct {
	CompanyId   int    `json:"Id"`
	CompanyName string `json:"CompanyName"`
	CompanyCode string `json:"CompanyCode"`
}
