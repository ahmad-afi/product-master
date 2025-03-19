package productu

import (
	"context"
	"product-master/internal/helper"
)

type ProductUsc interface {
	ListCategory(ctx context.Context) (res []ListCategory, err *helper.ErrorStruct)
	ListProduct(ctx context.Context, filter FilterProduct) (res MetaProduct, err *helper.ErrorStruct)
	CreateProduct(ctx context.Context, params CreateProduct) (productID string, err *helper.ErrorStruct)
}
