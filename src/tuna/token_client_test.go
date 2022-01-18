package tuna

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

//NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

func TestClient_BindCVV(t *testing.T) {
	t.Run("should return bind succeed", func(t *testing.T) {
		client := NewTestClient(func(req *http.Request) *http.Response {
			// Test request parameters
			assert.Equal(t, req.URL.String(), "/api/Token/Bind")
			return &http.Response{
				StatusCode: 200,
				// Send response to be tested
				Body: ioutil.NopCloser(bytes.NewBufferString(`{
					"code": 1,
					"message": "Bind succeeded"
				}`)),
				// Must be set to non-nil value or it panics
				Header: make(http.Header),
			}
		})
		api := NewTokenClient(client, Config{})

		req := BindCVVRequest{
			Token:     "test",
			SessionID: "test",
			CVV:       "test",
		}
		res, err := api.BindCVV(req)
		assert.NoError(t, err)
		assert.Equal(t, &BindCVVResponse{
			1,
			"Bind succeeded",
		}, res)
	})

	t.Run("should return session object invalid", func(t *testing.T) {
		client := NewTestClient(func(req *http.Request) *http.Response {
			// Test request parameters
			assert.Equal(t, req.URL.String(), "/api/Token/Bind")
			return &http.Response{
				StatusCode: 200,
				// Send response to be tested
				Body: ioutil.NopCloser(bytes.NewBufferString(`{
					"code": -1,
					"message": "Session object is invalid"
				}`)),
				// Must be set to non-nil value or it panics
				Header: make(http.Header),
			}
		})
		api := NewTokenClient(client, Config{})

		req := BindCVVRequest{
			Token:     "test",
			SessionID: "test",
			CVV:       "test",
		}
		res, err := api.BindCVV(req)
		assert.NoError(t, err)
		assert.Equal(t, &BindCVVResponse{
			Code:    -1,
			Message: "Session object is invalid",
		}, res)
	})

	t.Run("should return a http error bad request", func(t *testing.T) {
		client := NewTestClient(func(req *http.Request) *http.Response {
			// Test request parameters
			assert.Equal(t, req.URL.String(), "/api/Token/Bind")
			return &http.Response{
				StatusCode: http.StatusBadRequest,
				// Send response to be tested
				Body: ioutil.NopCloser(bytes.NewBufferString(`{
					"type": "https://tools.ietf.org/html/rfc7231#section-6.5.1",
					"title": "One or more validation errors occurred.",
					"status": 400,
					"traceId": "|9116a0f7-40600755b7d63ba3.",
					"errors": {
						"": [
							"A non-empty request body is required."
						],
						"value": [
							"The value field is required."
						]
					}
				}`)),
				// Must be set to non-nil value or it panics
				Header: make(http.Header),
			}
		})
		api := NewTokenClient(client, Config{})

		req := BindCVVRequest{}
		res, err := api.BindCVV(req)
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}
