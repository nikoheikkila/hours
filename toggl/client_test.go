package toggl

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	mocks "github.com/nikoheikkila/hours/toggl/utils"
	"github.com/stretchr/testify/assert"
)

const TEST_DATE string = "2021-01-01"

func init() {
	Client = &mocks.MockClient{}
}

func TestClientReturnEntriesForTimeRange(t *testing.T) {
	assert := assert.New(t)
	json := `{
		"data": [
			{
				"description": "Project work",
				"start": "2021-03-13T11:00:11+02:00",
				"end": "2021-03-13T14:21:17+02:00",
				"dur": 3600000,
				"client": "Client A",
				"project": "Product X",
				"project_hex_color": "#525266"
			},
			{
				"description": "Client meeting",
				"start": "2021-03-12T16:00:00+02:00",
				"end": "2021-03-12T18:30:00+02:00",
				"dur": 1800000,
				"client": "Client B",
				"project": "Product Y",
				"project_hex_color": "#06a893"
			}
		]
	}`

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		reader := ioutil.NopCloser(bytes.NewReader([]byte(json)))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       reader,
		}, nil
	}

	toggl := New(&Configuration{"some-token", "1"})
	entries, err := toggl.Entries(TEST_DATE, TEST_DATE)

	assert.Nil(err)
	assert.Len(entries, 2)

	assert.EqualValues("Project work", entries[0].GetDescription())
	assert.EqualValues("Client A", entries[0].GetClient())
	assert.EqualValues("Product X", entries[0].GetProject())
	assert.EqualValues(1.0, entries[0].GetHours())

	assert.EqualValues("Client meeting", entries[1].GetDescription())
	assert.EqualValues("Client B", entries[1].GetClient())
	assert.EqualValues("Product Y", entries[1].GetProject())
	assert.EqualValues(0.5, entries[1].GetHours())
}

func TestMissingValuesFromClientAreReplaced(t *testing.T) {
	assert := assert.New(t)
	json := `{
		"data": [
			{
				"description": "",
				"start": "2021-03-13T11:00:11+02:00",
				"end": "2021-03-13T14:21:17+02:00",
				"dur": 3600000,
				"client": "",
				"project": ""
			}
		]
	}`

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		reader := ioutil.NopCloser(bytes.NewReader([]byte(json)))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       reader,
		}, nil
	}

	toggl := New(&Configuration{"some-token", "1"})
	entries, err := toggl.Entries(TEST_DATE, TEST_DATE)

	assert.Nil(err)
	assert.Len(entries, 1)

	assert.EqualValues(NO_DESCRIPTION, entries[0].GetDescription())
	assert.EqualValues(NO_CLIENT, entries[0].GetClient())
	assert.EqualValues(NO_PROJECT, entries[0].GetProject())
	assert.EqualValues(1.0, entries[0].GetHours())
}

func TestClientReturnsErrorOnEmptyAPIResponse(t *testing.T) {
	assert := assert.New(t)
	json := `{}`

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		reader := ioutil.NopCloser(bytes.NewReader([]byte(json)))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       reader,
		}, nil
	}

	toggl := New(&Configuration{"some-token", "1"})
	entries, err := toggl.Entries(TEST_DATE, TEST_DATE)

	assert.Nil(entries)
	assert.NotNil(err)
	assert.EqualError(err, ErrEmptyResponse.Error())
}

func TestClientReturnErrorForMissingToken(t *testing.T) {
	assert := assert.New(t)

	toggl := New(&Configuration{"", "1"})
	entries, err := toggl.Entries(TEST_DATE, TEST_DATE)

	assert.Nil(entries)
	assert.EqualError(err, ErrMissingToken.Error())
}
