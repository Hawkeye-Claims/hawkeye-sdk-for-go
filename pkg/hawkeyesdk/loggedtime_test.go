package hawkeyesdk

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func startTestServer(t *testing.T) *httptest.Server {
	t.Helper()

	seededLogs := []map[string]any{
		{
			"date":       "2026-01-08",
			"filenumber": 26110431,
			"user":       "Gabriela Sokolow",
			"email":      "gabriela@hawkeyeclaims.com",
			"loggedtime": 0.1,
		},
		{
			"date":       "2026-01-09",
			"filenumber": 26110657,
			"user":       "Kevin Lopes",
			"email":      "kevin@hawkeyeclaims.com",
			"loggedtime": 0.2,
		},
		{
			"date":       "2026-01-10",
			"filenumber": 26110823,
			"user":       "Louise Johnson",
			"loggedtime": 0.3,
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/getloggedtime" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer test-token" {
			t.Fatalf("unexpected auth header: %s", got)
		}
		queriedDateFrom := r.URL.Query().Get("datefrom")
		if queriedDateFrom == "" {
			t.Fatalf("missing datefrom query parameter")
		}
		dateFrom, err := time.Parse("2006-01-02", queriedDateFrom)
		if err != nil {
			t.Fatalf("invalid datefrom format: %v", err)
		}
		queriedDateTo := r.URL.Query().Get("dateto")
		if queriedDateTo == "" {
			t.Fatalf("missing dateto query parameter")
		}
		dateTo, err := time.Parse("2006-01-02", queriedDateTo)
		if err != nil {
			t.Fatalf("invalid dateto format: %v", err)
		}
		if dateTo.After(time.Now()) {
			t.Fatalf("dateto cannot be in the future: %s", queriedDateTo)
		}
		responseLogs := []map[string]any{}
		for _, log := range seededLogs {
			dateValue, ok := log["date"].(string)
			if !ok {
				t.Fatalf("seeded log has an invalid date: %v", log["date"])
			}
			seededDate, err := time.Parse("2006-01-02", dateValue)
			if err != nil {
				t.Fatalf("seeded log has an invalid date: %v", err)
			}
			if !dateFrom.After(seededDate) && !dateTo.Before(seededDate) {
				responseLogs = append(responseLogs, log)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(responseLogs)
	}))
	t.Cleanup(func() { server.Close() })

	return server
}

func TestLoggedTimeService_GetLoggedTime(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		dateFrom      time.Time
		dateTo        time.Time
		expectedCount int
		wantFirst     LoggedTime
	}{
		{
			name:          "full range",
			dateFrom:      time.Date(2026, 1, 7, 0, 0, 0, 0, time.UTC),
			dateTo:        time.Date(2026, 1, 11, 0, 0, 0, 0, time.UTC),
			expectedCount: 3,
			wantFirst: LoggedTime{
				Date:       "2026-01-08",
				Filenumber: 26110431,
				User:       "Gabriela Sokolow",
				Email:      "gabriela@hawkeyeclaims.com",
				Time:       0.1,
			},
		},
		{
			name:          "date in future defaults to today",
			dateFrom:      time.Date(2026, 1, 7, 0, 0, 0, 0, time.UTC),
			dateTo:        time.Now().AddDate(0, 0, 1),
			expectedCount: 3,
			wantFirst: LoggedTime{
				Date:       "2026-01-08",
				Filenumber: 26110431,
				User:       "Gabriela Sokolow",
				Email:      "gabriela@hawkeyeclaims.com",
				Time:       0.1,
			},
		},
		{
			name:          "partial range",
			dateFrom:      time.Date(2026, 1, 9, 0, 0, 0, 0, time.UTC),
			dateTo:        time.Date(2026, 1, 11, 0, 0, 0, 0, time.UTC),
			expectedCount: 2,
			wantFirst: LoggedTime{
				Date:       "2026-01-09",
				Filenumber: 26110657,
				User:       "Kevin Lopes",
				Email:      "kevin@hawkeyeclaims.com",
				Time:       0.2,
			},
		},
	}

	server := startTestServer(t)

	client := &ClientSettings{
		AuthToken:  "test-token",
		BaseUrl:    server.URL,
		HTTPClient: server.Client(),
	}

	service := NewLoggedTimeService(client)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.GetLoggedTime(t.Context(), tt.dateFrom, tt.dateTo)
			if err != nil {
				t.Fatalf("GetLoggedTime() error = %v", err)
			}
			if len(got) != tt.expectedCount {
				t.Errorf("GetLoggedTime() got = %v, want %v", len(got), tt.expectedCount)
			}
			if len(got) > 0 && tt.wantFirst != (LoggedTime{}) && got[0] != tt.wantFirst {
				t.Errorf("GetLoggedTime() first record = %+v, want %+v", got[0], tt.wantFirst)
			}
		})
	}
}
