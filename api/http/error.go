package http

import (
    "context"
    "github.com/vivekworks/vbuy"
    "go.uber.org/zap"
    "net/http"
)

func HandleInternalServerError(ctx context.Context, w http.ResponseWriter) {
    rInfo := vbuy.RequestInfoFromContext(ctx)
    w.WriteHeader(vbuy.ErrInternalServer.Status)
    err := vbuy.WriteJSON(w, vbuy.ErrInternalServer.ToErrorResponse())
    if err != nil {
        rInfo.Logger.Error("error handling internal server error", zap.Error(err))
    }
}

func HandleInvalidPayload(ctx context.Context, w http.ResponseWriter, d []*vbuy.ErrorDetail) {
    rInfo := vbuy.RequestInfoFromContext(ctx)
    w.WriteHeader(vbuy.ErrInvalidPayload.Status)
    res := vbuy.ErrInvalidPayload
    if d != nil || len(d) > 0 {
        res.Detail = d
    }
    err := vbuy.WriteJSON(w, res.ToErrorResponse())
    if err != nil {
        rInfo.Logger.Error("error handling internal server error", zap.Error(err))
    }
}
