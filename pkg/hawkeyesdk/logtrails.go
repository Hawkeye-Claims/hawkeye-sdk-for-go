package hawkeyesdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type LogTrailOption func(*logTrailOptions)

type logTrailOptions struct {
	date string
}

func WithDate(date string) LogTrailOption {
	return func(opts *logTrailOptions) {
		opts.date = date
	}
}

type LogTrailsService struct {
	client *ClientSettings
}

func NewLogTrailsService(client *ClientSettings) *LogTrailsService {
	client.ensureHTTPClient()
	return &LogTrailsService{client: client}
}

func (s *LogTrailsService) CreateLogTrail(ctx context.Context, filenumber int, activity string, opts ...LogTrailOption) (ApiResponse, error) {
	options := logTrailOptions{
		date: time.Now().Format("01/02/2006"),
	}
	for _, opt := range opts {
		opt(&options)
	}

	type postData struct {
		Filenumber int    `json:"filenumber"`
		Activity   string `json:"activity"`
		Date       string `json:"date"`
	}

	payload := postData{
		Filenumber: filenumber,
		Activity:   activity,
		Date:       options.date,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return ApiResponse{}, fmt.Errorf("failed to marshal post data: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.client.BaseUrl+"/createLogTrailEntry", bytes.NewBuffer(jsonData))
	if err != nil {
		return ApiResponse{}, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.client.AuthToken))

	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return ApiResponse{}, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return ApiResponse{}, fmt.Errorf("failed to read response body: %w", err)
	}

	if err := checkResponse(resp); err != nil {
		return ApiResponse{}, err
	}

	var apiResp ApiResponse
	if err := json.Unmarshal(bodyBytes, &apiResp); err != nil {
		return ApiResponse{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return apiResp, nil
}
