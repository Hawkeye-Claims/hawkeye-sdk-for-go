package hawkeyesdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

type GetClaimsOption func(*getClaimsOptions)

type getClaimsOptions struct {
	includeInactive bool
}

func WithIncludeInactive(include bool) GetClaimsOption {
	return func(opts *getClaimsOptions) {
		opts.includeInactive = include
	}
}

func (cfg *ClientSettings) CreateClaim(ctx context.Context, claim ClaimPost) (ApiResponse, error) {
	var apiResp ApiResponse
	jsonData, err := json.Marshal(claim)
	if err != nil {
		return apiResp, fmt.Errorf("failed to marshal claim data: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", cfg.BaseUrl+"/createclaim", bytes.NewBuffer(jsonData))
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

func (cfg *ClientSettings) UpdateClaim(ctx context.Context, claim ClaimPost) (ApiResponse, error) {
	var apiResp ApiResponse
	jsonData, err := json.Marshal(claim)
	if err != nil {
		return apiResp, fmt.Errorf("failed to marshal claim data: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", cfg.BaseUrl+"/updateclaim", bytes.NewBuffer(jsonData))
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

func (cfg *ClientSettings) GetSingleClaim(ctx context.Context, filenumber int) (Claim, error) {
	var claims []Claim
	req, err := http.NewRequest("GET", cfg.BaseUrl+fmt.Sprintf("/getclaims/%d", filenumber), nil)
	if err != nil {
		return Claim{}, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.AuthToken))

	resp, err := cfg.HTTPClient.Do(req)
	if err != nil {
		return Claim{}, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return Claim{}, fmt.Errorf("failed to read response body: %v", err)
	}

	err = checkResponse(resp)
	if err != nil {
		return Claim{}, err
	}

	err = json.Unmarshal(bodyBytes, &claims)
	if err != nil {
		return Claim{}, fmt.Errorf("failed to decode response: %v", err)
	}

	if len(claims) == 0 {
		return Claim{}, fmt.Errorf("no claim found with filenumber %d", filenumber)
	}

	return claims[0], nil
}

func (cfg *ClientSettings) GetClaims(ctx context.Context, opts ...GetClaimsOption) ([]Claim, error) {
	options := getClaimsOptions{
		includeInactive: false,
	}
	for _, opt := range opts {
		opt(&options)
	}
	var claims []Claim
	req, err := http.NewRequest("GET", cfg.BaseUrl+fmt.Sprintf("/getclaims/all/%t", options.includeInactive), nil)
	if err != nil {
		return claims, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.AuthToken))

	resp, err := cfg.HTTPClient.Do(req)
	if err != nil {
		return claims, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return claims, fmt.Errorf("failed to read response body: %v", err)
	}

	err = checkResponse(resp)
	if err != nil {
		return claims, err
	}

	err = json.Unmarshal(bodyBytes, &claims)
	if err != nil {
		return claims, fmt.Errorf("failed to decode response: %v", err)
	}
	return claims, nil
}
