package mocks

import "net/http"

var (
	GetDoFunc func(request *http.Request) (*http.Response, error)
)

type MockClient struct {
	DoFunc func(request *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(request *http.Request) (*http.Response, error) {
	return GetDoFunc(request)
}
