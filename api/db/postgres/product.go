package postgres

import (
    "context"
    "github.com/vivekworks/vbuy"
    "github.com/vivekworks/vbuy/service"
    "go.uber.org/zap"
    "time"
)

type ProductRepository struct {
    db *DB
}

func NewProductRepository(db *DB) *ProductRepository {
    return &ProductRepository{
        db: db,
    }
}

func (pd *ProductRepository) SaveProduct(ctx context.Context, p *service.Product) (*service.Product, error) {
    rInfo := vbuy.RequestInfoFromContext(ctx)
    tx, err := pd.db.pool.Begin(ctx)
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

func (pd *ProductRepository) GetProductByID(ctx context.Context, id string) (*service.Product, error) {
    return nil, nil
}
func (pd *ProductRepository) GetAllProducts(ctx context.Context) ([]*service.Product, error) {
    return nil, nil
}
func (pd *ProductRepository) UpdateProduct(ctx context.Context, p *service.Product) (*service.Product, error) {
    return nil, nil
}
func (pd *ProductRepository) DeleteProduct(ctx context.Context, id string) (*service.Product, error) {
    return nil, nil
}
