package http

import (
	"encoding/json"
	"net/http"
)

// ProblemDetails merepresentasikan format error standar RFC 9457 (application/problem+json).
type ProblemDetails struct {
	Type      string `json:"type"`
	Title     string `json:"title"`
	Status    int    `json:"status"`
	Detail    string `json:"detail"`
	Instance  string `json:"instance"`
	RequestID string `json:"request_id,omitempty"`
}

// WriteJSON menulis respons JSON dengan status kode HTTP yang ditentukan.
func WriteJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

// WriteProblemJSON menulis error RFC 9457 tanpa mengekspos detail internal server atau stack trace.
func WriteProblemJSON(w http.ResponseWriter, r *http.Request, status int, errorCode string, detail string, requestID string) {
	title := http.StatusText(status)
	if title == "" {
		title = "Error"
	}

	problem := ProblemDetails{
		Type:      "https://sinergi.internal/errors/" + errorCode,
		Title:     title,
		Status:    status,
		Detail:    detail,
		Instance:  r.URL.Path,
		RequestID: requestID,
	}

	w.Header().Set("Content-Type", "application/problem+json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(problem)
}
