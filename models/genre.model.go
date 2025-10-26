package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ShuaibKhan786/mystreams/services/database"
	"github.com/ShuaibKhan786/mystreams/utils"
)

type Genre struct {
	ID        *int    `form:"id,omitempty" json:"id,omitempty"`
	Name      *string `validation:"required" form:"name,omitempty" json:"name,omitempty"`
	UpdatedAt *time.Time
	CreatedAt *time.Time
}

type PaginationGenres struct {
	TotalCount *int
	Genres     []*Genre
}

func (g *Genre) Create(ctx context.Context) error {
	return database.RunQuery(
		ctx,
		"INSERT INTO genres (name) VALUES ($1)",
		[]interface{}{g.Name},
	)
}

func (g *Genre) Update(ctx context.Context) error {
	return database.RunQuery(
		ctx,
		"UPDATE genres SET name=$1, updated_at=NOW() WHERE id=$2",
		[]interface{}{g.Name, g.ID},
	)
}

func (g *Genre) Delete(ctx context.Context) error {
	return database.RunQuery(
		ctx,
		"DELETE FROM genres WHERE id=$1",
		[]interface{}{g.ID},
	)
}

func ReadGenreByID(ctx context.Context, id *int) *Genre {
	return database.RunSelectQuery(
		ctx,
		"SELECT id, name FROM genres WHERE id=$1",
		func(row *sql.Row) (*Genre, error) {
			var genre Genre

			err := row.Scan(&genre.ID, &genre.Name)
			if err != nil {
				return nil, err
			}

			return &genre, nil
		},
		id,
	)
}

func ReadGenres(
	ctx context.Context,
	paginationQuery *utils.PaginationQuery,
) *PaginationGenres {
	query := "SELECT id, name, updated_at, created_at FROM genres "
	queryOrder := " ORDER BY "
	queryPagination := " LIMIT %d OFFSET %d"
	queryCondition := " WHERE"
	queryOR := " OR"
	countQuery := "SELECT COUNT(id) FROM genres"

	if len(paginationQuery.Filter) > 0 {
		query += queryCondition
		countQuery += queryCondition
	}
	index := 1
	for field, serchQuery := range paginationQuery.Filter {
		query += serchQuery.DBQueryBuilder(field)
		countQuery += serchQuery.DBQueryBuilder(field)
		if index != len(paginationQuery.Filter) {
			query += queryOR
			countQuery += queryOR
		}
		index++
	}

	query += queryOrder
	if _, exists := paginationQuery.Sort["updated_at"]; !exists {
		query += "updated_at DESC"
		if len(paginationQuery.Sort) > 0 {
			query += ", "
		}
	}
	index = 1
	for field, order := range paginationQuery.Sort {
		// you can sanitize here
		switch order {
		case "asc":
			query += fmt.Sprintf(" %s %s", field, "ASC")
		case "desc":
			query += fmt.Sprintf(" %s %s", field, "DESC")
		}
		if index != len(paginationQuery.Sort) {
			query += ", "
		}
		index++
	}

	offset := (paginationQuery.Page - 1) * paginationQuery.Size
	query += fmt.Sprintf(queryPagination, paginationQuery.Size, offset)

	genres := database.RunSelectQueryList(
		ctx,
		query,
		func(rows *sql.Rows) (*Genre, error) {
			var genre Genre
			err := rows.Scan(&genre.ID, &genre.Name, &genre.UpdatedAt, &genre.CreatedAt)
			if err != nil {
				return nil, err
			}

			return &genre, nil
		},
	)

	count := database.RunSelectQuery(
		ctx,
		countQuery,
		func(row *sql.Row) (*int, error) {
			var count int
			err := row.Scan(&count)
			if err != nil {
				return nil, err
			}
			return &count, nil
		},
	)

	return &PaginationGenres{
		TotalCount: count,
		Genres:     genres,
	}
}
