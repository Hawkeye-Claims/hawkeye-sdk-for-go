package hawkeyesdk

import (
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

	resp, err := service.CreateClaim(t.Context(), ClaimPost{
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

	_, err := service.CreateClaim(t.Context(), ClaimPost{
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

	_, err := service.CreateClaim(t.Context(), ClaimPost{
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

	_, err := service.GetSingleClaim(t.Context(), 999)
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

func TestAdminClaimUnmarshalJSON_SanitizesVehMileage(t *testing.T) {
	t.Parallel()

	var claim AdminClaim
	if err := json.Unmarshal([]byte(`{"vehmileage":"12.345.678"}`), &claim); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if claim.VehMileage != 12345678 {
		t.Fatalf("unexpected veh mileage: %d", claim.VehMileage)
	}

	if err := json.Unmarshal([]byte(`{"vehmileage":987654}`), &claim); err != nil {
		t.Fatalf("expected no error for numeric mileage, got %v", err)
	}
	if claim.VehMileage != 987654 {
		t.Fatalf("unexpected numeric veh mileage: %d", claim.VehMileage)
	}

	for raw, want := range map[string]int{`45.000`: 45000, `45000.5`: 450005, `"45000."`: 45000} {
		if err := json.Unmarshal([]byte(`{"vehmileage":`+raw+`}`), &claim); err != nil {
			t.Fatalf("expected no error for mileage %s, got %v", raw, err)
		}
		if claim.VehMileage != want {
			t.Fatalf("mileage %s: expected %d, got %d", raw, want, claim.VehMileage)
		}
	}
}

func TestAdminClaimUnmarshalJSON_DecodesHCAdjusterEmail(t *testing.T) {
	t.Parallel()

	var claim AdminClaim
	payload := `{"hcadjuster":"Adjuster Name","hcadjusteremail":"adjuster@hawkeyeclaims.test"}`
	if err := json.Unmarshal([]byte(payload), &claim); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if claim.HCAjusterEmail != "adjuster@hawkeyeclaims.test" {
		t.Fatalf("unexpected hc adjuster email: %q", claim.HCAjusterEmail)
	}
}

func TestAdminClaimUnmarshalJSON_DocTypes(t *testing.T) {
	t.Parallel()

	var claim AdminClaim
	payload := `{"docfiles":[
		{"doctype":"Incident Report (ACORD)","filename":"acord.pdf"},
		{"doctype":"Incident Report","filename":"incident.pdf"},
		{"doctype":"A Category Hawk Added Later","filename":"new.pdf"}
	]}`
	if err := json.Unmarshal([]byte(payload), &claim); err != nil {
		t.Fatalf("expected unknown doc types not to fail the claim, got %v", err)
	}
	want := []DocType{INCIDENT_REPORT_ACORD, INCIDENT_REPORT, DEFAULT}
	if len(claim.DocFiles) != len(want) {
		t.Fatalf("expected %d doc files, got %d", len(want), len(claim.DocFiles))
	}
	for i, doc := range claim.DocFiles {
		if doc.Doctype != want[i] {
			t.Fatalf("doc %d: expected %v, got %v", i, want[i], doc.Doctype)
		}
	}
	if INCIDENT_REPORT_ACORD.String() != "Incident Report (ACORD)" {
		t.Fatalf("unexpected ACORD name: %q", INCIDENT_REPORT_ACORD.String())
	}
	if FINAL_INVOICE != 57 {
		t.Fatalf("existing doc type values must not shift; FINAL_INVOICE = %d", FINAL_INVOICE)
	}
}
