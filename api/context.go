package vbuy

import (
    "context"
    "go.uber.org/zap"
)

const RequestInfoKey = "requestInfo"

type RequestInfo struct {
    User   string
    Logger *zap.Logger
    ID     string
}

func NewRequestContext(ctx context.Context, requestId string, logger *zap.Logger) context.Context {
    r := &RequestInfo{
        ID:     requestId,
        Logger: logger,
    }
    return context.WithValue(ctx, RequestInfoKey, r)
}

func RequestInfoFromContext(ctx context.Context) *RequestInfo {
    return ctx.Value(RequestInfoKey).(*RequestInfo)
}
