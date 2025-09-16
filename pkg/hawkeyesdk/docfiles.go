package hawkeyesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (cfg *ClientSettings) UploadFile(filenumber int, fileurl string, category DocType, visibleToClient bool, notes string) error {
	type PostData struct {
		Filenumber      int    `json:"filenumber"`
		Link            string `json:"link"`
		Category        string `json:"category"`
		VisibleToClient bool   `json:"visible_to_client"`
		Notes           string `json:"notes"`
	}
	postData := PostData{
		Filenumber:      filenumber,
		Link:            fileurl,
		Category:        category.String(),
		VisibleToClient: visibleToClient,
		Notes:           notes,
	}
	var apiResp ApiResponse

	jsonData, err := json.Marshal(postData)
	if err != nil {
		return fmt.Errorf("failed to marshal post data: %v", err)
	}

	req, err := http.NewRequest("POST", cfg.BaseUrl+"/savefile", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.AuthToken))

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&apiResp)
	if err != nil {
		return fmt.Errorf("failed to decode response: %v", err)
	}
	return nil
}
