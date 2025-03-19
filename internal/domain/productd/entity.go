package productd

import (
	"product-master/internal/helper"
	"time"
)

type CategoryEntity struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type ProductEntity struct {
	ID         string    `db:"id"`
	Name       string    `db:"name"`
	CategoryID string    `db:"category_id"`
	Price      float64   `db:"price"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type FilterProduct struct {
	ID         string `db:"id"`
	Name       string `db:"name"`
	CategoryID string `db:"categoryID"`
	OrderBy    string `db:"order_by"`
	SortType   string `db:"sort_type"`
	helper.PaginationStruct
}

type ListProduct struct {
	ID           string    `db:"id"`
	Name         string    `db:"name"`
	CategoryName string    `db:"category_name"`
	Price        float64   `db:"price"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
