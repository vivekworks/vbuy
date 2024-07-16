package service

import (
    "context"
    "github.com/vivekworks/vbuy"
    "github.com/vivekworks/vbuy/db"
    "time"
)

type ProductService struct {
    pr db.ProductRepository
}

func NewProductService(pr db.ProductRepository) *ProductService {
    return &ProductService{
        pr: pr,
    }
}

func (ps *ProductService) CreateProduct(ctx context.Context, p vbuy.ProductCreate) (*vbuy.Product, error) {
    product, err := ps.pr.SaveProduct(ctx, &db.Product{
        Name:         p.Name,
        ReleasedDate: time.Time(*p.ReleasedDate),
        Model:        p.Model,
        Price:        float64(p.Price),
        Manufacturer: p.Manufacturer,
        IsActive:     p.IsActive,
        Category:     string(p.Category),
    })
    if err != nil {
        return nil, err
    }
    releasedDate := vbuy.Date(product.ReleasedDate)
    price := vbuy.Money(product.Price)
    return &vbuy.Product{
        ID:           product.ID,
        Name:         product.Name,
        ReleasedDate: &releasedDate,
        Model:        product.Model,
        Price:        &price,
        Category:     vbuy.Category(product.Category),
        Manufacturer: product.Manufacturer,
        IsActive:     product.IsActive,
        CreatedAt:    product.CreatedAt,
        UpdatedAt:    product.UpdatedAt,
        CreatedBy:    product.CreatedBy,
        UpdatedBy:    product.UpdatedBy,
    }, nil
}
func (ps *ProductService) GetProduct(ctx context.Context, id string) (*vbuy.Product, error) {
    product, err := ps.pr.GetProductByID(ctx, id)
    if err != nil {
        return nil, err
    }
    releasedDate := vbuy.Date(product.ReleasedDate)
    price := vbuy.Money(product.Price)
    return &vbuy.Product{
        ID:           product.ID,
        Name:         product.Name,
        ReleasedDate: &releasedDate,
        Model:        product.Model,
        Price:        &price,
        Category:     vbuy.Category(product.Category),
        Manufacturer: product.Manufacturer,
        IsActive:     product.IsActive,
        CreatedAt:    product.CreatedAt,
        UpdatedAt:    product.UpdatedAt,
        CreatedBy:    product.CreatedBy,
        UpdatedBy:    product.UpdatedBy,
    }, nil
}
func (ps *ProductService) ListAllProducts(ctx context.Context) ([]*vbuy.Product, error) {
    return nil, nil
}
func (ps *ProductService) UpdateProduct(ctx context.Context, id string, p vbuy.ProductUpdate) (*vbuy.Product, error) {
    return nil, nil
}
func (ps *ProductService) DeleteProduct(ctx context.Context, id string) (*vbuy.Product, error) {
    return nil, nil
}
