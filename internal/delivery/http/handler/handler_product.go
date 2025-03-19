package handler

import (
	"net/http"
	"product-master/internal/helper"
	"product-master/internal/usecase/productu"

	"github.com/gofiber/fiber/v2"
)

type productHandler struct {
	productUsc productu.ProductUsc
}

func NewProductHandler(productUsc productu.ProductUsc) productHandler {
	return productHandler{productUsc: productUsc}
}

func (h *productHandler) ListCategory(ctx *fiber.Ctx) error {
	res, err := h.productUsc.ListCategory(ctx.Context())
	if err != nil {
		return helper.BuildResponse(ctx, false, err.Message, nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, res, http.StatusOK)
}

func (h *productHandler) ListProduct(ctx *fiber.Ctx) error {
	data := new(productu.FilterProduct)
	if err := ctx.QueryParser(data); err != nil {
		return helper.BuildResponse(ctx, false, err.Error(), nil, http.StatusBadRequest)
	}

	res, err := h.productUsc.ListProduct(ctx.Context(), *data)
	if err != nil {
		return helper.BuildResponse(ctx, false, err.Message, nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, res, http.StatusOK)
}

func (h *productHandler) CreateProduct(ctx *fiber.Ctx) error {
	data := new(productu.CreateProduct)
	if err := ctx.BodyParser(data); err != nil {
		return helper.BuildResponse(ctx, false, err.Error(), nil, http.StatusBadRequest)
	}

	productID, err := h.productUsc.CreateProduct(ctx.Context(), *data)
	if err != nil {
		return helper.BuildResponse(ctx, false, err.Message, nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDPOSTDATA, productID, http.StatusOK)
}
