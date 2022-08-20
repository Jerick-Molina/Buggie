package access

import "github.com/Jerick-Molina/Buggie/api/types"

func InsertTicket(compId int, projectId int, tkt *types.Tickets) int {

	return 200
}

func ModifyTicket() {

}

func QueryTicketByCompanyId(compId int) ([]types.User, int) {

	return nil, 200
}

func QueryTicketByProjectId(compId int, projectId int) ([]types.User, int) {

	return nil, 200
}
