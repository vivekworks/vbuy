package postgres

import (
    "context"
    "github.com/jackc/pgx/v5"
    "github.com/vivekworks/vbuy"
    "log"
    "time"
)

type ProductRepository struct {
    db *pgx.Conn
}

func NewProductRepository(conn *pgx.Conn) *ProductRepository {
    return &ProductRepository{db: conn}
}

func (pr *ProductRepository) CreateProduct(ctx context.Context, p vbuy.ProductCreate) (*vbuy.Product, error) {
    tx, err := pr.db.Begin(ctx)
    if err != nil {
        log.Fatal(err)
        return nil, vbuy.ErrInternalServer
    }
    defer tx.Rollback(ctx)
    user := ctx.Value("user").(string)
    sql := "INSERT INTO PRODUCTS(name, released_date, model, manufacturer, price, is_active, created_by, updated_by) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, created_at, updated_at, is_active"
    var product vbuy.Product
    err = tx.QueryRow(ctx, sql, p.Name, time.Time(p.ReleasedDate), p.Model, p.Manufacturer, p.Price, p.IsActive, user, user).Scan(&product.ID, &product.CreatedAt, &product.UpdatedAt, &product.IsActive)
    if err != nil {
        log.Fatal(err)
        return nil, vbuy.ErrInternalServer
    }
    tx.Commit(ctx)
    product.CreatedBy = user
    product.UpdatedBy = user
    p.ToProduct(&product)
    return &product, nil
}

func (pr *ProductRepository) GetProduct(ctx context.Context, id string) (*vbuy.Product, error) {
    return nil, nil
}
func (pr *ProductRepository) ListAllProducts(ctx context.Context) ([]*vbuy.Product, error) {
    return nil, nil
}
func (pr *ProductRepository) UpdateProduct(ctx context.Context, id string, p vbuy.ProductUpdate) (*vbuy.Product, error) {
    return nil, nil
}
func (pr *ProductRepository) DeleteProduct(ctx context.Context, id string) (*vbuy.Product, error) {
    return nil, nil
}
