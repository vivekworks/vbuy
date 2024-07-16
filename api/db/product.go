package db

import (
    "context"
    "time"
)

type Product struct {
    ID           string    `db:"id"`
    Name         string    `db:"name"`
    ReleasedDate time.Time `db:"released_date"`
    Model        string    `db:"model"`
    Price        float64   `db:"price"`
    Manufacturer string    `db:"manufacturer"`
    Category     string    `db:"category"`
    IsActive     bool      `db:"is_active"`
    CreatedAt    time.Time `db:"created_at"`
    UpdatedAt    time.Time `db:"updated_at"`
    CreatedBy    string    `db:"created_by"`
    UpdatedBy    string    `db:"updated_by"`
}

type ProductRepository interface {
    SaveProduct(ctx context.Context, p *Product) (*Product, error)
    GetProductByID(ctx context.Context, id string) (*Product, error)
    GetAllProducts(ctx context.Context) ([]*Product, error)
    UpdateProduct(ctx context.Context, p *Product) (*Product, error)
    DeleteProduct(ctx context.Context, id string) (*Product, error)
}
