package postgres

import (
    "context"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/vivekworks/vbuy"
    "github.com/vivekworks/vbuy/db"
    "github.com/vivekworks/vbuy/service"
    "go.uber.org/zap"
    "time"
)

type ProductRepository struct {
    db *pgxpool.Pool
}

func NewProductRepository(pool *pgxpool.Pool) *ProductRepository {
    return &ProductRepository{
        db: pool,
    }
}

func (pr *ProductRepository) SaveProduct(ctx context.Context, p *service.Product) (*db.Product, error) {
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
    var product service.Product
    err = tx.QueryRow(ctx, query, p.Name, time.Time(p.ReleasedDate), p.Model, p.Manufacturer, p.Price, p.IsActive, rInfo.User, rInfo.User).Scan(&product.ID, &product.CreatedAt, &product.UpdatedAt, &product.IsActive)
    if err != nil {
        rInfo.Logger.Error("error querying row", zap.Error(err))
        return nil, vbuy.ErrInternalServer
    }
    tx.Commit(ctx)
    product.CreatedBy, product.UpdatedBy = rInfo.User, rInfo.User
    return nil, nil
}

func (pr *ProductRepository) GetProductByID(ctx context.Context, id string) (*db.Product, error) {
    return nil, nil
}
func (pr *ProductRepository) GetAllProducts(ctx context.Context) ([]*db.Product, error) {
    return nil, nil
}
func (pr *ProductRepository) UpdateProduct(ctx context.Context, p *service.Product) (*db.Product, error) {
    return nil, nil
}
func (pr *ProductRepository) DeleteProduct(ctx context.Context, id string) (*db.Product, error) {
    return nil, nil
}
