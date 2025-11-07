package hawkeyesdk

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInsCompaniesService_GetInsuranceCompanies_Success(t *testing.T) {
	t.Parallel()

	expectedCompanies := []InsCompany{
		{Id: 1, Name: "State Farm"},
		{Id: 2, Name: "Geico"},
		{Id: 3, Name: "Progressive"},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/inscompanies" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer test-token" {
			t.Fatalf("unexpected auth header: %s", got)
		}
		w.Header().Set("Content-Type", "application/json")
		response := map[string]any{
			"data": expectedCompanies,
		}
		_ = json.NewEncoder(w).Encode(response)
	}))
	t.Cleanup(server.Close)

	client := &ClientSettings{
		AuthToken:  "test-token",
		BaseUrl:    server.URL,
		HTTPClient: server.Client(),
	}

	service := NewInsCompaniesService(client)

	companies, err := service.GetInsuranceCompanies(context.Background())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(companies) != len(expectedCompanies) {
		t.Fatalf("expected %d companies, got %d", len(expectedCompanies), len(companies))
	}

	if companies[0].Name != "State Farm" {
		t.Fatalf("unexpected first company name: %s", companies[0].Name)
	}
}

func TestInsCompaniesService_GetInsuranceCompanies_WithQuery(t *testing.T) {
	t.Parallel()

	expectedSuggestions := []InsCompany{
		{Id: 1, Name: "State Farm", Probability: 95},
		{Id: 4, Name: "Farmers", Probability: 85},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/inscompanies" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		query := r.URL.Query()
		if query.Get("q") != "farm" {
			t.Fatalf("expected query 'farm', got '%s'", query.Get("q"))
		}
		if query.Get("limit") != "5" {
			t.Fatalf("expected limit '5', got '%s'", query.Get("limit"))
		}
		w.Header().Set("Content-Type", "application/json")
		response := map[string]any{
			"query":       "farm",
			"suggestions": expectedSuggestions,
		}
		_ = json.NewEncoder(w).Encode(response)
	}))
	t.Cleanup(server.Close)

	client := &ClientSettings{
		AuthToken:  "test-token",
		BaseUrl:    server.URL,
		HTTPClient: server.Client(),
	}

	service := NewInsCompaniesService(client)

	companies, err := service.GetInsuranceCompanies(context.Background(), WithQueryParameters("farm", 5))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(companies) != len(expectedSuggestions) {
		t.Fatalf("expected %d companies, got %d", len(expectedSuggestions), len(companies))
	}

	if companies[0].Probability != 95 {
		t.Fatalf("unexpected first company name: %s", companies[0].Name)
	}
}

func TestInsCompaniesService_GetInsuranceCompanies_LimitEnforcement(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		limit := query.Get("limit")
		if limit != "20" {
			t.Fatalf("expected limit '20', got '%s'", limit)
		}
		w.Header().Set("Content-Type", "application/json")
		response := map[string]any{
			"query":       "test",
			"suggestions": []InsCompany{},
		}
		_ = json.NewEncoder(w).Encode(response)
	}))
	t.Cleanup(server.Close)

	client := &ClientSettings{
		AuthToken:  "test-token",
		BaseUrl:    server.URL,
		HTTPClient: server.Client(),
	}

	service := NewInsCompaniesService(client)

	_, err := service.GetInsuranceCompanies(context.Background(), WithQueryParameters("test", 100))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestInsCompaniesService_GetInsuranceCompanies_HTTPError(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ApiResponse{Message: "internal server error"})
	}))
	t.Cleanup(server.Close)

	client := &ClientSettings{
		AuthToken:  "test-token",
		BaseUrl:    server.URL,
		HTTPClient: server.Client(),
	}

	service := NewInsCompaniesService(client)

	_, err := service.GetInsuranceCompanies(context.Background())
	if err == nil {
		t.Fatalf("expected error, got none")
	}
}

func TestInsCompaniesService_GetInsuranceCompanies_UnexpectedFormat(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]any{
			"unexpected_field": "unexpected_value",
		}
		_ = json.NewEncoder(w).Encode(response)
	}))
	t.Cleanup(server.Close)

	client := &ClientSettings{
		AuthToken:  "test-token",
		BaseUrl:    server.URL,
		HTTPClient: server.Client(),
	}

	service := NewInsCompaniesService(client)

	_, err := service.GetInsuranceCompanies(context.Background())
	if err == nil {
		t.Fatalf("expected error due to unexpected response format, got none")
	}
}

func TestInsCompaniesService_GetInsuranceCompanies_DefaultLimit(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		limit := query.Get("limit")
		if limit != "5" {
			t.Fatalf("expected default limit '5', got '%s'", limit)
		}
		w.Header().Set("Content-Type", "application/json")
		response := map[string]any{
			"query":       "test",
			"suggestions": []InsCompany{},
		}
		_ = json.NewEncoder(w).Encode(response)
	}))
	t.Cleanup(server.Close)

	client := &ClientSettings{
		AuthToken:  "test-token",
		BaseUrl:    server.URL,
		HTTPClient: server.Client(),
	}

	service := NewInsCompaniesService(client)

	_, err := service.GetInsuranceCompanies(context.Background(), WithQueryParameters("test", 0))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
