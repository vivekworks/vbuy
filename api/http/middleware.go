package http

import (
    "github.com/google/uuid"
    "github.com/vivekworks/vbuy"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "net/http"
)

func RequestLoggerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        requestID := uuid.NewString()
        encoderConfig := zap.NewProductionEncoderConfig()
        encoderConfig.TimeKey = "timestamp"
        encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
        config := zap.Config{
            Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
            Development:       false,
            DisableCaller:     false,
            DisableStacktrace: false,
            Sampling:          nil,
            Encoding:          "json",
            EncoderConfig:     encoderConfig,
            OutputPaths: []string{
                "stderr",
            },
            ErrorOutputPaths: []string{
                "stderr",
            },
            InitialFields: map[string]interface{}{
                "requestID": requestID,
            },
        }
        logger := zap.Must(config.Build())
        ctx := vbuy.NewRequestContext(r.Context(), requestID, logger)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
