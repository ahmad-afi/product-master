package productu

import (
	"product-master/internal/helper"
	"time"
)

type ListCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ListProduct struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	CateogryName string    `json:"categoryName"`
	Price        float64   `json:"price"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type MetaProduct struct {
	TotalData int64         `json:"totalData"`
	Data      []ListProduct `json:"data"`
}

type FilterProduct struct {
	ID         string `query:"id"`
	Name       string `query:"name"`
	CategoryID string `query:"categoryID"`
	OrderBy    string `query:"orderBy"`
	SortType   string `query:"sortType"`
	helper.PaginationStruct
}

type CreateProduct struct {
	Name       string  `json:"name" validate:"required"`
	CategoryID string  `json:"categoryID" validate:"required"`
	Price      float64 `json:"price" validate:"required"`
}
