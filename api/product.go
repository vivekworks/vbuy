package vbuy

import (
    "context"
    "time"
)

type Product struct {
    ID           string    `json:"id"`
    Name         string    `json:"name"`
    ReleasedDate Date      `json:"releasedDate"`
    Model        string    `json:"model"`
    Price        Money     `json:"price"`
    Manufacturer string    `json:"manufacturer"`
    IsActive     bool      `json:"isActive"`
    CreatedAt    time.Time `json:"createdAt"`
    UpdatedAt    time.Time `json:"updatedAt"`
    CreatedBy    string    `json:"createdBy"`
    UpdatedBy    string    `json:"updatedBy"`
}

type ProductService interface {
    CreateProduct(ctx context.Context, p ProductCreate) (*Product, error)
    GetProduct(ctx context.Context, id string) (*Product, error)
    ListAllProducts(ctx context.Context) ([]*Product, error)
    UpdateProduct(ctx context.Context, id string, p ProductUpdate) (*Product, error)
    DeleteProduct(ctx context.Context, id string) (*Product, error)
}

type ProductCreate struct {
    Name         string `json:"name"`
    ReleasedDate Date   `json:"releasedDate"`
    Model        string `json:"model"`
    Price        Money  `json:"price"`
    Manufacturer string `json:"manufacturer"`
    IsActive     bool   `json:"isActive"`
}

func (pc *ProductCreate) ToProduct(p *Product) {
    p.Manufacturer = pc.Manufacturer
    p.Model = pc.Model
    p.Name = pc.Name
    p.Price = pc.Price
    p.ReleasedDate = pc.ReleasedDate
}

type ProductUpdate struct {
    Price    Money `json:"price"`
    IsActive bool  `json:"isActive"`
}
