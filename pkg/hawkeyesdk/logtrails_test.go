package hawkeyesdk

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogTrailsService_CreateLogTrail_Success(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/createLogTailEntry" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		var payload struct {
			Filenumber int    `json:"filenumber"`
			Activity   string `json:"activity"`
			Date       string `json:"date"`
		}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			t.Fatalf("failed to decode payload: %v", err)
		}
		if payload.Filenumber != 7 || payload.Activity != "note" || payload.Date != "12/25/2024" {
			t.Fatalf("unexpected payload: %+v", payload)
		}
		_ = json.NewEncoder(w).Encode(ApiResponse{Message: "logged", Success: true})
	}))
	t.Cleanup(server.Close)

	client := &ClientSettings{
		AuthToken:  "token",
		BaseUrl:    server.URL,
		HTTPClient: server.Client(),
	}

	service := NewLogTrailsService(client)

	resp, err := service.CreateLogTrail(context.Background(), 7, "note", WithDate("12/25/2024"))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !resp.Success {
		t.Fatalf("expected success response")
	}
}

func TestLogTrailsService_CreateLogTrail_HTTPError(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ApiResponse{Message: "fail"})
	}))
	t.Cleanup(server.Close)

	client := &ClientSettings{
		AuthToken:  "token",
		BaseUrl:    server.URL,
		HTTPClient: server.Client(),
	}

	service := NewLogTrailsService(client)

	if _, err := service.CreateLogTrail(context.Background(), 1, "oops"); err == nil {
		t.Fatalf("expected error from server")
	}
}
