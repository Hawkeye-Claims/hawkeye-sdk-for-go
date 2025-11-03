package hawkeyesdk

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestClaimsService_CreateClaim_Success(t *testing.T) {
	t.Parallel()

	expectedResponse := ApiResponse{Filenumber: 123, Message: "ok", Success: true}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/createclaim" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer test-token" {
			t.Fatalf("unexpected auth header: %s", got)
		}
		var body ClaimPost
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("failed to decode body: %v", err)
		}
		if body.RenterName != "Test Renter" {
			t.Fatalf("unexpected renter name: %s", body.RenterName)
		}
		if body.InsCompaniesID != "Hawkeye" {
			t.Fatalf("unexpected insurance company: %s", body.InsCompaniesID)
		}
		if body.VehVIN != "VIN123" {
			t.Fatalf("unexpected VIN: %s", body.VehVIN)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(expectedResponse)
	}))
	t.Cleanup(server.Close)

	client := &ClientSettings{
		AuthToken:  "test-token",
		BaseUrl:    server.URL,
		HTTPClient: server.Client(),
	}

	service := NewClaimsService(client)

	resp, err := service.CreateClaim(context.Background(), ClaimPost{
		RenterName:     "Test Renter",
		InsCompaniesID: "Hawkeye",
		DateOfLoss:     "2024-01-01",
		VehMake:        "Ford",
		VehModel:       "F150",
		VehColor:       "Blue",
		VehVIN:         "VIN123",
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp != expectedResponse {
		t.Fatalf("unexpected response: %+v", resp)
	}
}

func TestClaimsService_CreateClaim_HTTPError(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ApiResponse{Message: "bad request"})
	}))
	t.Cleanup(server.Close)

	client := &ClientSettings{
		AuthToken:  "test-token",
		BaseUrl:    server.URL,
		HTTPClient: server.Client(),
	}

	service := NewClaimsService(client)

	_, err := service.CreateClaim(context.Background(), ClaimPost{
		RenterName:     "Test Renter",
		InsCompaniesID: "Hawkeye",
		DateOfLoss:     "2024-01-01",
		VehMake:        "Ford",
		VehModel:       "F150",
		VehColor:       "Blue",
		VehVIN:         "VIN123",
	})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestClaimsService_CreateClaim_MissingRequiredFields(t *testing.T) {
	t.Parallel()

	client := &ClientSettings{
		AuthToken:  "token",
		BaseUrl:    "http://example.com",
		HTTPClient: http.DefaultClient,
	}

	service := NewClaimsService(client)

	_, err := service.CreateClaim(context.Background(), ClaimPost{
		RenterName:     "Test Renter",
		InsCompaniesID: "Hawkeye",
		DateOfLoss:     "2024-01-01",
		VehMake:        "Ford",
		VehModel:       "F150",
		VehColor:       "Blue",
		VehVIN:         "",
	})
	if err == nil {
		t.Fatalf("expected validation error, got nil")
	}
	if !strings.Contains(err.Error(), "VehVIN") {
		t.Fatalf("expected error to mention missing VehVIN, got %v", err)
	}
}

func TestClaimsService_GetSingleClaim_NoResults(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/getclaims/999" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("[]"))
	}))
	t.Cleanup(server.Close)

	client := &ClientSettings{
		AuthToken:  "test-token",
		BaseUrl:    server.URL,
		HTTPClient: server.Client(),
	}

	service := NewClaimsService(client)

	_, err := service.GetSingleClaim(context.Background(), 999)
	if err == nil {
		t.Fatalf("expected error when no claim returned")
	}
}

func TestClaimPostRequiredFields(t *testing.T) {
	t.Parallel()

	expected := []string{
		"RenterName",
		"InsCompaniesID",
		"DateOfLoss",
		"VehMake",
		"VehModel",
		"VehColor",
		"VehVIN",
	}

	got := ClaimPostRequiredFields()
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("unexpected fields: %v", got)
	}

	got[0] = "mutated"
	fresh := ClaimPostRequiredFields()
	if fresh[0] != expected[0] {
		t.Fatalf("expected helper to return copy; got %v", fresh)
	}
}
