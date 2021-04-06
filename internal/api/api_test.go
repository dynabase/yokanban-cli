package api_test

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"testing"
	"yokanban-cli/internal/accesstoken"
	"yokanban-cli/internal/api"
)

func TestAPI_Test_Success(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	gock.New("https://api.yokanban.io").
		MatchHeader("Authorization", "Bearer mock-token").
		Get("/auth/oauth2/test").
		Reply(200).
		BodyString("Ok")

	tokenMock := new(accesstoken.Mock)
	tokenMock.On("Get").Return("mock-token")

	a := api.API{AccessToken: tokenMock}
	res := a.Test()
	assert.Equal(t, "Ok", res)
	assert.Equal(t, true, gock.IsDone())
}

func TestAPI_Test_Retry_Success(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	gock.New("https://api.yokanban.io").
		MatchHeader("Authorization", "Bearer mock-token").
		Get("/auth/oauth2/test").
		Reply(500)

	gock.New("https://api.yokanban.io").
		MatchHeader("Authorization", "Bearer mock-refresh-token").
		Get("/auth/oauth2/test").
		Reply(200).
		BodyString("Ok")

	tokenMock := new(accesstoken.Mock)
	tokenMock.On("Get").Return("mock-token")
	tokenMock.On("Refresh").Return("mock-refresh-token")

	a := api.API{AccessToken: tokenMock}
	res := a.Test()
	assert.Equal(t, "Ok", res)
	assert.Equal(t, true, gock.IsDone())
}
