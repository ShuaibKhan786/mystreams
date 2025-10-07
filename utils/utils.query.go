package utils

import (
	"fmt"
	"strings"
)

type FilterQuery struct {
	Field  string
	Search string
}

// FORMAT: ?filter=foo;name
func ParseFilterQuery(query string) (*FilterQuery, error) {
	queryArr := strings.Split(query, ";")
	if len(queryArr) < 2 {
		return nil, fmt.Errorf("invalid filter query format")
	}

	return &FilterQuery{
		Field:  queryArr[1],
		Search: queryArr[0],
	}, nil
}

type SortQuery struct {
	Field string
	Order string
}

// FORMAT: ?sort=name,asc;created_at,desc
func ParseSortQuery(query string) ([]*SortQuery, error) {
	queryArr := strings.Split(query, ";")

	sortQueries := make([]*SortQuery, 0)

	for _, query := range queryArr {
		queryArr := strings.Split(query, ",")
		if len(queryArr) < 2 {
			return nil, fmt.Errorf("invalid filter query format")
		}

		order := ""

		switch queryArr[1] {
		case "asc":
			order = "ASC"
		case "desc":
			order = "DESC"
		}

		sortQueries = append(sortQueries, &SortQuery{
			Field: queryArr[0],
			Order: order,
		})
	}

	return sortQueries, nil
}
