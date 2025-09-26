package hawkeyesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("Api returned status code %d: %s", e.StatusCode, e.Message)
}

func checkResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read error response body: %v", err)
	}

	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	var apiResp ApiResponse

	err = json.Unmarshal(bodyBytes, &apiResp)
	if err != nil {
		return fmt.Errorf("failed to decode response: %v", err)
	}

	return &APIError{
		StatusCode: resp.StatusCode,
		Message:    apiResp.Message,
	}
}
