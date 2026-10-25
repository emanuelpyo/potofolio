package chat

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"deryl-portfolio/internal/config"
)

func post(h *Handler, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/api/chat", strings.NewReader(body))
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	return rr
}

func TestEmptyTextRejected(t *testing.T) {
	h := New(config.Config{AIAPIKey: "key", AIBaseURL: "http://unused", AIModel: "m"})
	rr := post(h, `{"text":"   "}`)
	if rr.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", rr.Code)
	}
	var payload map[string]any
	if err := json.Unmarshal(rr.Body.Bytes(), &payload); err != nil {
		t.Fatalf("body is not json: %v", err)
	}
	if payload["message"] != "Pesan tidak boleh kosong" {
		t.Fatalf("message = %v", payload["message"])
	}
}

func TestMissingKeyReturnsClearError(t *testing.T) {
	h := New(config.Config{AIBaseURL: "http://unused", AIModel: "m"})
	rr := post(h, `{"text":"hai"}`)
	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, want 500", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "Konfigurasi API AI tidak ditemukan.") {
		t.Fatalf("unexpected body: %s", rr.Body.String())
	}
	if strings.Contains(rr.Body.String(), "key") && strings.Contains(rr.Body.String(), "AI_API_KEY=") {
		t.Fatalf("response leaks configuration: %s", rr.Body.String())
	}
}

func TestMethodNotAllowed(t *testing.T) {
	h := New(config.Config{AIAPIKey: "key", AIBaseURL: "http://unused", AIModel: "m"})
	req := httptest.NewRequest(http.MethodGet, "/api/chat", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status = %d, want 405", rr.Code)
	}
}

func TestStreamsDeltaContent(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "Bearer test-key" {
			t.Errorf("authorization = %q", got)
		}
		w.Header().Set("Content-Type", "text/event-stream")
		for _, part := range []string{"Halo", " ", "Dunia"} {
			fmt.Fprintf(w, "data: {\"choices\":[{\"delta\":{\"content\":%q}}]}\n\n", part)
		}
		fmt.Fprint(w, "data: [DONE]\n\n")
	}))
	defer upstream.Close()

	h := New(config.Config{AIAPIKey: "test-key", AIBaseURL: upstream.URL, AIModel: "m"})
	rr := post(h, `{"text":"hai"}`)
	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rr.Code)
	}
	if got := rr.Body.String(); got != "Halo Dunia" {
		t.Fatalf("stream = %q, want %q", got, "Halo Dunia")
	}
}
