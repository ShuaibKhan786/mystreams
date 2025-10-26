package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ShuaibKhan786/mystreams/services/database"
	"github.com/ShuaibKhan786/mystreams/utils"
)

type Gender string

const (
	MaleGender   Gender = "male"
	FemaleGender Gender = "female"
)

type Person struct {
	ID        *int    `form:"id,omitempty" json:"id,omitempty"`
	Name      *string `validation:"required" form:"name,omitempty" json:"name,omitempty"`
	Gender    *Gender `validation:"required" form:"gender,omitempty" json:"gender,omitempty"`
	UpdatedAt *time.Time
	CreatedAt *time.Time
	// if you want you can add more information
}

type PaginationPeople struct {
	TotalCount *int
	People     []*Person
}

func (p *Person) Create(ctx context.Context) error {
	return database.RunQuery(
		ctx,
		"INSERT INTO people (name, gender) VALUES ($1, $2)",
		[]interface{}{p.Name, p.Gender},
	)
}

func (p *Person) Update(ctx context.Context) error {
	return database.RunQuery(
		ctx,
		"UPDATE people SET name=$1, gender=$2, updated_at=NOW() WHERE id=$3",
		[]interface{}{p.Name, p.Gender, p.ID},
	)
}

// warning: these is an hard delete
func (p *Person) Delete(ctx context.Context) error {
	return database.RunQuery(
		ctx,
		"DELETE FROM people WHERE id=$1",
		[]interface{}{p.ID},
	)
}

func ReadPersonByID(ctx context.Context, id *int) *Person {
	return database.RunSelectQuery(
		ctx,
		"SELECT id, name, gender FROM people WHERE id = $1",
		func(row *sql.Row) (*Person, error) {
			var person Person

			err := row.Scan(&person.ID, &person.Name, &person.Gender)
			if err != nil {
				return nil, err
			}

			return &person, nil
		},
		id,
	)
}

func ReadPeople(
	ctx context.Context,
	paginationQuery *utils.PaginationQuery,
) *PaginationPeople {
	query := "SELECT id, name, gender, updated_at, created_at FROM people "
	queryOrder := " ORDER BY "
	queryPagination := " LIMIT %d OFFSET %d"
	queryCondition := " WHERE"
	queryOR := " OR"
	countQuery := "SELECT COUNT(id) FROM people"

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

	people := database.RunSelectQueryList(
		ctx,
		query,
		func(rows *sql.Rows) (*Person, error) {
			var person Person
			err := rows.Scan(&person.ID, &person.Name, &person.Gender, &person.UpdatedAt, &person.CreatedAt)
			if err != nil {
				return nil, err
			}

			return &person, nil
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

	return &PaginationPeople{
		TotalCount: count,
		People:     people,
	}
}
