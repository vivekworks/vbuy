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

func NewWithRequestID(ctx context.Context, requestId string) context.Context {
    r := &RequestInfo{
        ID: requestId,
    }
    return context.WithValue(ctx, RequestInfoKey, r)
}

func RequestInfoFromContext(ctx context.Context) *RequestInfo {
    return ctx.Value(RequestInfoKey).(*RequestInfo)
}
