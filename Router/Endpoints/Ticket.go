package endpoints

import (
	"net/http"

	properties "github.com/Jerick-Molina/Buggie/Properties"
	database "github.com/Jerick-Molina/Buggie/src/Database"
)

func TicketCreate() {
	//sqlExc := "insert into Ticket "
}

func TicketFindByCompany(companyId int) ([]properties.Tickets, int, string) {
	db := database.Access()
	defer db.Close()
	var tickets []properties.Tickets
	sqlQuery := `select TicketId,Name,Description,Status,AssignedTo,Priority,
	ProjectId from Ticket where CompanyId = ?`

	rows, err := db.Query(sqlQuery, companyId)
	if err != nil {
		return nil, http.StatusBadRequest, "Invalid request"
	}

	for rows.Next() {
		var tkt properties.Tickets
		if err := rows.Scan(&tkt.Id, &tkt.Name, &tkt.Description, &tkt.Status,
			&tkt.AssignedTo, &tkt.Priority, &tkt.ProjectId); err != nil {
			return nil, http.StatusInternalServerError, "Something happened in our end"
		}
		tickets = append(tickets, tkt)
	}
	if err := rows.Err(); err != nil {
		return nil, http.StatusInternalServerError, "Something happened in our end"
	}
	return tickets, http.StatusOK, "Valid"
}
func TicketsFindByProject(projectId int, companyId int) ([]properties.Tickets, int, string) {
	db := database.Access()
	defer db.Close()
	var tickets []properties.Tickets
	sqlQuery := `select TicketId,Name,Description,Status,AssignedTo,Priority,
	ProjectId from Ticket where ProjectId = ? and CompanyId = ?`

	rows, err := db.Query(sqlQuery, projectId, companyId)
	if err != nil {
		return nil, http.StatusBadRequest, "Invalid request"
	}

	for rows.Next() {
		var tkt properties.Tickets
		if err := rows.Scan(&tkt.Id, &tkt.Name, &tkt.Description, &tkt.Status,
			&tkt.AssignedTo, &tkt.Priority, &tkt.ProjectId); err != nil {
			return nil, http.StatusInternalServerError, "Something happened in our end"
		}
		tickets = append(tickets, tkt)
	}
	if err := rows.Err(); err != nil {
		return nil, http.StatusInternalServerError, "Something happened in our end"
	}
	return tickets, http.StatusOK, "Valid"
}

func TicketEdit() {

}
