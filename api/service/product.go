package service

import (
    "context"
    "github.com/vivekworks/vbuy/db"
    "github.com/vivekworks/vbuy/http"
    "time"
)

type Product struct {
    ID           string
    Name         string
    ReleasedDate time.Time
    Model        string
    Price        float64
    Manufacturer string
    IsActive     bool
    CreatedAt    time.Time
    UpdatedAt    time.Time
    CreatedBy    string
    UpdatedBy    string
}

type ProductService struct {
    pr *db.ProductRepository
}

func NewProductService(pr *db.ProductRepository) *ProductService {
    return &ProductService{
        pr: pr,
    }
}

func (ps *ProductService) CreateProduct(ctx context.Context, p http.ProductCreate) (*Product, error) {
    return nil, nil
}
func (ps *ProductService) GetProduct(ctx context.Context, id string) (*Product, error) {
    return nil, nil
}
func (ps *ProductService) ListAllProducts(ctx context.Context) ([]*Product, error) {
    return nil, nil
}
func (ps *ProductService) UpdateProduct(ctx context.Context, id string, p http.ProductUpdate) (*Product, error) {
    return nil, nil
}
func (ps *ProductService) DeleteProduct(ctx context.Context, id string) (*Product, error) {
    return nil, nil
}
