package vbuy

import (
	"context"
	"time"
)

type Product struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	ReleasedDate time.Time `json:"releasedDate"`
}

type ProductService interface {
	CreateProduct(ctx context.Context, p Product) (*Product, error)
}
