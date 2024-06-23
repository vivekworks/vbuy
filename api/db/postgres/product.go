package postgres

import (
    "context"
    "github.com/jackc/pgx/v5"
    "github.com/vivekworks/vbuy"
    "go.uber.org/zap"
    "time"
)

type ProductRepository struct {
    db *pgx.Conn
}

func NewProductRepository(conn *pgx.Conn) *ProductRepository {
    return &ProductRepository{db: conn}
}

func (pr *ProductRepository) CreateProduct(ctx context.Context, p vbuy.ProductCreate) (*vbuy.Product, error) {
    rInfo := vbuy.RequestInfoFromContext(ctx)
    tx, err := pr.db.Begin(ctx)
    if err != nil {
        rInfo.Logger.Error("error beginning transaction", zap.Error(err))
        return nil, vbuy.ErrInternalServer
    }
    defer tx.Rollback(ctx)
    query := `INSERT INTO PRODUCTS(
                name,
                released_date,
                model,
                manufacturer,
                price,
                is_active,
                created_by,
                updated_by)
            VALUES($1, $2, $3, $4, $5, $6, $7, $8)
            RETURNING id, created_at, updated_at, is_active`
    var product vbuy.Product
    err = tx.QueryRow(ctx, query, p.Name, time.Time(p.ReleasedDate), p.Model, p.Manufacturer, p.Price, p.IsActive, rInfo.User, rInfo.User).Scan(&product.ID, &product.CreatedAt, &product.UpdatedAt, &product.IsActive)
    if err != nil {
        rInfo.Logger.Error("error querying row", zap.Error(err))
        return nil, vbuy.ErrInternalServer
    }
    tx.Commit(ctx)
    product.CreatedBy, product.UpdatedBy = rInfo.User, rInfo.User
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
