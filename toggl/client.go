package toggl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	USER_AGENT string = "github.com/nikoheikkila/hours"
	BASE_URL string = "https://api.track.toggl.com/api/v8"
	CONTENT_TYPE string = "application/json"
	PASSWORD string = "api_token"
)

var (
	Client HTTPClient
)

var (
	ErrMissingToken error = errors.New("missing Toggl API token for the client")
)

type HTTPClient interface {
	Do(request *http.Request) (*http.Response, error)
}

type TogglClient struct {
	baseURL string
	contentType string
	token string
}

func init() {
	Client = &http.Client{}
}

func WithToken(token string) *TogglClient {
	return &TogglClient{
		baseURL: BASE_URL,
		contentType: CONTENT_TYPE,
		token: token,
	}
}

func (c *TogglClient) Entries(start, end time.Time) ([]TimeEntry, error) {
	startDate := url.QueryEscape(start.Format(time.RFC3339))
	endDate := url.QueryEscape(end.Format(time.RFC3339))
	url := fmt.Sprintf("%s/time_entries?start_date=%s&end_date=%s", c.baseURL, startDate, endDate)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := c.sendRequest(request)
	if err != nil {
		return nil, err
	}

	var entries []TimeEntry

	err = json.Unmarshal(bytes, &entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func (c *TogglClient) sendRequest(request *http.Request) ([]byte, error) {
	if (c.token == "") {
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
		return nil, fmt.Errorf("request failed with status %d: %s", response.StatusCode, body)
	}

	return body, nil
}