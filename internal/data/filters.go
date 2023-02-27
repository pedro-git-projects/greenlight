package data

import (
	"strings"

	"github.com/pedro-git-projects/greenlight/internal/validator"
)

// Filters represents the filtering options and parameters
type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafeList []string
}

// NewFilters returns a pointer to a new Filters instance
func NewFilters(page, pageSize int, sort string, sortSafeList []string) *Filters {
	return &Filters{
		Page:         page,
		PageSize:     pageSize,
		Sort:         sort,
		SortSafeList: make([]string, 0),
	}
}

// NewEmptyFilters retuns a safe zeroed pointer to a new Filters instance
func NewEmptyFilters() *Filters {
	return &Filters{
		Page:         *new(int),
		PageSize:     *new(int),
		Sort:         *new(string),
		SortSafeList: make([]string, 0),
	}
}

// ValidateFilters checks if the Filters fields are valid
func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Page > 0, "page", "must be greater than zero")
	v.Check(f.Page <= 10_000_000, "page", "must be at maximum 10 million")
	v.Check(f.PageSize > 0, "page", "must be greater than zero")
	v.Check(f.PageSize <= 100, "page", "must at maximum 100")

	v.Check(validator.In(f.Sort, f.SortSafeList...), "sort", "invalid sort value")
}

// Checks if the cliend provided sort field matches with a entrie
// in the safe list
func (f Filters) sortColumn() string {
	for _, safe := range f.SortSafeList {
		if f.Sort == safe {
			return strings.TrimPrefix(f.Sort, "-")
		}
	}
	panic("unsafe sort parameter: " + f.Sort)
}

// sortDirection returns "ASC" or "DESC" depending on the prefix
// rune of the sort field
func (f Filters) sortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}
	return "ASC"
}

// limit returns the pagination limit
func (f Filters) limit() int {
	return f.PageSize
}

// offset returns the pagination offset
func (f Filters) offset() int {
	return (f.Page - 1) * f.PageSize
}
