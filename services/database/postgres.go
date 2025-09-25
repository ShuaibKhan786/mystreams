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
	connStr := "postgres://postgres:dbMyStreams@localhost:5432/mystreamsdb?sslmode=disable"
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
		log.Errorf("Failed to prepare an insert query: %s : %v", query, err)
		return err
	}

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		log.Errorf("Failed to run an insert query: %s : %v", query, err)
		return err
	}

	return nil
}
