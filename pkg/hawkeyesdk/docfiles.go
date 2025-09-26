package hawkeyesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UploadFileOption func(*uploadFileOptions)

type uploadFileOptions struct {
	category        DocType
	visibleToClient bool
	notes           string
}

func WithCategory(category DocType) UploadFileOption {
	return func(opts *uploadFileOptions) {
		opts.category = category
	}
}

func WithVisibleToClient(visible bool) UploadFileOption {
	return func(opts *uploadFileOptions) {
		opts.visibleToClient = visible
	}
}

func WithNotes(notes string) UploadFileOption {
	return func(opts *uploadFileOptions) {
		opts.notes = notes
	}
}

func (cfg *ClientSettings) UploadFile(filenumber int, fileurl string, opts ...UploadFileOption) (ApiResponse, error) {
	options := uploadFileOptions{
		category:        DEFAULT,
		visibleToClient: false,
		notes:           "",
	}
	for _, opt := range opts {
		opt(&options)
	}
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
		Category:        options.category.String(),
		VisibleToClient: options.visibleToClient,
		Notes:           options.notes,
	}
	var apiResp ApiResponse

	jsonData, err := json.Marshal(postData)
	if err != nil {
		return apiResp, fmt.Errorf("failed to marshal post data: %v", err)
	}

	req, err := http.NewRequest("POST", cfg.BaseUrl+"/savefile", bytes.NewBuffer(jsonData))
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

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return apiResp, fmt.Errorf("failed to read response body: %v", err)
	}

	err = checkResponse(resp)
	if err != nil {
		return apiResp, err
	}

	apiResp.Message = "File uploaded successfully"
	apiResp.Success = true

	return apiResp, nil
}
