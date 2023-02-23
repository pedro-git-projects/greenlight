package data

import "github.com/pedro-git-projects/greenlight/internal/validator"

type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafeList []string
}

func NewFilters(page, pageSize int, sort string, sortSafeList []string) *Filters {
	return &Filters{
		Page:         page,
		PageSize:     pageSize,
		Sort:         sort,
		SortSafeList: make([]string, 0),
	}
}

func NewEmptyFilters() *Filters {
	return &Filters{
		Page:         *new(int),
		PageSize:     *new(int),
		Sort:         *new(string),
		SortSafeList: make([]string, 0),
	}
}

func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Page > 0, "page", "must be greater than zero")
	v.Check(f.Page <= 10_000_000, "page", "must be at maximum 10 million")
	v.Check(f.PageSize > 0, "page", "must be greater than zero")
	v.Check(f.PageSize <= 100, "page", "must at maximum 100")

	v.Check(validator.In(f.Sort, f.SortSafeList...), "sort", "invalid sort value")
}
