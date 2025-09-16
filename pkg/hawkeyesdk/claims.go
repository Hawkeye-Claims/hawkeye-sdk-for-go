package hawkeyesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func (cfg *ClientSettings) CreateClaim(claim ClaimPost) (ApiResponse, error) {
	var apiResp ApiResponse
	jsonData, err := json.Marshal(claim)
	if err != nil {
		return apiResp, fmt.Errorf("failed to marshal claim data: %v", err)
	}

	req, err := http.NewRequest("POST", cfg.BaseUrl+"/createclaim", bytes.NewBuffer(jsonData))
	if err != nil {
		return apiResp, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.AuthToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return apiResp, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&apiResp)
	if err != nil {
		return apiResp, fmt.Errorf("failed to decode response: %v", err)
	}
	return apiResp, nil
}

func (cfg *ClientSettings) UpdateClaim(claim ClaimPost) (ApiResponse, error) {
	var apiResp ApiResponse
	jsonData, err := json.Marshal(claim)
	if err != nil {
		return apiResp, fmt.Errorf("failed to marshal claim data: %v", err)
	}

	req, err := http.NewRequest("POST", cfg.BaseUrl+"/updateclaim", bytes.NewBuffer(jsonData))
	if err != nil {
		return apiResp, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.AuthToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return apiResp, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&apiResp)
	if err != nil {
		return apiResp, fmt.Errorf("failed to decode response: %v", err)
	}
	return apiResp, nil
}

func (cfg *ClientSettings) GetSingleClaim(filenumber int) (Claim, error) {
	var claim Claim
	req, err := http.NewRequest("GET", cfg.BaseUrl+fmt.Sprintf("/getclaims/%d", filenumber), nil)
	if err != nil {
		return claim, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.AuthToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return claim, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&claim)
	if err != nil {
		return claim, fmt.Errorf("failed to decode response: %v", err)
	}
	return claim, nil
}

func (cfg *ClientSettings) GetClaims(includeinactive bool) ([]Claim, error) {
	var claims []Claim
	req, err := http.NewRequest("GET", cfg.BaseUrl+fmt.Sprintf("/getclaims?includeinactive=%t", includeinactive), nil)
	if err != nil {
		return claims, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.AuthToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return claims, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&claims)
	if err != nil {
		return claims, fmt.Errorf("failed to decode response: %v", err)
	}
	return claims, nil
}
