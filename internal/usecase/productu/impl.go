package productu

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"product-master/internal/domain/productd"
	redisRepo "product-master/internal/domain/redis_repo"
	"product-master/internal/helper"
	"product-master/internal/utils"

	"github.com/jmoiron/sqlx"
)

type ProductUsecase struct {
	productRepo productd.ProductRepo
	redisRepo   redisRepo.RedisUsersRepository
}

func NewProductUsecase(productRepo productd.ProductRepo, redisRepo redisRepo.RedisUsersRepository) ProductUsc {
	return &ProductUsecase{productRepo: productRepo, redisRepo: redisRepo}
}

func (u *ProductUsecase) ListCategory(ctx context.Context) (res []ListCategory, err *helper.ErrorStruct) {
	resRepo, errRepo := u.productRepo.GetListCategory(ctx)
	if errRepo != nil {
		helper.Logger(helper.LoggerLevelError, "productu.ListCategory Error at GetListCategory", errRepo)
		err = helper.HelperErrorResponse(errRepo)
		return
	}

	for _, v := range resRepo {
		res = append(res, ListCategory{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	return
}

func (u *ProductUsecase) ListProduct(ctx context.Context, filter FilterProduct) (res MetaProduct, err *helper.ErrorStruct) {
	filter.DefaultPagination()
	res.Data = make([]ListProduct, 0)

	key := fmt.Sprintf("search:product:id:%s:name:%s:category_id:%s:orderby:%s:sorttype:%s:page:%d:limit:%d",
		filter.ID, filter.Name, filter.CategoryID, filter.OrderBy, filter.SortType, filter.Page, filter.Limit)

	resData, errRepo := u.redisRepo.GetData(ctx, key)
	if err != nil || resData == nil {
		err = nil
		helper.Logger(helper.LoggerLevelError, "productu.ListProduct Error at redisRepo.GetData", errRepo)
		if filter.OrderBy != "" {
			switch filter.OrderBy {
			case "date":
				filter.OrderBy = "created_at"
			case "price", "name":
			default:
				err = helper.HelperErrorResponse(fmt.Errorf("invalid orderby"), "invalid orderby")
				return
			}
		}

		if strings.EqualFold("asc", filter.SortType) {
			filter.SortType = "asc"
		} else if strings.EqualFold("desc", filter.SortType) {
			filter.SortType = "desc"
		} else {
			filter.SortType = "asc"
		}

		filterRepo := productd.FilterProduct{
			ID:               filter.ID,
			Name:             filter.Name,
			CategoryID:       filter.CategoryID,
			OrderBy:          filter.OrderBy,
			SortType:         filter.SortType,
			PaginationStruct: filter.PaginationStruct,
		}

		resRepo, errRepo := u.productRepo.GetListProduct(ctx, filterRepo)
		if errRepo != nil {
			helper.Logger(helper.LoggerLevelError, "productu.ListProduct Error at GetListCategory", errRepo)
			err = helper.HelperErrorResponse(errRepo)
			return
		}
		res.TotalData, errRepo = u.productRepo.CountListProduct(ctx, filterRepo)
		if errRepo != nil {
			helper.Logger(helper.LoggerLevelError, "productu.ListProduct Error at CountListProduct", errRepo)
			err = helper.HelperErrorResponse(errRepo)
			return
		}

		for _, v := range resRepo {
			res.Data = append(res.Data, ListProduct{
				ID:           v.ID,
				Name:         v.Name,
				CateogryName: v.CategoryName,
				Price:        v.Price,
				CreatedAt:    v.CreatedAt,
				UpdatedAt:    v.UpdatedAt,
			})
		}

		resByte, errMarshal := json.Marshal(res)
		if err != nil {
			helper.Logger(helper.LoggerLevelError, "productu.ListProduct Error at json.Marshal", errMarshal)
			err = helper.HelperErrorResponse(errMarshal)
			return
		}

		go func() {
			errRepo = u.redisRepo.SetData(context.Background(), key, string(resByte), time.Duration(time.Hour))
			if err != nil {
				helper.Logger(helper.LoggerLevelError, "productu.ListProduct Error at redisRepo.SetData", errRepo)
			}
		}()
	} else {
		errMarshal := json.Unmarshal([]byte(*resData), &res)
		if errMarshal != nil {
			helper.Logger(helper.LoggerLevelError, "productu.ListProduct Error at json.Unmarshal", errMarshal)
			err = helper.HelperErrorResponse(errMarshal)
			return
		}
	}

	return
}

func (u *ProductUsecase) CreateProduct(ctx context.Context, params CreateProduct) (productID string, err *helper.ErrorStruct) {
	if errValidate := utils.Validator(params); errValidate != nil {
		helper.Logger(helper.LoggerLevelError, "productu.CreateProduct Error at len(listProductID) <0", errValidate.Err)
		err = helper.HelperErrorResponse(errValidate.Err, errValidate.Message)
		return
	}

	productID, errIDGenerator := utils.IDGenerator()
	if errIDGenerator != nil {
		helper.Logger(helper.LoggerLevelError, "productu.CreateProduct Error at errIDGenerator", errIDGenerator)
		err = helper.HelperErrorResponse(errIDGenerator)
		return
	}

	if params.Price < 0 {
		err = helper.HelperErrorResponse(fmt.Errorf("price can't be negative"))
		return
	}

	dataRepo := productd.ProductEntity{
		ID:         productID,
		Name:       params.Name,
		CategoryID: params.CategoryID,
		Price:      params.Price,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	errRepo := u.productRepo.WrapperTransaction(ctx, func(tx *sqlx.Tx) (err error) {
		err = u.productRepo.CreateProduct(ctx, tx, dataRepo)
		if err != nil {
			helper.Logger(helper.LoggerLevelError, "productu.ListCategory Error at productRepo.CreateProduct", err)
			// err = helper.HelperErrorResponse(err)
			return
		}
		return
	})

	if errRepo != nil {
		helper.Logger(helper.LoggerLevelError, "productu.CreateProduct Error at WrapperTransaction", errRepo)
		err = helper.HelperErrorResponse(errRepo)
		return
	}

	return
}
