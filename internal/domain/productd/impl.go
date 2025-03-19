package productd

import (
	"context"
	"fmt"
	"product-master/internal/utils"
	"strings"

	"github.com/jmoiron/sqlx"
)

type ProductDomain struct {
	pg *sqlx.DB
	utils.SQLTransaction
}

func NewProductDomain(pg *sqlx.DB, transaction utils.SQLTransaction) ProductRepo {
	return &ProductDomain{pg: pg, SQLTransaction: transaction}
}

func (d *ProductDomain) GetListCategory(ctx context.Context) (res []CategoryEntity, err error) {
	query := `select id, name, created_at, updated_at from categories where deleted_at is null order by name asc`
	err = d.pg.SelectContext(ctx, &res, query)
	return
}

func (d *ProductDomain) CreateProduct(ctx context.Context, tx *sqlx.Tx, params ProductEntity) (err error) {
	_, err = tx.NamedExecContext(ctx, `INSERT INTO products (id, name, category_id, price, created_at, updated_at)
	 VALUES(:id, :name, :category_id, :price, :created_at, :updated_at)`, params)
	return
}

func (d *ProductDomain) GetListProduct(ctx context.Context, filter FilterProduct) (res []ListProduct, err error) {
	var inputArgs []any
	var buffer strings.Builder
	buffer.WriteString(`select p.id, p.name, c.name as category_name, 
	p.price, p.created_at, p.updated_at
	from products p join categories c on p.category_id = c.id
	where p.deleted_at is null and c.deleted_at is null
	`)

	if filter.ID != "" {
		buffer.WriteString(" and p.id = ? ")
		inputArgs = append(inputArgs, filter.ID)
	}
	if filter.Name != "" {
		buffer.WriteString(" and LOWER(p.name) like ? ")
		inputArgs = append(inputArgs, "%"+strings.ToLower(filter.Name)+"%")
	}

	if filter.CategoryID != "" {
		buffer.WriteString(" and p.category_id = ? ")
		inputArgs = append(inputArgs, filter.CategoryID)
	}

	buffer.WriteString(" ORDER BY ")
	if filter.OrderBy == "" {
		buffer.WriteString(" p.name asc ")
	} else {
		if filter.SortType == "" {
			filter.SortType = "asc"
		}
		buffer.WriteString(fmt.Sprintf(" %s %s ", filter.OrderBy, filter.SortType))
	}

	filter.Page = (filter.Page - 1) * filter.Limit
	buffer.WriteString(" offset ? limit ?")
	inputArgs = append(inputArgs, filter.Page, filter.Limit)

	query := buffer.String()
	query, args, err := sqlx.In(query, inputArgs...)
	if err != nil {
		return
	}

	query = d.pg.Rebind(query)

	err = d.pg.SelectContext(ctx, &res, query, args...)
	return
}

func (d *ProductDomain) CountListProduct(ctx context.Context, filter FilterProduct) (res int64, err error) {
	var inputArgs []any
	var buffer strings.Builder
	buffer.WriteString(`select count(1)
	from products p join categories c on p.category_id = c.id
	where p.deleted_at is null and c.deleted_at is null
	`)

	if filter.ID != "" {
		buffer.WriteString(" and p.id = ? ")
		inputArgs = append(inputArgs, filter.ID)
	}
	if filter.Name != "" {
		buffer.WriteString(" and LOWER(p.name) like ? ")
		inputArgs = append(inputArgs, "%"+strings.ToLower(filter.Name)+"%")
	}
	if filter.CategoryID != "" {
		buffer.WriteString(" and p.category_id = ? ")
		inputArgs = append(inputArgs, filter.CategoryID)
	}

	query := buffer.String()
	query, args, err := sqlx.In(query, inputArgs...)
	if err != nil {
		return
	}

	query = d.pg.Rebind(query)
	err = d.pg.GetContext(ctx, &res, query, args...)
	return
}
