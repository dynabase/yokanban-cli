package http_test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	yohttp "yokanban-cli/internal/http"
)

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

func TestHTTPAuthSuccess(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, "https://api.yokanban.io/auth/oauth2/token", req.URL.String())

		// Test request headers
		assert.Equal(t, "application/x-www-form-urlencoded", req.Header.Get("Content-Type"))

		// Test form values
		assert.Equal(t, "urn:ietf:params:oauth:grant-type:jwt-bearer", req.FormValue("grant_type"))
		assert.Equal(t, "foo.bar.123", req.FormValue("assertion"))

		token := yohttp.TokenData{ExpiresIn: 111, Scope: "test", AccessToken: "123"}
		res := yohttp.TokenResponse{Data: token}
		body, _ := json.Marshal(res)

		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewReader(body)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})

	h := yohttp.HTTP{Client: client}
	token, err := h.Auth("foo.bar.123")

	assert.Nil(t, err)
	assert.Equal(t, "123", token.AccessToken)
	assert.Equal(t, "test", token.Scope)
	assert.Equal(t, 111, token.ExpiresIn)
}

func TestHTTPAuthFail(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 500,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`Fail`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})

	h := yohttp.HTTP{Client: client}
	_, err := h.Auth("foo.bar.123")

	assert.NotNil(t, err)
}

func TestHTTPGetSuccess(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, "https://api.yokanban.io/foo/bar", req.URL.String())

		// Test request headers
		assert.Equal(t, "Bearer abc123", req.Header.Get("Authorization"))

		// Test request method
		assert.Equal(t, http.MethodGet, req.Method)

		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`OK`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})

	h := yohttp.HTTP{Client: client}
	body, err := h.Get("/foo/bar", "abc123")

	assert.Nil(t, err)
	assert.Equal(t, "OK", body)
}

func TestHTTPPostSuccess(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, "https://api.yokanban.io/foo/bar", req.URL.String())

		// Test request headers
		assert.Equal(t, "Bearer abc123", req.Header.Get("Authorization"))

		// Test request method
		assert.Equal(t, http.MethodPost, req.Method)

		// Test request body
		buf := new(bytes.Buffer)
		buf.ReadFrom(req.Body)
		body := buf.String()

		assert.Equal(t, "{\"foo\": 123}", body)

		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`OK`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})

	h := yohttp.HTTP{Client: client}
	body, err := h.Post("/foo/bar", "abc123", "{\"foo\": 123}")

	assert.Nil(t, err)
	assert.Equal(t, "OK", body)
}

func TestHTTPPatchSuccess(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, "https://api.yokanban.io/foo/bar", req.URL.String())

		// Test request headers
		assert.Equal(t, "Bearer abc123", req.Header.Get("Authorization"))

		// Test request method
		assert.Equal(t, http.MethodPatch, req.Method)

		// Test request body
		buf := new(bytes.Buffer)
		buf.ReadFrom(req.Body)
		body := buf.String()

		assert.Equal(t, "{\"foo\": 123}", body)

		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`OK`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})

	h := yohttp.HTTP{Client: client}
	body, err := h.Patch("/foo/bar", "abc123", "{\"foo\": 123}")

	assert.Nil(t, err)
	assert.Equal(t, "OK", body)
}
