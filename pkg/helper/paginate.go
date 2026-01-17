package helper

import (
	"gorm.io/gorm"
)

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func Paginate(pagination Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		defaultPage := 1
		defaultPageSize := 10

		page := pagination.Page
		if page <= 0 {
			page = defaultPage
		}

		pageSize := pagination.PageSize
		if pageSize <= 0 {
			pageSize = defaultPageSize
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = defaultPageSize
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
