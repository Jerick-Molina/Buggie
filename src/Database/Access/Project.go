package access

import (
	"fmt"
	"net/http"

	properties "github.com/Jerick-Molina/Buggie/Properties"
	database "github.com/Jerick-Molina/Buggie/src/Database"
)

func ProjectSearchByCompanyId(companyId int) ([]properties.Projects, int, string) {

	db := database.Access()

	var projects []properties.Projects
	sqlQuery := `select ProjectId,Name,Description from Projects where CompanyId = ?`

	rows, err := db.Query(sqlQuery, companyId)
	if err != nil {
		fmt.Println(err.Error())
		return nil, http.StatusBadRequest, "Invalid request"
	}

	for rows.Next() {
		var project properties.Projects
		if err := rows.Scan(&project.Id, &project.Name, &project.Description); err != nil {
			return nil, http.StatusInternalServerError, "Something happened in our end"
		}
		projects = append(projects, project)
	}
	if err := rows.Err(); err != nil {
		return nil, http.StatusInternalServerError, "Something happened in our end"
	}
	return projects, http.StatusOK, "Valid"

}

func ProjectSearchByProjectId(projectId int, companyId int) (properties.Projects, int, string) {
	db := database.Access()

	var projects properties.Projects
	sqlQuery := `select Name,Description from Projects where CompanyId = ? and ProjectId = ?`

	if err := db.QueryRow(sqlQuery, companyId, projectId).Scan(&projects.Name, &projects.Description); err != nil {
		return properties.Projects{}, http.StatusBadRequest, "Invalid Project Id"
	}

	return projects, http.StatusOK, "Valid"
}

func EditProject() {
	sql := ""

	fmt.Println(sql)
}

func DeleteProject() {
	sql := ""

	fmt.Println(sql)
}
