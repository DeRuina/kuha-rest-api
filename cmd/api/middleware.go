package main

import (
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/DeRuina/KUHA-REST-API/internal/auth/authn"
	"github.com/DeRuina/KUHA-REST-API/internal/ratelimiter"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

func (app *api) BasicAuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// read the auth header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				utils.UnauthorizedBasicErrorResponse(w, r, fmt.Errorf("authorization header is missing"))
				return
			}

			// parse it -> get the base64
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Basic" {
				utils.UnauthorizedBasicErrorResponse(w, r, fmt.Errorf("authorization header is malformed"))
				return
			}

			// decode it
			decoded, err := base64.StdEncoding.DecodeString(parts[1])
			if err != nil {
				utils.UnauthorizedBasicErrorResponse(w, r, err)
				return
			}

			// check the credentials
			username := app.config.auth.basic.user
			pass := app.config.auth.basic.pass

			creds := strings.SplitN(string(decoded), ":", 2)
			if len(creds) != 2 || creds[0] != username || creds[1] != pass {
				utils.UnauthorizedBasicErrorResponse(w, r, fmt.Errorf("invalid credentials"))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func JWTMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				utils.UnauthorizedResponse(w, r, fmt.Errorf("missing or malformed Authorization header"))
				return
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			_, claims, err := authn.ValidateJWT(tokenStr)
			if err != nil {
				utils.UnauthorizedResponse(w, r, err)
				return
			}

			clientName, _ := claims["sub"].(string)
			rawRoles, _ := claims["roles"].([]interface{})

			var roles []string
			for _, r := range rawRoles {
				if s, ok := r.(string); ok {
					roles = append(roles, s)
				}
			}

			ctx := authn.WithClientMetadata(r.Context(), clientName, roles)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func ExtractClientIDMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				next.ServeHTTP(w, r)
				return
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			_, claims, err := authn.ValidateJWT(tokenStr)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			clientName, _ := claims["sub"].(string)
			ctx := authn.WithClientMetadata(r.Context(), clientName, nil)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func (app *api) RateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !app.config.rateLimiter.Enabled {
			next.ServeHTTP(w, r)
			return
		}

		clientID := authn.GetClientName(r.Context())
		if clientID == "" {
			clientID = r.RemoteAddr
		}

		limit, window := ratelimiter.GetLimitForRole(clientID)

		if app.redisRateLimiter != nil {
			allowed, retryAfter, err := app.redisRateLimiter.Allow(r.Context(), clientID, limit, window)
			if err != nil {
				utils.InternalServerError(w, r, err)
				return
			}
			if !allowed {
				utils.RateLimitExceededResponse(w, r, retryAfter.String())
				return
			}
		} else if app.localRateLimiter != nil {
			allowed, retryAfter := app.localRateLimiter.Allow(clientID)
			if !allowed {
				utils.RateLimitExceededResponse(w, r, retryAfter.String())
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func GzipDecompressionMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !strings.EqualFold(r.Header.Get("Content-Encoding"), "gzip") {
				next.ServeHTTP(w, r)
				return
			}

			const maxCompressed = int64(200 * 1024 * 1024)    // 200 MB (compressed)
			const maxDecompressed = int64(1024 * 1024 * 1024) // 1 GB   (after decompression)

			r.Body = http.MaxBytesReader(w, r.Body, maxCompressed)

			gz, err := gzip.NewReader(r.Body)
			if err != nil {
				utils.BadRequestResponse(w, r, fmt.Errorf("invalid gzip payload: %w", err))
				return
			}
			defer gz.Close()

			rc := http.MaxBytesReader(w, gz, maxDecompressed)

			r.Body = rc

			r.Header.Del("Content-Encoding")
			r.Header.Set("X-Was-Gzipped", "true")

			next.ServeHTTP(w, r)
		})
	}
}
