package models

import (
	"context"

	"github.com/ShuaibKhan786/mystreams/services/database"
)

type Gender string

const (
	MaleGender   Gender = "male"
	FemaleGender Gender = "female"
)

type Person struct {
	ID     *int64  `form:"id,omitempty" json:"id,omitempty"`
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
