package domain

import "testing"

func TestEvaluateReadiness(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		appEnv       string
		dbConfigured bool
		dbReachable  bool
		wantOK       bool
		wantStatus   string
		wantDatabase string
	}{
		{
			name:         "production db configured and reachable",
			appEnv:       "production",
			dbConfigured: true,
			dbReachable:  true,
			wantOK:       true,
			wantStatus:   "ready",
			wantDatabase: "up",
		},
		{
			name:         "development db not configured still ready for skeleton",
			appEnv:       "development",
			dbConfigured: false,
			dbReachable:  false,
			wantOK:       true,
			wantStatus:   "ready",
			wantDatabase: "not_configured",
		},
		{
			name:         "production db not configured fails closed",
			appEnv:       "production",
			dbConfigured: false,
			dbReachable:  false,
			wantOK:       false,
			wantStatus:   "not_ready",
			wantDatabase: "not_configured",
		},
		{
			name:         "db configured but unreachable fails closed",
			appEnv:       "development",
			dbConfigured: true,
			dbReachable:  false,
			wantOK:       false,
			wantStatus:   "not_ready",
			wantDatabase: "down",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, ok := EvaluateReadiness(tc.appEnv, tc.dbConfigured, tc.dbReachable)
			if ok != tc.wantOK {
				t.Fatalf("ok = %v, want %v", ok, tc.wantOK)
			}
			if got.Status != tc.wantStatus {
				t.Fatalf("status = %q, want %q", got.Status, tc.wantStatus)
			}
			if got.Database != tc.wantDatabase {
				t.Fatalf("database = %q, want %q", got.Database, tc.wantDatabase)
			}
		})
	}
}
