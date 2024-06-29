package db

import (
    "context"
    "github.com/vivekworks/vbuy/service"
)

type ProductRepository interface {
    SaveProduct(ctx context.Context, p *service.Product) (*service.Product, error)
    GetProductByID(ctx context.Context, id string) (*service.Product, error)
    GetAllProducts(ctx context.Context) ([]*service.Product, error)
    UpdateProduct(ctx context.Context, p *service.Product) (*service.Product, error)
    DeleteProduct(ctx context.Context, id string) (*service.Product, error)
}
