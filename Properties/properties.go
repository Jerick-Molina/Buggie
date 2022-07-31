package properties

type User struct {
	Id        int    `json:"Id"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
	Role      string `json:"Role"`
	CompanyId int    `json:"companyId"`
}

type Company struct {
	Id          int    `json:"Id"`
	Name        string `json:"Name"`
	CompanyCode string `json:"companyCode"`
}

type Projects struct {
	Id          int    `json:"Id"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	CompanyId   string `json:"CompanyId"`
}

type Tickets struct {
	Id          int `json:"CompanyId"`
	Name        string
	Description string
	Status      string
	AssignedTo  string
	//	Date  time TODO: Figure datetime out
	Priority  string
	CompanyId string
	ProjectId string
}

type Comments struct {
	CompanyId int    `json:"CompanyId"`
	ProjectId int    `json:"ProjectId"`
	TicketId  int    `json:"TicketId"`
	UserIdId  int    `json:"UserIdId"`
	Message   string `json:"Message"`
}
