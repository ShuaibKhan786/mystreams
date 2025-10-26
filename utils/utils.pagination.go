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

type Filter struct {
	Operator string
	Value    string
}

func (f *Filter) DBQueryBuilder(field string) string {
	switch f.Operator {
	case "match":
		return fmt.Sprintf(" %s ILIKE '%%%s%%'", field, f.Value)
	case "similar":
		return fmt.Sprintf(" similarity(%s, '%s') > 0.3", field, f.Value)
	case "eql":
		return fmt.Sprintf(" %s = '%s'", field, f.Value)
	default:
		return fmt.Sprintf(" %s = '%s'", field, f.Value)
	}
}

type PaginationQuery struct {
	Page   int `query:"page"`
	Size   int `query:"size"`
	Filter map[string]Filter
	Sort   map[string]string
}

func NewPaginationQuery() PaginationQuery {
	return PaginationQuery{
		Page:   0,
		Size:   0,
		Filter: make(map[string]Filter),
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
			fields := key[len("filter[") : len(key)-1]
			if strings.Contains(fields, ":") && value != "" {
				fieldArr := strings.Split(fields, ":")
				if len(fieldArr) >= 2 {
					p.Filter[fieldArr[0]] = Filter{
						Operator: fieldArr[1],
						Value:    value,
					}
				}
			}
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
		params = append(params, fmt.Sprintf("filter[%s:%s]=%s", field, value.Operator, value.Value))
	}

	for field, value := range p.Sort {
		params = append(params, fmt.Sprintf("sort[%s]=%s", field, value))
	}

	return strings.Join(params, "&")
}

func (p *PaginationQuery) FilterQuery() string {
	params := make([]string, 0)

	for field, value := range p.Filter {
		params = append(params, fmt.Sprintf("filter[%s:%s]=%s", field, value.Operator, value.Value))
	}

	return strings.Join(params, "&")
}

func (p *PaginationQuery) SortQuery() string {
	params := make([]string, 0)

	for field, value := range p.Sort {
		params = append(params, fmt.Sprintf("sort[%s]=%s", field, value))
	}

	return strings.Join(params, "&")
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
