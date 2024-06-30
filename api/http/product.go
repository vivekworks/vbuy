package http

import (
	"github.com/vivekworks/vbuy/service"
    "log"
    "net/http"
)

type ProductHandler struct {
	ps *service.ProductService
}

func NewProductHandler(ps *service.ProductService) *ProductHandler {
	return &ProductHandler{
		ps: ps,
	}
}

func (ph *ProductHandler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside HandleGetUser")
}
