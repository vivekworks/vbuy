package main

import (
    "context"
    "flag"
    "fmt"
    "github.com/spf13/viper"
    "github.com/vivekworks/vbuy"
    "github.com/vivekworks/vbuy/db"
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
    DB         db.DB
    Logger     *zap.Logger
}

var (
    InvalidDBUsernameOrPassword = "invalid DB username or password"
)

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
        log.Fatal(InvalidDBUsernameOrPassword)
    }
    log.Println("After parsing flags")
    m := &Main{}
    m.LoadConfig(*Env)
    ctx = m.SetLogger(ctx, *Env)
    log.Println("After setting logger")
    if err := m.OpenDBConnection(ctx, *DBUser, *DBPassword); err != nil {
        m.Logger.Error("error opening db connection", zap.Error(err))
        os.Exit(1)
    }
    log.Println("After opening db connection pool")
    if err := m.StartServer(ctx); err != nil {
        m.Logger.Error("error starting server", zap.Error(err))
        os.Exit(1)
    }
    log.Println("After starting server")
    <-ctx.Done()
    log.Println("After Done")
    if err := m.Close(); err != nil {
        log.Fatalf("error closing resources: %v", err)
    }
    log.Println("After Close()")
}

func (m *Main) LoadConfig(env string) {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")

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
