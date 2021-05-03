package api_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"testing"
	"yokanban-cli/internal/accesstoken"
	"yokanban-cli/internal/api"
)

func TestAPI_CreateCard(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	id := uuid.New()

	gock.New("https://api.yokanban.io").
		MatchHeader("Authorization", "Bearer mock-token").
		Patch("/boards/606a34fc20904527edaf3790").
		MatchType("json").
		BodyString(`{
		  "event": {
			"type": "BULK",
			"events": [
			  {
				"type": "ADD",
				"elementId": "` + id.String() + `",
				"oldValues": null,
				"newValues": {
				  "type": "CARD",
				  "title": "test-card",
				  "shape": {
					"x": 497.5,
					"y": 438,
					"width": 250,
					"height": 147
				  },
				  "color": "#FFFFFF",
				  "isArchived": false,
				  "id": "` + id.String() + `",
				  "zIndex": 1000
				},
				"softwareVersion": "0.4.12"
			  }
			],
			"softwareVersion": "0.4.12"
		  }
		}
		`).
		Reply(200).
		JSON(api.CreateCardResponse{
			Success: true,
			Data:    append([]api.CreateCardResponseDetails{}, api.CreateCardResponseDetails{OldID: "abc", NewID: "def"}),
		})

	tokenMock := new(accesstoken.Mock)
	tokenMock.On("Get").Return("mock-token")

	a := api.API{AccessToken: tokenMock}
	res := a.CreateCard("606a34fc20904527edaf3790", "test-card", id)
	assert.Equal(t, "abc", res[0].OldID)
	assert.Equal(t, "def", res[0].NewID)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, true, gock.IsDone())
}
