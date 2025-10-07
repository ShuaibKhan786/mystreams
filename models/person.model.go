package models

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ShuaibKhan786/mystreams/services/database"
	"github.com/ShuaibKhan786/mystreams/utils"
)

type Gender string

const (
	MaleGender   Gender = "male"
	FemaleGender Gender = "female"
)

type Person struct {
	ID     *int    `form:"id,omitempty" json:"id,omitempty"`
	Name   *string `form:"name,omitempty" json:"name,omitempty"`
	Gender *Gender `form:"gender,omitempty" json:"gender,omitempty"`
	// if you want you can add more information
}

func (p *Person) Create(ctx context.Context) error {
	return database.RunInsertQuery(
		ctx,
		"INSERT INTO people (name, gender) VALUES ($1, $2);",
		[]interface{}{*p.Name, *p.Gender},
	)
}

func ReadPersonByID(ctx context.Context, id *int) *Person {
	return database.RunSelectQuery(
		ctx,
		"SELECT name, gender FROM people WHERE id = $1",
		func(row *sql.Row) (*Person, error) {
			var person Person

			err := row.Scan(&person.Name, &person.Gender)
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
	page, size int,
	filter *utils.FilterQuery,
	sort []*utils.SortQuery,
) []*Person {
	query := "SELECT id, name, gender FROM people "
	queryOrder := " ORDER BY %s %s"
	queryPagination := " LIMIT %d OFFSET %d"
	queryCondition := " WHERE %s=%s"

	queryPagination = fmt.Sprintf(queryPagination, size, (page-1)*size)

	fmt.Println("SOrt ", sort)
	for i, s := range sort {
		switch s.Field {
		case "name":
			query += fmt.Sprintf(queryOrder, "name", s.Order)
		case "gender":
			query += fmt.Sprintf(queryOrder, "gender", s.Order)
		}
		if i < len(sort)-1 {
			query += ", "
		}
	}

	query += queryPagination

	if filter != nil {
		query += fmt.Sprintf(queryCondition, filter.Field, filter.Search)
	}

	fmt.Println(query)

	people := database.RunSelectQueryList(
		ctx,
		query,
		func(rows *sql.Rows) (*Person, error) {
			var person Person
			err := rows.Scan(&person.ID, &person.Name, &person.Gender)
			if err != nil {
				return nil, err
			}

			return &person, nil
		},
	)

	return people
}
