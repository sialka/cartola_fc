package repository

import (
	"database/sql"

	"github.com/sialka/cartola_fc/internal/infra/db"
)

//var ErrQueriesNotSet = errors.New("queries not set")

// Informa que precisamos do conteudo de infra/db/queries.sql.go
type Repository struct {
	dbConn *sql.DB
	*db.Queries
}

/*
func (r *Repository) SetQuery(q *db.Queries) {
	r.Queries = q
}

func (r *Repository) Validade() error {
	if r.Queries == nil {
		return ErrQueriesNotSet
	}
	return nil
}
*/
