package access

import (
	"fmt"
	"net/http"

	properties "github.com/Jerick-Molina/Buggie/Properties"
	database "github.com/Jerick-Molina/Buggie/src/Database"
)

func CreateTicket(companyId int, creatorId int, ticket *properties.Tickets) (int, string) {

	db := database.Access()

	sqlExec := `insert into Tickets
	 (Name,AssignedTo,CreatedById,CompanyId,Description,ProjectId,Priority)
	 values(?,?,?,?,?,?,?)`

	result, err := db.Exec(sqlExec, ticket.Name, ticket.AssignedTo,
		creatorId, companyId, ticket.Description, ticket.ProjectId, ticket.Priority)
	if err != nil {
		fmt.Println(err.Error())
		return http.StatusInternalServerError, "something happened in out end"
	}
	fmt.Println(result)

	return http.StatusOK, "Valid"
}

func EditTicket() {

}

func TicketFindByCompany(companyId int) ([]properties.Tickets, int, string) {
	db := database.Access()

	var tickets []properties.Tickets

	sqlQuery := `select TicketId,Name,Description,Status,AssignedTo,Priority,
	ProjectId from Tickets where CompanyId = ?`

	rows, err := db.Query(sqlQuery, companyId)
	if err != nil {
		fmt.Println(err.Error())
		return nil, http.StatusBadRequest, "Invalid request"
	}

	for rows.Next() {
		var tkt properties.Tickets
		if err := rows.Scan(&tkt.Id, &tkt.Name, &tkt.Description, &tkt.Status,
			&tkt.AssignedTo, &tkt.Priority, &tkt.ProjectId); err != nil {
			fmt.Println(companyId)
			fmt.Println(err.Error())
			return nil, http.StatusInternalServerError, "Something happened in our end"
		}
		tickets = append(tickets, tkt)
	}

	//defer db.Close()
	return tickets, http.StatusOK, "Valid"
}
func TicketsFindByProject(projectId int, companyId int) ([]properties.Tickets, int, string) {
	db := database.Access()
	var tickets []properties.Tickets
	sqlQuery := `select TicketId,Name,Description,Status,AssignedTo,Priority,
	ProjectId,CreatedById from Tickets where ProjectId = ? and CompanyId = ?`
	fmt.Println(projectId, companyId)
	rows, err := db.Query(sqlQuery, projectId, companyId)
	if err != nil {
		fmt.Println(err.Error())
		return nil, http.StatusBadRequest, "Invalid request"
	}

	for rows.Next() {
		var tkt properties.Tickets
		if err := rows.Scan(&tkt.Id, &tkt.Name, &tkt.Description, &tkt.Status,
			&tkt.AssignedTo, &tkt.Priority, &tkt.ProjectId, &tkt.CreatedBy); err != nil {
			fmt.Println(err.Error())
			return nil, http.StatusInternalServerError, "Something happened in our end"
		}
		tickets = append(tickets, tkt)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, http.StatusInternalServerError, "Something happened in our end"
	}

	return tickets, http.StatusOK, "Valid"
}
