package http

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"

	"nerasync-dashboard/services/api/internal/domain"
)

// ReadinessPinger mengecek keterhubungan ke penyimpanan data publik.
type ReadinessPinger interface {
	Configured() bool
	Ping(ctx context.Context) error
}

// NewRouter membangun Chi router dengan middleware CORS, RequestID, rate limiting,
// serta endpoint walking skeleton (/healthz, /readyz, /api/v1/public/ping).
func NewRouter(appEnv string, corsAllowedOrigins []string, pinger ReadinessPinger) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(injectRequestIDHeader)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   corsAllowedOrigins,
		AllowedMethods:   []string{http.MethodGet, http.MethodHead, http.MethodOptions},
		AllowedHeaders:   []string{"Accept", "Content-Type", "If-None-Match", "X-Request-Id"},
		ExposedHeaders:   []string{"ETag", "X-Request-Id"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Use(httprate.LimitByIP(120, time.Minute))

	r.NotFound(func(w http.ResponseWriter, req *http.Request) {
		reqID := middleware.GetReqID(req.Context())
		WriteProblemJSON(
			w,
			req,
			http.StatusNotFound,
			"RESOURCE_NOT_FOUND",
			"Resource tidak ditemukan atau belum tersedia untuk publik.",
			reqID,
		)
	})

	r.Get("/healthz", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		WriteJSON(w, http.StatusOK, domain.HealthReport{
			Status:    "ok",
			Service:   "sinergi-api",
			Phase:     "F0",
			Timestamp: time.Now().UTC(),
		})
	})

	r.Get("/readyz", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		reqID := middleware.GetReqID(req.Context())

		dbConfigured := pinger != nil && pinger.Configured()
		dbReachable := false
		if dbConfigured {
			if err := pinger.Ping(req.Context()); err == nil {
				dbReachable = true
			}
		}

		report, ok := domain.EvaluateReadiness(appEnv, dbConfigured, dbReachable)
		if !ok {
			WriteProblemJSON(
				w,
				req,
				http.StatusServiceUnavailable,
				"SERVICE_UNAVAILABLE",
				"Layanan belum siap melayani permintaan publik.",
				reqID,
			)
			return
		}

		WriteJSON(w, http.StatusOK, report)
	})

	r.Route("/api/v1/public", func(pub chi.Router) {
		pub.Get("/ping", func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Cache-Control", "public, max-age=60, s-maxage=600, stale-while-revalidate=86400")
			WriteJSON(w, http.StatusOK, domain.PublicPingResponse{
				Message: "Hello SINERGI",
				Service: "sinergi-api",
				Phase:   "F0",
			})
		})
	})

	return r
}

func injectRequestIDHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if reqID := middleware.GetReqID(r.Context()); reqID != "" {
			w.Header().Set("X-Request-Id", reqID)
		}
		next.ServeHTTP(w, r)
	})
}
