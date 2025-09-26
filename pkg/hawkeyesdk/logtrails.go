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

func (cfg *ClientSettings) CreateLogTrail(ctx context.Context, filenumber int, activity string, opts ...LogTrailOption) (ApiResponse, error) {
	options := logTrailOptions{
		date: time.Now().Format("03/31/2025"),
	}
	for _, opt := range opts {
		opt(&options)
	}
	type PostData struct {
		Filenumber int    `json:"filenumber"`
		Activity   string `json:"activity"`
		Date       string `json:"date"`
	}
	postData := PostData{
		Filenumber: filenumber,
		Activity:   activity,
		Date:       options.date,
	}
	var apiResp ApiResponse

	jsonData, err := json.Marshal(postData)
	if err != nil {
		return apiResp, fmt.Errorf("failed to marshal post data: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", cfg.BaseUrl+"/createLogTailEntry", bytes.NewBuffer(jsonData))
	if err != nil {
		return apiResp, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.AuthToken))

	resp, err := cfg.HTTPClient.Do(req)
	if err != nil {
		return apiResp, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return apiResp, fmt.Errorf("failed to read response body: %v", err)
	}

	err = checkResponse(resp)
	if err != nil {
		return apiResp, err
	}

	err = json.Unmarshal(bodyBytes, &apiResp)
	if err != nil {
		return apiResp, fmt.Errorf("failed to decode response: %v", err)
	}
	return apiResp, nil
}
