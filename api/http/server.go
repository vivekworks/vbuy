package http

import (
    "github.com/go-chi/chi/v5"
    "github.com/vivekworks/vbuy"
    "net/http"
)

type Server struct {
}

func NewServer() *http.Server {
    r := chi.NewRouter()
    r.Use(RequestLoggerMiddleware)
    userRouter := chi.NewRouter()
    userRouter.Route("/", func(rr chi.Router) {
        rr.Get("/{userID}", func(w http.ResponseWriter, r *http.Request) {
            rInfo := vbuy.RequestInfoFromContext(r.Context())
            rInfo.Logger.Info("Inside User ID route")
        })
    })
    productRouter := chi.NewRouter()
    productRouter.Route("/", func(rr chi.Router) {
        rr.Get("/{productID}", func(w http.ResponseWriter, r *http.Request) {
            rInfo := vbuy.RequestInfoFromContext(r.Context())
            rInfo.Logger.Info("Inside Product ID route")
        })
    })
    r.Mount("/api/products", productRouter)
    r.Mount("/api/users", userRouter)
    return &http.Server{
        Addr:    ":8080",
        Handler: r,
    }
}
