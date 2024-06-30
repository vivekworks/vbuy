package main

import (
    "context"
    "flag"
    "fmt"
    "github.com/vivekworks/vbuy"
    "github.com/vivekworks/vbuy/db"
    "github.com/vivekworks/vbuy/db/postgres"
    "github.com/vivekworks/vbuy/http"
    "go.uber.org/zap"
    "log"
    "os/signal"
    "syscall"
)

type Main struct {
    Config     *vbuy.Config
    HTTPServer *http.Server
    DB         db.DB
    Logger     *zap.Logger
}

var ErrInvalidDBUsernameOrPassword = "invalid DB username or password"

var (
    Env        = flag.String("ENV", vbuy.Development, "running environment")
    DBUser     = flag.String("DB_USER", "", "database username")
    DBPassword = flag.String("DB_PASSWORD", "", "database password")
)

func main() {
    ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
    defer stop()

    flag.Parse()
    if *DBUser == "" || *DBPassword == "" {
        log.Fatal(ErrInvalidDBUsernameOrPassword)
    }

    m := &Main{
        Config: vbuy.NewConfig(*Env),
    }

    ctx = m.SetLogger(ctx, *Env)

    if err := m.OpenDBConnection(ctx, *DBUser, *DBPassword); err != nil {
        log.Fatal(err)
    }

    errChan := make(chan error, 1)
    if err := m.StartServer(ctx, errChan); err != nil {
        log.Fatal(err)
    }
    select {
    case <-ctx.Done():
        log.Println("context cancelled, shutting down server")
    case err := <-errChan:
        log.Println(err)
    }

    if err := m.Close(); err != nil {
        log.Fatalf("server shutdown failed: %v", err)
    }
    log.Println("server shutdown gracefully")
}

func (m *Main) SetLogger(ctx context.Context, env string) context.Context {
    logger := vbuy.NewLogger(m.Config.Logger.Level, env)
    m.Logger = logger
    return vbuy.WithLogger(ctx, logger)
}

func (m *Main) OpenDBConnection(ctx context.Context, user string, password string) error {
    dbURL := fmt.Sprintf(postgres.UrlFormat, user, password, m.Config.DB.Host, m.Config.DB.Port, m.Config.DB.Name)
    pdb, err := postgres.NewDB(ctx, dbURL)
    if err != nil {
        return err
    }
    m.DB = pdb
    return nil
}

func (m *Main) StartServer(ctx context.Context, errChan chan<- error) error {
    pr := postgres.NewProductRepository(m.DB)
    server := http.NewServer(ctx, m.Config, pr)
    m.HTTPServer = server
    return server.Start(ctx, errChan)
}

func (m *Main) Close() error {
    if err := m.HTTPServer.Close(m.Config.HTTP.Timeout.Shutdown); err != nil {
        return err
    }
    m.DB.Close()
    return nil
}
