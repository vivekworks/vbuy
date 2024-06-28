package http

import (
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "net/http"
)

type Server struct {
    server *http.Server
    router *chi.Mux
}

func NewServer() *Server {
    s := &Server{
        server: &http.Server{
            Addr: ":8080",
        },
        router: chi.NewRouter(),
    }
    s.router.Use(RequestLoggerMiddleware)
    s.router.Use(middleware.StripSlashes)
    apiRouter := chi.NewRouter()
    s.router.Mount("/api", apiRouter)
    s.registerProductRoutes(apiRouter)
    //    &http.Server{
    //        Addr:    ":8080",
    //        Handler: r,
    //        IdleTimeout: 30 * time.Second,
    //        ReadTimeout: 200 * time.Millisecond,
    //        ReadHeaderTimeout: 50 * time.Millisecond,
    //        WriteTimeout: 1 * time.Second,
    //    }
    return s
}

func (s *Server) registerProductRoutes(mr *chi.Mux) {
    ph := NewProductHandler(nil)
    mr.Route("/products", func(r chi.Router) {
        r.Get("/", ph.HandleGetUser)
    })
}
