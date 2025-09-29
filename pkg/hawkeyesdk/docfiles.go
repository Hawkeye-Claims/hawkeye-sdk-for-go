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

type DocFilesService struct {
	client *ClientSettings
}

func NewDocFilesService(client *ClientSettings) *DocFilesService {
	client.ensureHTTPClient()
	return &DocFilesService{client: client}
}

func (s *DocFilesService) UploadFile(filenumber int, fileurl string, opts ...UploadFileOption) (ApiResponse, error) {
	options := uploadFileOptions{
		category:        DEFAULT,
		visibleToClient: false,
		notes:           "",
	}
	for _, opt := range opts {
		opt(&options)
	}

	type postData struct {
		Filenumber      int    `json:"filenumber"`
		Link            string `json:"link"`
		Category        string `json:"category"`
		VisibleToClient bool   `json:"visible_to_client"`
		Notes           string `json:"notes"`
	}

	payload := postData{
		Filenumber:      filenumber,
		Link:            fileurl,
		Category:        options.category.String(),
		VisibleToClient: options.visibleToClient,
		Notes:           options.notes,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return ApiResponse{}, fmt.Errorf("failed to marshal post data: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, s.client.BaseUrl+"/savefile", bytes.NewBuffer(jsonData))
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

	if _, err := io.ReadAll(resp.Body); err != nil {
		return ApiResponse{}, fmt.Errorf("failed to read response body: %w", err)
	}

	if err := checkResponse(resp); err != nil {
		return ApiResponse{}, err
	}

	return ApiResponse{
		Message: "File uploaded successfully",
		Success: true,
	}, nil
}
