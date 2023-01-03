// package repo

// import (
// 	"database/sql"

// 	db "github.com/erick291992/transaction-backend/src/db/sqlc"
// )

// type Repo struct {
// 	*db.Queries
// 	DB *sql.DB
// }

package db

import (
	"database/sql"
	// dbp "github.com/erick291992/transaction-backend/src/db/sqlc"
)

type Repo struct {
	*Queries
	DB *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		DB:      db,
		Queries: New(db),
	}
}
