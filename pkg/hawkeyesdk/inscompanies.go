package hawkeyesdk

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type GetInsCompaniesOptions func(*getInsCompaniesOptions)

type getInsCompaniesOptions struct {
	query string
	limit int
}

func WithQueryParameters(query string, limit int) GetInsCompaniesOptions {
	return func(opts *getInsCompaniesOptions) {
		opts.query = query
		if limit > 0 {
			opts.limit = limit
		} else {
			opts.limit = 5
		}
	}
}

type InsCompaniesService struct {
	client *ClientSettings
}

func NewInsCompaniesService(client *ClientSettings) *InsCompaniesService {
	client.ensureHTTPClient()
	return &InsCompaniesService{client: client}
}

func (s *InsCompaniesService) GetInsuranceCompanies(ctx context.Context, opts ...GetInsCompaniesOptions) ([]InsCompany, error) {
	const MAX_LIMIT = 20
	options := &getInsCompaniesOptions{
		limit: 5,
	}

	for _, opt := range opts {
		opt(options)
	}

	if options.limit > MAX_LIMIT {
		options.limit = MAX_LIMIT
	}

	u, _ := url.Parse(s.client.BaseUrl + "/inscompanies")

	if options.query != "" {
		queryParams := url.Values{}
		queryParams.Add("q", options.query)
		queryParams.Add("limit", strconv.Itoa(options.limit))
		u.RawQuery = queryParams.Encode()
	}

	var insCompanies []InsCompany

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return insCompanies, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.client.AuthToken))

	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return insCompanies, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return insCompanies, fmt.Errorf("failed to read response body: %w", err)
	}

	if err := checkResponse(resp); err != nil {
		return insCompanies, err
	}

	type fullResponse struct {
		Data []InsCompany `json:"data"`
	}

	type queryResponse struct {
		Query       string       `json:"query"`
		Suggestions []InsCompany `json:"suggestions"`
	}

	var genericResp map[string]any

	if err := json.Unmarshal(bodyBytes, &genericResp); err != nil {
		return insCompanies, fmt.Errorf("failed to decode response: %w", err)
	}

	if _, ok := genericResp["data"]; ok {
		var fr fullResponse
		if err := json.Unmarshal(bodyBytes, &fr); err != nil {
			return insCompanies, fmt.Errorf("failed to decode full response: %w", err)
		}
		insCompanies = fr.Data
	} else if _, ok := genericResp["suggestions"]; ok {
		var qr queryResponse
		if err := json.Unmarshal(bodyBytes, &qr); err != nil {
			return insCompanies, fmt.Errorf("failed to decode query response: %w", err)
		}
		insCompanies = qr.Suggestions
	} else {
		return insCompanies, fmt.Errorf("unexpected response format")
	}

	return insCompanies, nil
}
