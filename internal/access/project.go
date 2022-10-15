package access

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/Jerick-Molina/Buggie/api/types"
// )

// func InsertProject() {

// }
// func ModifyProject() {

// }

// func QueryProjectByCompanyId(compId int) ([]types.Projects, int) {

// 	var projects []types.Projects
// 	sqlQuery := `select ProjectId,Name,Description from Projects where CompanyId = ?`

// 	rows, err := db.Query(sqlQuery, compId)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil, http.StatusBadRequest
// 	}

// 	for rows.Next() {
// 		var project types.Projects
// 		if err := rows.Scan(&project.Id, &project.Name, &project.Description); err != nil {
// 			return nil, http.StatusInternalServerError
// 		}
// 		projects = append(projects, project)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, http.StatusInternalServerError
// 	}
// 	return projects, 200
// }

// func QueryProjectByProjectId(compId int, projectId int) (types.Projects, int) {
// 	var projects types.Projects

// 	sqlQuery := `select Name,Description from Projects where CompanyId = ? and ProjectId = ?`

// 	if err := db.QueryRow(sqlQuery, compId, projectId).Scan(&projects.Name, &projects.Description); err != nil {
// 		return types.Projects{}, http.StatusBadRequest
// 	}
// 	return projects, 200
// }
