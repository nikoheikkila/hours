package toggl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	BASE_URL string = "https://api.track.toggl.com/api/v8"
	CONTENT_TYPE string = "application/json"
	PASSWORD string = "api_token"
)

type Client struct {
	baseURL string
	contentType string
	token string
}

func WithToken(token string) *Client {
	return &Client{
		baseURL: BASE_URL,
		contentType: CONTENT_TYPE,
		token: token,
	}
}

func (c *Client) Entries(start, end time.Time) (*[]TimeEntry, error) {
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

	return &entries, nil
}

func (c *Client) sendRequest(request *http.Request) ([]byte, error) {
	request.SetBasicAuth(c.token, PASSWORD)
	client := &http.Client{}

	response, err := client.Do(request)
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