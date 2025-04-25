package logger

import (
	"net/http"
	"time"

	"github.com/DeRuina/KUHA-REST-API/internal/auth/authn"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

// Initialize the global zap logger
func Init() {
	rawLogger, _ := zap.NewProduction()
	Logger = rawLogger.Sugar()
}

// Ensures logs are flushed before the application exits
func Cleanup() {
	_ = Logger.Sync()
}

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseRecorder) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rr := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(rr, r)

		clientID := authn.GetClientName(r.Context())
		if clientID == "" {
			clientID = "anonymous"
		}

		requestID := middleware.GetReqID(r.Context())

		logFields := []zap.Field{
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.String("query_params", r.URL.RawQuery),
			zap.String("client_id", clientID),
			zap.String("request_id", requestID),
			zap.String("ip", r.RemoteAddr),
			zap.String("user_agent", r.UserAgent()),
			zap.Int("status", rr.statusCode),
			zap.Duration("response_time", time.Since(start)),
		}

		switch {
		case rr.statusCode >= 500:
			Logger.Desugar().With(logFields...).Error("Internal server error")
		case rr.statusCode >= 400:
			Logger.Desugar().With(logFields...).Warn("error response")
		default:
			Logger.Desugar().With(logFields...).Info("Request completed")
		}
	})
}
