package domain

import "time"

// HealthReport menyatakan status liveness proses backend Go.
type HealthReport struct {
	Status    string    `json:"status"`
	Service   string    `json:"service"`
	Phase     string    `json:"phase"`
	Timestamp time.Time `json:"timestamp"`
}

// ReadinessReport menyatakan status kesiapan melayani trafik (liveness + database).
type ReadinessReport struct {
	Status    string    `json:"status"`
	Database  string    `json:"database"`
	Service   string    `json:"service"`
	Timestamp time.Time `json:"timestamp"`
}

// PublicPingResponse adalah payload walking skeleton untuk menguji konektivitas publik.
type PublicPingResponse struct {
	Message string `json:"message"`
	Service string `json:"service"`
	Phase   string `json:"phase"`
}

// EvaluateReadiness menentukan kesiapan layanan secara fail-closed.
// Di luar mode development, database wajib terkonfigurasi dan dapat dijangkau.
func EvaluateReadiness(appEnv string, dbConfigured bool, dbReachable bool) (ReadinessReport, bool) {
	now := time.Now().UTC()

	if dbConfigured && dbReachable {
		return ReadinessReport{
			Status:    "ready",
			Database:  "up",
			Service:   "sinergi-api",
			Timestamp: now,
		}, true
	}

	if !dbConfigured && appEnv == "development" {
		return ReadinessReport{
			Status:    "ready",
			Database:  "not_configured",
			Service:   "sinergi-api",
			Timestamp: now,
		}, true
	}

	dbState := "down"
	if !dbConfigured {
		dbState = "not_configured"
	}

	return ReadinessReport{
		Status:    "not_ready",
		Database:  dbState,
		Service:   "sinergi-api",
		Timestamp: now,
	}, false
}
