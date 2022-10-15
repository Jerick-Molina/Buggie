package types

import "database/sql"

type Database struct {
	*sql.DB
}
