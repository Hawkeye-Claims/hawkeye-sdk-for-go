package hawkeyesdk

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type LoggedTimeService struct {
	client *ClientSettings
}

func NewLoggedTimeService(client *ClientSettings) *LoggedTimeService {
	client.ensureHTTPClient()
	return &LoggedTimeService{client: client}
}

func (s *LoggedTimeService) GetLoggedTime(ctx context.Context, dateFrom, dateTo time.Time) ([]LoggedTime, error) {
	var maxTime = time.Now()
	u, err := url.Parse(s.client.BaseUrl + "/getloggedtime")
	if err != nil {
		return nil, err
	}

	queryParams := url.Values{}
	queryParams.Add("datefrom", dateFrom.Format(time.DateOnly))
	if dateTo.After(maxTime) {
		dateTo = maxTime
	}
	queryParams.Add("dateto", dateTo.Format(time.DateOnly))
	u.RawQuery = queryParams.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("get logged time: create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.client.AuthToken))

	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("get logged time: request failed: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("get logged time: read response body: %w", err)
	}

	if err := checkResponse(resp); err != nil {
		return nil, err
	}

	var loggedTimes []LoggedTime

	if err := json.Unmarshal(bodyBytes, &loggedTimes); err != nil {
		return nil, fmt.Errorf("get logged time: decode response: %w", err)
	}

	return loggedTimes, nil
}
