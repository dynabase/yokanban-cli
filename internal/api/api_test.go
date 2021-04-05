package api_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gopkg.in/h2non/gock.v1"
	"testing"
	"yokanban-cli/internal/api"
)

type AccessToken struct {
	mock.Mock
}

func (m *AccessToken) Get() string {
	return "mock-token"
}

func (m *AccessToken) Refresh() string {
	return "mock-token"
}

func TestTestSuccess(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	gock.New("https://api.yokanban.io").
		MatchHeader("Authorization", "Bearer mock-token").
		Get("/auth/oauth2/test").
		Reply(200).
		BodyString("Ok")

	tokenMock := new(AccessToken)
	a := api.API{AccessToken: tokenMock}

	res := a.Test()
	assert.Equal(t, "Ok", res)
	assert.Equal(t, gock.IsDone(), true)
}

func TestTestRetrySuccess(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	gock.New("https://api.yokanban.io").
		MatchHeader("Authorization", "Bearer mock-token").
		Get("/auth/oauth2/test").
		Reply(500)

	gock.New("https://api.yokanban.io").
		MatchHeader("Authorization", "Bearer mock-token").
		Get("/auth/oauth2/test").
		Reply(200).
		BodyString("Ok")

	tokenMock := new(AccessToken)
	a := api.API{AccessToken: tokenMock}

	res := a.Test()
	assert.Equal(t, "Ok", res)
	assert.Equal(t, gock.IsDone(), true)
}
