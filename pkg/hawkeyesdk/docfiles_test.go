package hawkeyesdk

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDocFilesService_UploadFile_Success(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/savefile" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer token" {
			t.Fatalf("unexpected auth header: %s", got)
		}

		var payload map[string]any
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			t.Fatalf("failed to decode payload: %v", err)
		}
		if payload["category"] != DEFAULT.String() {
			t.Fatalf("expected default category, got %v", payload["category"])
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(ApiResponse{Message: "ok", Success: true})
	}))
	t.Cleanup(server.Close)

	client := &ClientSettings{
		AuthToken:  "token",
		BaseUrl:    server.URL,
		HTTPClient: server.Client(),
	}

	service := NewDocFilesService(client)

	resp, err := service.UploadFile(100, "https://file")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !resp.Success {
		t.Fatalf("expected success response")
	}
}

func TestDocFilesService_UploadFile_WithOptions(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload struct {
			Filenumber      int    `json:"filenumber"`
			Link            string `json:"link"`
			Category        string `json:"category"`
			VisibleToClient bool   `json:"visible_to_client"`
			Notes           string `json:"notes"`
		}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			t.Fatalf("failed to decode payload: %v", err)
		}
		if payload.Filenumber != 321 || payload.Link != "https://file" || payload.Category != IMAGES.String() || !payload.VisibleToClient || payload.Notes != "docs" {
			t.Fatalf("unexpected payload: %+v", payload)
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(ApiResponse{Message: "ok", Success: true})
	}))
	t.Cleanup(server.Close)

	client := &ClientSettings{
		AuthToken:  "token",
		BaseUrl:    server.URL,
		HTTPClient: server.Client(),
	}

	service := NewDocFilesService(client)

	if _, err := service.UploadFile(321, "https://file", WithCategory(IMAGES), WithVisibleToClient(true), WithNotes("docs")); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestDocFilesService_UploadFile_HTTPError(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		_ = json.NewEncoder(w).Encode(ApiResponse{Message: "nope"})
	}))
	t.Cleanup(server.Close)

	client := &ClientSettings{
		AuthToken:  "token",
		BaseUrl:    server.URL,
		HTTPClient: server.Client(),
	}

	service := NewDocFilesService(client)

	if _, err := service.UploadFile(123, "https://file"); err == nil {
		t.Fatalf("expected an error")
	}
}
