package main

import (
	"product-master/internal/delivery/http"
	"product-master/internal/infrastructure/container"
)

func main() {
	cont := container.NewContainer()
	http.HTTPRouteInit(cont)
}
