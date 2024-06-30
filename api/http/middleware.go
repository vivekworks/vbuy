package http

import (
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

func RequestLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.NewString()
		RequestIdFieldType.String = requestID
		logger := vbuy.NewLoggerWithFields(vbuy.LoggerFromContext(r.Context()))
		ctx := vbuy.NewRequestContext(r.Context(), requestID, logger)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
