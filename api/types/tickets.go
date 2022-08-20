package types

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
