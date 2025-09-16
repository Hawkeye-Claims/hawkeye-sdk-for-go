package hawkeyesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (cfg *ClientSettings) CreateLogTrail(filenumber int, activity string, date string) (ApiResponse, error) {
	type PostData struct {
		Filenumber int    `json:"filenumber"`
		Activity   string `json:"activity"`
		Date       string `json:"date"`
	}
	postData := PostData{
		Filenumber: filenumber,
		Activity:   activity,
		Date:       date,
	}
	var apiResp ApiResponse

	jsonData, err := json.Marshal(postData)
	if err != nil {
		return apiResp, fmt.Errorf("failed to marshal post data: %v", err)
	}

	req, err := http.NewRequest("POST", cfg.BaseUrl+"/createLogTailEntry", bytes.NewBuffer(jsonData))
	if err != nil {
		return apiResp, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.AuthToken))

	client := http.Client{}
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
