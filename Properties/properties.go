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
	Id          int    `json:"Id"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Status      string `json:"Status"`
	AssignedTo  int    `json:"AssignedTo"`
	CreatedBy   int    `json:"CreatedBy"`
	//	Date  time TODO: Figure datetime out
	Priority  string `json:"Priority"`
	CompanyId int    `json:"CompanyId"`
	ProjectId int    `json:"ProjectId"`
}

type Comments struct {
	CompanyId int    `json:"CompanyId"`
	ProjectId int    `json:"ProjectId"`
	TicketId  int    `json:"TicketId"`
	UserId    int    `json:"UserIdId"`
	Message   string `json:"Message"`
}

type AssignedProject struct {
	CompanyId    int `json:"CompanyId"`
	ProjectId    int `json:"ProjectId"`
	AssignedToId int `json:"AssignedTo"`
}

type AssignedTicket struct {
	CompanyId    int `json:"CompanyId"`
	ProjectId    int `json:"ProjectId"`
	TicketId     int `json:"TicketId"`
	AssignedToId int `json:"AssignedTo"`
}
