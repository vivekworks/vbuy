package main

import (
    "context"
    "flag"
    "fmt"
    "github.com/spf13/viper"
    "github.com/vivekworks/vbuy"
    "github.com/vivekworks/vbuy/db/postgres"
    "github.com/vivekworks/vbuy/http"
    "go.uber.org/zap"
    "log"
    "os"
    "os/signal"
    "syscall"
)

type Main struct {
    Config     *vbuy.Config
    HTTPServer *http.Server
    DB         *postgres.DB
    Logger     *zap.Logger
}

var (
    Env        = flag.String("ENV", vbuy.Development, "running environment")
    DBUser     = flag.String("DB_USER", "", "database username")
    DBPassword = flag.String("DB_PASSWORD", "", "database password")
)

func main() {
    ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
    go func() {
        stop()
        if r := recover(); r != nil {
            log.Printf("application panic: %v", r)
        }
    }()

    flag.Parse()

    m := &Main{}
    m.LoadConfig(*Env)
    m.SetLogger(ctx, *Env)

    if err := m.OpenDBConnection(ctx, *DBUser, *DBPassword); err != nil {
        m.Logger.Error("error opening db connection", zap.Error(err))
        os.Exit(1)
    }
    if err := m.StartServer(ctx); err != nil {
        m.Logger.Error("error starting server", zap.Error(err))
        os.Exit(1)
    }
    <-ctx.Done()
    if err := m.Close(); err != nil {
        log.Fatalf("error closing resources: %v", err)
    }
}

func (m *Main) LoadConfig(env string) {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("../..")

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("error reading config: %v", err)
    }

    viper.SetConfigName(fmt.Sprintf("config.%s", env))

    if err := viper.MergeInConfig(); err != nil {
        log.Fatalf("error merging config: %v", err)
    }

    var c vbuy.Config
    err := viper.Unmarshal(&c)
    if err != nil {
        log.Fatalf("error unmarshalling config: %v", err)
    }
    m.Config = &c
}

func (m *Main) SetLogger(ctx context.Context, env string) {
    logger := vbuy.NewLogger(m.Config.Logger.Level, env)
    m.Logger = logger
    vbuy.WithLogger(ctx, logger)
}

func (m *Main) OpenDBConnection(ctx context.Context, user string, password string) error {
    dbURL := fmt.Sprintf("postgres://%s:%s&%s:%d/%s?sslmode=%s",
        user, password, m.Config.DB.Host, m.Config.DB.Port, m.Config.DB.Name, m.Config.DB.SslMode)
    db, err := postgres.NewDB(ctx, dbURL)
    if err != nil {
        return err
    }
    m.DB = db
    return nil
}

func (m *Main) StartServer(ctx context.Context) error {
    pr := postgres.NewProductRepository(m.DB)
    server := http.NewServer(ctx, m.Config, pr)
    m.HTTPServer = server
    return server.Start(ctx)
}

func (m *Main) Close() error {
    if err := m.HTTPServer.Close(m.Config.HTTP.Timeout.Shutdown); err != nil {
        return err
    }
    m.DB.Close()
    return nil
}
