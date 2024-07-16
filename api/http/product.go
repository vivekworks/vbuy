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
    ctx := r.Context()
    rInfo := vbuy.RequestInfoFromContext(ctx)
    var pc vbuy.ProductCreate
    if err := vbuy.ReadJSON(r.Body, &pc); err != nil {
        rInfo.Logger.Error("error decoding request body", zap.Error(err))
        HandleInvalidPayload(ctx, w, nil)
        return
    }
    if err := pc.Validate(); err != nil {
        HandleInvalidPayload(ctx, w, err)
        return
    }
    res, err := ph.ps.CreateProduct(ctx, pc)
    if err != nil {
        HandleInternalServerError(ctx, w)
        return
    }
    if err = vbuy.WriteJSON(w, res); err != nil {
        HandleInternalServerError(ctx, w)
        return
    }
}
