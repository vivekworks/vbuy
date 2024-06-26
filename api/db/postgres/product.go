package postgres

import (
    "context"
    "github.com/vivekworks/vbuy"
    "github.com/vivekworks/vbuy/db"
    "go.uber.org/zap"
)

type ProductRepository struct {
    db *DB
}

func NewProductRepository(pdb db.DB) *ProductRepository {
    return &ProductRepository{
        db: pdb.(*DB),
    }
}

func (pr *ProductRepository) SaveProduct(ctx context.Context, p *db.Product) (*db.Product, error) {
    rInfo := vbuy.RequestInfoFromContext(ctx)
    tx, err := pr.db.pool.Begin(ctx)
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
    err = tx.QueryRow(ctx, query, p.Name, p.ReleasedDate, p.Model, p.Manufacturer, p.Price, p.IsActive, rInfo.User, rInfo.User).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt, &p.IsActive)
    if err != nil {
        rInfo.Logger.Error("error querying row", zap.Error(err))
        return nil, vbuy.ErrInternalServer
    }
    tx.Commit(ctx)
    p.CreatedBy, p.UpdatedBy = rInfo.User, rInfo.User
    return nil, nil
}

func (pr *ProductRepository) GetProductByID(ctx context.Context, id string) (*db.Product, error) {
    return nil, nil
}
func (pr *ProductRepository) GetAllProducts(ctx context.Context) ([]*db.Product, error) {
    return nil, nil
}
func (pr *ProductRepository) UpdateProduct(ctx context.Context, p *db.Product) (*db.Product, error) {
    return nil, nil
}
func (pr *ProductRepository) DeleteProduct(ctx context.Context, id string) (*db.Product, error) {
    return nil, nil
}
