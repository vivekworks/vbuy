package vbuy

import (
    "context"
    "go.uber.org/zap"
)

const (
    RequestInfoKey = "requestInfo"
    LoggerKey      = "logger"
)

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

func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
    return context.WithValue(ctx, LoggerKey, logger)
}

func RequestInfoFromContext(ctx context.Context) *RequestInfo {
    return ctx.Value(RequestInfoKey).(*RequestInfo)
}

func LoggerFromContext(ctx context.Context) *zap.Logger {
    return ctx.Value(LoggerKey).(*zap.Logger)
}
