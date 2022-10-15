package access

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/Jerick-Molina/Buggie/api/types"
// )

// func InsertTicket(compId int, creatorId int, tkt *types.Tickets) int {
// 	sqlExec := `insert into Tickets
// 	(Name,AssignedTo,CreatedById,CompanyId,Description,ProjectId,Priority)
// 	values(?,?,?,?,?,?,?)`

// 	result, err := db.Exec(sqlExec, tkt.Name, tkt.AssignedTo,
// 		creatorId, compId, tkt.Description, tkt.ProjectId, tkt.Priority)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return http.StatusInternalServerError
// 	}
// 	fmt.Println(result)

// 	return 200
// }

// func ModifyTicket() {

// }

// func QueryTicketByCompanyId(compId int) ([]types.Tickets, int) {

// 	var tickets []types.Tickets

// 	sqlQuery := `select TicketId,Name,Description,Status,AssignedTo,Priority,
// 	ProjectId from Tickets where CompanyId = ?`

// 	rows, err := db.Query(sqlQuery, compId)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil, http.StatusBadRequest
// 	}

// 	for rows.Next() {
// 		var tkt types.Tickets
// 		if err := rows.Scan(&tkt.Id, &tkt.Name, &tkt.Description, &tkt.Status,
// 			&tkt.AssignedTo, &tkt.Priority, &tkt.ProjectId); err != nil {
// 			fmt.Println(compId)
// 			fmt.Println(err.Error())
// 			return nil, http.StatusInternalServerError
// 		}
// 		tickets = append(tickets, tkt)
// 	}

// 	//defer db.Close()
// 	return tickets, http.StatusOK

// }

// func QueryTicketByProjectId(compId int, projectId int) ([]types.Tickets, int) {
// 	var tickets []types.Tickets
// 	sqlQuery := `select TicketId,Name,Description,Status,AssignedTo,Priority,
// 	ProjectId,CreatedById from Tickets where ProjectId = ? and CompanyId = ?`
// 	fmt.Println(projectId, compId)
// 	rows, err := db.Query(sqlQuery, projectId, compId)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil, http.StatusBadRequest
// 	}

// 	for rows.Next() {
// 		var tkt types.Tickets
// 		if err := rows.Scan(&tkt.Id, &tkt.Name, &tkt.Description, &tkt.Status,
// 			&tkt.AssignedTo, &tkt.Priority, &tkt.ProjectId, &tkt.CreatedBy); err != nil {
// 			fmt.Println(err.Error())
// 			return nil, http.StatusInternalServerError

// 		}
// 		tickets = append(tickets, tkt)
// 	}
// 	if err := rows.Err(); err != nil {
// 		fmt.Println(err.Error())
// 		return nil, http.StatusInternalServerError
// 	}
// 	return tickets, 200
// }
