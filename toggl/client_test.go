// build +integration

package toggl

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	mocks "github.com/nikoheikkila/hours/toggl/utils"
	"github.com/stretchr/testify/assert"
)

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

	toggl := New("some-token", "1")
	entries, err := toggl.Entries(time.Now(), time.Now())

	assert.Nil(err)
	assert.Equal(2, len(entries))

	assert.EqualValues("Project work", entries[0].Description)
	assert.EqualValues("Product X", entries[0].Project)
	assert.EqualValues(3600000, entries[0].Duration)

	assert.EqualValues("Client meeting", entries[1].Description)
	assert.EqualValues("Product Y", entries[1].Project)
	assert.EqualValues(1800000, entries[1].Duration)
}

func TestClientReturnErrorForMissingToken(t *testing.T) {
	assert := assert.New(t)

	client := New("", "1")
	entries, err := client.Entries(time.Now(), time.Now())

	assert.Nil(entries)
	assert.EqualError(err, "missing Toggl API token for the client")
}