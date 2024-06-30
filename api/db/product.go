package db

import (
	"context"
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

type ProductRepository interface {
	SaveProduct(ctx context.Context, p *Product) (*Product, error)
	GetProductByID(ctx context.Context, id string) (*Product, error)
	GetAllProducts(ctx context.Context) ([]*Product, error)
	UpdateProduct(ctx context.Context, p *Product) (*Product, error)
	DeleteProduct(ctx context.Context, id string) (*Product, error)
}
