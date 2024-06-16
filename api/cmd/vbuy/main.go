package main

import (
    "context"
    "github.com/jackc/pgx/v5"
    "github.com/vivekworks/vbuy"
    "github.com/vivekworks/vbuy/db/postgres"
    "log"
    "time"
)

func main() {
    ctx := context.Background()
    connStr := "postgresql://vivekts:vivekts@localhost:5432/vbuy"
    conn, err := pgx.Connect(ctx, connStr)
    if err != nil {
        log.Fatal(err)
    }
    ctx = context.WithValue(ctx, "user", "system")
    pr := postgres.NewProductRepository(conn)
    p := vbuy.ProductCreate{
        Name:         "Vivek T S",
        ReleasedDate: vbuy.Date(time.Now()),
        Model:        "Human",
        Price:        23.1237,
        Manufacturer: "God",
        IsActive:     false,
    }
    product, err := pr.CreateProduct(ctx, p)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Created Product: %v", product)
}
