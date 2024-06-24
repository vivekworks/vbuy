package main

import (
//    "context"
//    "github.com/jackc/pgx/v5"
//    "github.com/vivekworks/vbuy"
//    "github.com/vivekworks/vbuy/db/postgres"
    "github.com/vivekworks/vbuy/http"
//    "go.uber.org/zap"
//    "log"
//    "time"
)

func main() {
//    ctx := context.Background()
    server := http.NewServer()
    if err := server.ListenAndServe(); err != nil {
        panic(err)
    }

//    logger, err := zap.NewProduction()
//    if err != nil {
//        log.Fatal("zap error", err)
//    }
//    connStr := "postgresql://vivekts:vivekts@localhost:5432/vbuy"
//    conn, err := pgx.Connect(ctx, connStr)
//    if err != nil {
//        logger.Info("postgres connection error", zap.Error(err))
//        panic(err)
//    }
//    ctx = context.WithValue(ctx, "user", "system")
//    pr := postgres.NewProductRepository(conn)
//    p := vbuy.ProductCreate{
//        Name:         "Vivek T S",
//        ReleasedDate: vbuy.Date(time.Now()),
//        Model:        "Human",
//        Price:        23.1237,
//        Manufacturer: "God",
//        IsActive:     false,
//    }
//    product, err := pr.CreateProduct(ctx, p)
//    if err != nil {
//        logger.Info("create product error", zap.Error(err))
//    }
//    logger.Info("Created Product", zap.Any("product", product))
}
