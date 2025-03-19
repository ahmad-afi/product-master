package productd

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	WrapperTransaction(ctx context.Context, fn func(tx *sqlx.Tx) error) (err error)
	GetListCategory(ctx context.Context) (res []CategoryEntity, err error)
	CreateProduct(ctx context.Context, tx *sqlx.Tx, params ProductEntity) (err error)
	GetListProduct(ctx context.Context, filter FilterProduct) (res []ListProduct, err error)
	CountListProduct(ctx context.Context, filter FilterProduct) (res int64, err error)
}
