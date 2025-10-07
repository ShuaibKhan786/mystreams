package database

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2/log"
)

var db *sql.DB

func GetPSQLInstance() *sql.DB {
	return db
}

func ConnectPSQL() (*sql.DB, error) {
	connStr := "postgres://postgres:myStreams@localhost:5432/mystreamsdb?sslmode=disable"
	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Error("Failed to connect the postgres database: ", err)
		return nil, err
	}
	log.Info("Successfully connected to postgres database")

	db = DB

	return db, nil
}

func RunInsertQuery(ctx context.Context, query string, args []interface{}) error {
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Errorf("Failed to prepare an insert query: %s : %v\n", query, err)
		return err
	}

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		log.Errorf("Failed to run an insert query: %s : %v\n", query, err)
		return err
	}
	return nil
}

func RunSelectQuery[T any](
	ctx context.Context,
	query string,
	fn func(row *sql.Row) (*T, error),
	args ...any,
) *T {
	row := db.QueryRowContext(ctx, query, args...)
	result, err := fn(row)
	if err != nil {
		log.Errorf("Failed to scan a row query: %s : %v\n", query, err)
		return nil
	}

	return result
}

func RunSelectQueryList[T any](
	ctx context.Context,
	query string,
	fn func(rows *sql.Rows) (*T, error),
	args ...any,
) []*T {
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Errorf("Failed to run a select query: %s : %v\n", query, err)
		return nil
	}

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		log.Errorf("Failed to query a select query: %s : %v\n", query, err)
		return nil
	}
	defer rows.Close()

	var results []*T
	for rows.Next() {
		row, err := fn(rows)
		if err != nil {
			log.Errorf("Failed to scan the row query: %s : %v\n", query, err)
			return nil
		}

		results = append(results, row)
	}

	return results
}
