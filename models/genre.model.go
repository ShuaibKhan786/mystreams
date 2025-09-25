package models

import (
	"context"

	"github.com/ShuaibKhan786/mystreams/services/database"
)

type Genre struct {
	ID   *int64  `form:"id,omitempty" json:"id,omitempty"`
	Name *string `form:"name,omitempty" json:"name,omitempty"`
}

func (g *Genre) Create(ctx context.Context) error {
	return database.RunInsertQuery(
		ctx,
		"INSERT INTO genres (name) VALUES ($1)",
		[]interface{}{g.Name},
	)
}
