package handler

import "product-master/internal/infrastructure/container"

type handler struct {
	ProductHandler productHandler
}

func SetupHandler(cont container.Container) handler {
	return handler{
		ProductHandler: NewProductHandler(cont.ProductUsc),
	}
}
