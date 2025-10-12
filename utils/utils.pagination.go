package utils

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	DEFAULT_PAGINATION_PAGE int = 1
	DEFAULT_PAGINATION_SIZE int = 10
)

type PaginationQuery struct {
	Page   int `query:"page"`
	Size   int `query:"size"`
	Filter map[string]string
	Sort   map[string]string
}

func NewPaginationQuery() PaginationQuery {
	return PaginationQuery{
		Page:   0,
		Size:   0,
		Filter: make(map[string]string),
		Sort:   make(map[string]string),
	}
}

func (p *PaginationQuery) Parse(c *fiber.Ctx) error {
	err := c.QueryParser(p)
	if err != nil {
		return err
	}

	queries := c.Queries()

	for key, value := range queries {
		switch {
		case strings.HasPrefix(key, "filter["):
			field := key[len("filter[") : len(key)-1]
			p.Filter[field] = value
		case strings.HasPrefix(key, "sort["):
			field := key[len("sort[") : len(key)-1]
			p.Sort[field] = value
		}
	}

	return nil
}

func (p *PaginationQuery) Encode() string {
	params := make([]string, 0)

	if p.Page > 0 {
		params = append(params, fmt.Sprintf("page=%d", p.Page))
	}
	if p.Size > 0 {
		params = append(params, fmt.Sprintf("size=%d", p.Size))
	}

	for field, value := range p.Filter {
		params = append(params, fmt.Sprintf("filter[%s]=%s", field, value))
	}

	for field, value := range p.Sort {
		params = append(params, fmt.Sprintf("sort[%s]=%s", field, value))
	}

	return strings.Join(params, "&")
}

func (p *PaginationQuery) FilterQuery() string {
	query := ""

	for field, value := range p.Filter {
		query += fmt.Sprintf("filter[%s]=%s", field, value)
	}

	return query
}

func (p *PaginationQuery) SortQuery() string {
	query := ""

	for field, value := range p.Sort {
		query += fmt.Sprintf("sort[%s]=%s", field, value)
	}

	return query
}

func CalculatePaginationPages(totalCount, size int) int {
	if size <= 0 {
		return 0
	}
	pages := totalCount / size
	if totalCount%size != 0 {
		pages++
	}
	return pages
}
