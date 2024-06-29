package http

import (
    "context"
    "fmt"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/vivekworks/vbuy"
    "github.com/vivekworks/vbuy/db"
    "github.com/vivekworks/vbuy/service"
    "go.uber.org/zap"
    "net"
    "net/http"
    "time"
)

type Server struct {
    server *http.Server
    logger *zap.Logger
}

func NewServer(
    ctx context.Context,
    c *vbuy.Config,
    pr db.ProductRepository,
) *Server {
    logger := vbuy.LoggerFromContext(ctx)
    mux := chi.NewRouter()
    mux.Use(RequestLoggerMiddleware)
    mux.Use(middleware.Recoverer)
    apiRouter := chi.NewRouter()
    mux.Mount("/api", apiRouter)
    s := &Server{
        server: &http.Server{
            Addr:              fmt.Sprintf(":%s", c.HTTP.Port),
            Handler:           mux,
            IdleTimeout:       time.Duration(c.HTTP.Timeout.Idle) * time.Millisecond,
            ReadHeaderTimeout: time.Duration(c.HTTP.Timeout.ReadHeader) * time.Millisecond,
            ReadTimeout:       time.Duration(c.HTTP.Timeout.Read) * time.Millisecond,
            WriteTimeout:      time.Duration(c.HTTP.Timeout.Write) * time.Millisecond,
        },
    }
    s.logger = logger
    s.registerProductRoutes(apiRouter, pr)
    return s
}

func (s *Server) Close(shutdownTimeout int) error {
    ctx, cancel := context.WithTimeout(context.Background(), time.Duration(shutdownTimeout)*time.Millisecond)
    defer cancel()
    err := s.server.Shutdown(ctx)
    if err != nil {
        return err
    }
    return nil
}

func (s *Server) Start(ctx context.Context) error {
    logger := vbuy.LoggerFromContext(ctx)
    ln, err := net.Listen("tcp", s.server.Addr)
    if err != nil {
        logger.Error("error listening to tcp network", zap.String("addr", s.server.Addr), zap.Error(err))
        return err
    }
    go func() {
        err = s.server.Serve(ln)
        if err != nil {
            logger.Error("error serving incoming connections", zap.Error(err))
        }
    }()
    return nil
}

func (s *Server) registerProductRoutes(mux *chi.Mux, pr db.ProductRepository) {
    ps := service.NewProductService(&pr)
    ph := NewProductHandler(ps)
    mux.Route("/products", func(r chi.Router) {
        r.Get("/", ph.HandleGetUser)
    })
}
