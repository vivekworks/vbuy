package service

import (
	"context"
	"github.com/vivekworks/vbuy"
	"github.com/vivekworks/vbuy/db"
)

type ProductService struct {
	pr *db.ProductRepository
}

func NewProductService(pr *db.ProductRepository) *ProductService {
	return &ProductService{
		pr: pr,
	}
}

func (ps *ProductService) CreateProduct(ctx context.Context, p vbuy.ProductCreate) (*db.Product, error) {
	return nil, nil
}
func (ps *ProductService) GetProduct(ctx context.Context, id string) (*db.Product, error) {
	return nil, nil
}
func (ps *ProductService) ListAllProducts(ctx context.Context) ([]*db.Product, error) {
	return nil, nil
}
func (ps *ProductService) UpdateProduct(ctx context.Context, id string, p vbuy.ProductUpdate) (*db.Product, error) {
	return nil, nil
}
func (ps *ProductService) DeleteProduct(ctx context.Context, id string) (*db.Product, error) {
	return nil, nil
}
