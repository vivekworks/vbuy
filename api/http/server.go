package http

import (
    "context"
    "errors"
    "fmt"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/vivekworks/vbuy"
    "github.com/vivekworks/vbuy/db"
    "github.com/vivekworks/vbuy/service"
    "log"
    "net"
    "net/http"
    "time"
)

type Server struct {
    server *http.Server
}

func NewServer(
    ctx context.Context,
    c *vbuy.Config,
    pr db.ProductRepository,
) *Server {
    mux := chi.NewRouter()
    mux.Use(RequestLoggerMiddleware(ctx))
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
    s.registerProductRoutes(apiRouter, pr)
    return s
}

func (s *Server) Close(shutdownTimeout int) error {
    ctx, cancel := context.WithTimeout(context.Background(), time.Duration(shutdownTimeout)*time.Millisecond)
    defer cancel()
    if err := s.server.Shutdown(ctx); err != nil {
        return fmt.Errorf("error shutting down HTTP Server: %v", err)
    }
    return nil
}

func (s *Server) Start(ctx context.Context, errChan chan<- error) error {
    ln, err := net.Listen("tcp", s.server.Addr)
    if err != nil {
        return fmt.Errorf("error listening to tcp network at port %s: %v", s.server.Addr, err)
    }
    log.Println("Server listening on port", s.server.Addr)
    go func() {
        err = s.server.Serve(ln)
        if err != nil && !errors.Is(err, http.ErrServerClosed) {
            errChan <- fmt.Errorf("error serving incoming connections: %v", err)
        }
    }()
    return nil
}

func (s *Server) registerProductRoutes(mux *chi.Mux, pr db.ProductRepository) {
    ps := service.NewProductService(&pr)
    ph := NewProductHandler(ps)
    mux.Route("/products", func(r chi.Router) {
        r.Get("/", ph.HandleGetUser)
        r.Post("/", ph.HandlePostUser)
    })
}
