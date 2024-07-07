package http

import (
    "encoding/json"
    "github.com/go-chi/chi/v5"
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
    pId := chi.URLParam(r, "id")
    res, err := ph.ps.GetProduct(r.Context(), pId)
    if err != nil {
        _ = json.NewEncoder(w).Encode(err.(vbuy.Error).ToErrorResponse())
        return
    }
    _ = json.NewEncoder(w).Encode(res)
}

func (ph *ProductHandler) HandlePostUser(w http.ResponseWriter, r *http.Request) {
    rInfo := vbuy.RequestInfoFromContext(r.Context())
    var pc vbuy.ProductCreate
    if err := json.NewDecoder(r.Body).Decode(&pc); err != nil {
        rInfo.Logger.Error("error decoding request body", zap.Error(err))
        _ = json.NewEncoder(w).Encode(vbuy.ErrInternalServer.ToErrorResponse())
        return
    }
    res, err := ph.ps.CreateProduct(r.Context(), pc)
    if err != nil {
        _ = json.NewEncoder(w).Encode(err.(vbuy.Error).ToErrorResponse())
        return
    }
    _ = json.NewEncoder(w).Encode(res)
}
