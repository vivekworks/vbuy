package http

import (
    "context"
    "github.com/google/uuid"
    "github.com/vivekworks/vbuy"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "net/http"
)

var RequestIdFieldType = zap.Field{
    Key:  "requestID",
    Type: zapcore.StringType,
}

func RequestLoggerMiddleware(ctx context.Context) func(handler http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            requestID := uuid.NewString()
            RequestIdFieldType.String = requestID
            logger := vbuy.NewLoggerWithFields(vbuy.LoggerFromContext(ctx))
            c := vbuy.NewRequestContext(r.Context(), requestID, logger)
            next.ServeHTTP(w, r.WithContext(c))
        })
    }
}

func ResponseContentTypeMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        next.ServeHTTP(w, r)
    })
}
