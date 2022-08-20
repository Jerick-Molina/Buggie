package types

type Projects struct {
	Id          int    `json:"Id"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	CompanyId   string `json:"CompanyId"`
}
