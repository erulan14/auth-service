package pgx

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewClient(host, port, username, password, database string) (*sql.DB, error) {
	config := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", username, password, host, port, database)

	sqlDB, err := sql.Open("pgx", config)
	if err != nil {
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	return sqlDB, nil
}
