package pgx

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewDb(host, port, username, password, database string) (*sqlx.DB, error) {
	config := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", username, password, host, port, database)
	db, err := sqlx.Connect("postgres", config)
	if err != nil {
		return nil, err
	}
	return db, nil
}
