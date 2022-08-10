package types

type User struct {
	Id        int    `json:"Id"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
	Role      string `json:"Role"`
	CompanyId int    `json:"companyId"`
}

type User_Assigned_Project struct {
	CompanyId    int `json:"CompanyId"`
	ProjectId    int `json:"ProjectId"`
	AssignedToId int `json:"AssignedId"`
}

type User_Assignesd_Ticket struct {
	CompanyId    int `json:"CompanyId"`
	ProjectId    int `json:"ProjectId"`
	TicketId     int `json:"TicketId"`
	AssignedToId int `json:"AssignedId"`
}
