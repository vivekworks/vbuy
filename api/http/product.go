package http

import (
    "github.com/vivekworks/vbuy"
    "github.com/vivekworks/vbuy/service"
    "net/http"
    "time"
)

type Product struct {
    ID           string     `json:"id"`
    Name         string     `json:"name"`
    ReleasedDate vbuy.Date  `json:"releasedDate"`
    Model        string     `json:"model"`
    Price        vbuy.Money `json:"price"`
    Manufacturer string     `json:"manufacturer"`
    IsActive     bool       `json:"isActive"`
    CreatedAt    time.Time  `json:"createdAt"`
    UpdatedAt    time.Time  `json:"updatedAt"`
    CreatedBy    string     `json:"createdBy"`
    UpdatedBy    string     `json:"updatedBy"`
}

type ProductCreate struct {
    Name         string     `json:"name"`
    ReleasedDate vbuy.Date  `json:"releasedDate"`
    Model        string     `json:"model"`
    Price        vbuy.Money `json:"price"`
    Manufacturer string     `json:"manufacturer"`
    IsActive     bool       `json:"isActive"`
}

type ProductUpdate struct {
    Price    vbuy.Money `json:"price"`
    IsActive bool       `json:"isActive"`
}

type ProductHandler struct {
    ps *service.ProductService
}

func NewProductHandler(ps *service.ProductService) *ProductHandler {
    return &ProductHandler{
        ps: ps,
    }
}

func (ph *ProductHandler) HandleGetUser(w http.ResponseWriter, r *http.Request) {

}
