package toggl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	USER_AGENT   string = "https://github.com/nikoheikkila/hours"
	BASE_URL     string = "https://api.track.toggl.com/reports/api/v2"
	CONTENT_TYPE string = "application/json"
	PASSWORD     string = "api_token"
)

var (
	Client           HTTPClient
	ErrMissingToken  error = errors.New("missing Toggl API token for the client")
	ErrEmptyResponse error = errors.New("empty response from Toggl API")
)

type HTTPClient interface {
	Do(request *http.Request) (*http.Response, error)
}

type TogglClient struct {
	baseURL     string
	contentType string
	token       string
	workspace   string
}

type ErrorResponse struct {
	Error struct {
		Message string `json:"message"`
		Tip     string `json:"tip"`
		Code    int    `json:"code"`
	} `json:"error"`
}

func init() {
	Client = &http.Client{}
}

func New(token, workspace string) *TogglClient {
	return &TogglClient{
		baseURL:     BASE_URL,
		contentType: CONTENT_TYPE,
		token:       token,
		workspace:   workspace,
	}
}

func (c *TogglClient) Entries(start, end string) ([]TimeEntry, error) {
	startDate := url.QueryEscape(start)
	endDate := url.QueryEscape(end)
	url := fmt.Sprintf("%s/details?workspace_id=%s&since=%s&until=%s&user_agent=%s", c.baseURL, c.workspace, startDate, endDate, USER_AGENT)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := c.sendRequest(request)
	if err != nil {
		return nil, err
	}

	var report DetailedReport

	err = json.Unmarshal(bytes, &report)
	if err != nil {
		return nil, err
	}

	if len(report.Data) == 0 {
		return nil, ErrEmptyResponse
	}

	return report.Data, nil
}

func (c *TogglClient) sendRequest(request *http.Request) ([]byte, error) {
	if c.token == "" {
		return nil, ErrMissingToken
	}

	request.SetBasicAuth(c.token, PASSWORD)
	request.Header.Add("User-Agent", USER_AGENT)
	request.Header.Add("Content-Type", c.contentType)

	response, err := Client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse

		err := json.Unmarshal(body, &errorResponse)
		if err != nil {
			return nil, fmt.Errorf("request failed with status %d and unschematic error response: %s", response.StatusCode, body)
		}

		code := errorResponse.Error.Code
		message := fmt.Sprintf("%s %s", errorResponse.Error.Message, errorResponse.Error.Tip)
		return nil, fmt.Errorf("request failed with status %d and message: '%s'", code, message)
	}

	return body, nil
}
