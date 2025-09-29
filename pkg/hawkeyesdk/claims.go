package hawkeyesdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ClaimPost struct {
	FileNumber         int    `json:"filenumber,omitempty"`
	ClientClaimNo      string `json:"clientclaimno,omitempty"`
	RenterName         string `json:"rentername"`
	RenterPhone        string `json:"renterphone,omitempty"`
	RenterEmail        string `json:"renteremail,omitempty"`
	InsuranceCompany   string `json:"insurancecompany"`
	ClaimNumber        string `json:"claimnumber,omitempty"`
	InsuredName        string `json:"insuredname,omitempty"`
	PolicyNumber       string `json:"policynumber,omitempty"`
	DateOfLoss         string `json:"dateofloss"`
	VehYear            int    `json:"vehyear,omitempty"`
	VehMake            string `json:"vehmake"`
	VehModel           string `json:"vehmodel"`
	VehColor           string `json:"vehcolor"`
	VehVIN             string `json:"vehvin"`
	VehEdition         string `json:"vehedition,omitempty"`
	VehPlateNumber     string `json:"vehplatenumber,omitempty"`
	VehUnitNumber      string `json:"vehunitnumber,omitempty"`
	VehLocationDetails string `json:"vehlocationdetails,omitempty"`
	VehLocationCity    string `json:"vehlocationcity,omitempty"`
	VehLocationState   string `json:"vehlocationstate,omitempty"`
	Note               string `json:"note,omitempty"`
}

var claimPostRequiredFieldDefs = []struct {
	name     string
	accessor func(ClaimPost) string
}{
	{name: "RenterName", accessor: func(c ClaimPost) string { return c.RenterName }},
	{name: "InsuranceCompany", accessor: func(c ClaimPost) string { return c.InsuranceCompany }},
	{name: "DateOfLoss", accessor: func(c ClaimPost) string { return c.DateOfLoss }},
	{name: "VehMake", accessor: func(c ClaimPost) string { return c.VehMake }},
	{name: "VehModel", accessor: func(c ClaimPost) string { return c.VehModel }},
	{name: "VehColor", accessor: func(c ClaimPost) string { return c.VehColor }},
	{name: "VehVIN", accessor: func(c ClaimPost) string { return c.VehVIN }},
}

func ClaimPostRequiredFields() []string {
	fields := make([]string, len(claimPostRequiredFieldDefs))
	for i, def := range claimPostRequiredFieldDefs {
		fields[i] = def.name
	}
	return fields
}

func (c ClaimPost) ValidateForCreate() error {
	var missing []string

	for _, def := range claimPostRequiredFieldDefs {
		if strings.TrimSpace(def.accessor(c)) == "" {
			missing = append(missing, def.name)
		}
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing required fields: %s", strings.Join(missing, ", "))
	}

	return nil
}

type GetClaimsOption func(*getClaimsOptions)

type getClaimsOptions struct {
	includeInactive bool
}

func WithIncludeInactive(include bool) GetClaimsOption {
	return func(opts *getClaimsOptions) {
		opts.includeInactive = include
	}
}

type ClaimsService struct {
	client *ClientSettings
}

func NewClaimsService(client *ClientSettings) *ClaimsService {
	client.ensureHTTPClient()
	return &ClaimsService{client: client}
}

func (s *ClaimsService) CreateClaim(ctx context.Context, claim ClaimPost) (ApiResponse, error) {
	var apiResp ApiResponse

	if err := claim.ValidateForCreate(); err != nil {
		return apiResp, fmt.Errorf("claim validation failed: %w", err)
	}

	jsonData, err := json.Marshal(claim)
	if err != nil {
		return apiResp, fmt.Errorf("failed to marshal claim data: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.client.BaseUrl+"/createclaim", bytes.NewBuffer(jsonData))
	if err != nil {
		return apiResp, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.client.AuthToken))

	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return apiResp, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return apiResp, fmt.Errorf("failed to read response body: %w", err)
	}

	if err := checkResponse(resp); err != nil {
		return apiResp, err
	}

	if err := json.Unmarshal(bodyBytes, &apiResp); err != nil {
		return apiResp, fmt.Errorf("failed to decode response: %w", err)
	}

	return apiResp, nil
}

func (s *ClaimsService) UpdateClaim(ctx context.Context, claim ClaimPost) (ApiResponse, error) {
	var apiResp ApiResponse

	jsonData, err := json.Marshal(claim)
	if err != nil {
		return apiResp, fmt.Errorf("failed to marshal claim data: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.client.BaseUrl+"/updateclaim", bytes.NewBuffer(jsonData))
	if err != nil {
		return apiResp, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.client.AuthToken))

	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return apiResp, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return apiResp, fmt.Errorf("failed to read response body: %w", err)
	}

	if err := checkResponse(resp); err != nil {
		return apiResp, err
	}

	if err := json.Unmarshal(bodyBytes, &apiResp); err != nil {
		return apiResp, fmt.Errorf("failed to decode response: %w", err)
	}

	return apiResp, nil
}

func (s *ClaimsService) GetSingleClaim(ctx context.Context, filenumber int) (Claim, error) {
	var claims []Claim

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, s.client.BaseUrl+fmt.Sprintf("/getclaims/%d", filenumber), nil)
	if err != nil {
		return Claim{}, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.client.AuthToken))

	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return Claim{}, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return Claim{}, fmt.Errorf("failed to read response body: %w", err)
	}

	if err := checkResponse(resp); err != nil {
		return Claim{}, err
	}

	if err := json.Unmarshal(bodyBytes, &claims); err != nil {
		return Claim{}, fmt.Errorf("failed to decode response: %w", err)
	}

	if len(claims) == 0 {
		return Claim{}, fmt.Errorf("no claim found with filenumber %d", filenumber)
	}

	return claims[0], nil
}

func (s *ClaimsService) GetClaims(ctx context.Context, opts ...GetClaimsOption) ([]Claim, error) {
	options := getClaimsOptions{includeInactive: false}
	for _, opt := range opts {
		opt(&options)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, s.client.BaseUrl+fmt.Sprintf("/getclaims/all/%t", options.includeInactive), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.client.AuthToken))

	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if err := checkResponse(resp); err != nil {
		return nil, err
	}

	var claims []Claim
	if err := json.Unmarshal(bodyBytes, &claims); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return claims, nil
}
