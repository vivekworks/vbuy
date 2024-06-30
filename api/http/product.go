package http

import (
    "github.com/vivekworks/vbuy"
    "github.com/vivekworks/vbuy/service"
    "go.uber.org/zap"
    "net/http"
)

type ProductHandler struct {
    ps *service.ProductService
}

func NewProductHandler(ps *service.ProductService) *ProductHandler {
    return &ProductHandler{
        ps: ps,
    }
}

func (ph *ProductHandler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
    rInfo := vbuy.RequestInfoFromContext(r.Context())
    rInfo.Logger.Info("Inside HandlerGetUser", zap.String("requestID", rInfo.ID))
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("All OK!"))
}
