package api_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"testing"
	"yokanban-cli/internal/accesstoken"
	"yokanban-cli/internal/api"
)

func TestAPI_CreateColumn_WithoutOtherColumnOnBoard(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	id := uuid.New()

	gock.New("https://api.yokanban.io").
		MatchHeader("Authorization", "Bearer mock-token").
		Get("/boards/606a34fc20904527edaf3790").
		Reply(200).
		JSON(api.GetBoardResponse{
			Success: true,
			Data: api.BoardDetails{
				Avatars:  []api.AvatarDTO{},
				Elements: []api.BoardElement{},
			},
		})

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
				  "type": "COLUMN",
				  "title": "test-col",
				  "shape": {
					"x": 396.5,
					"y": 112.5,
					"width": 350,
					"height": 800
				  },
				  "color": "white",
				  "wipLimit": 0,
				  "isLocked": false,
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
		JSON(api.CreateColumnResponse{
			Success: true,
			Data:    append([]api.CreateColumnResponseDetails{}, api.CreateColumnResponseDetails{OldID: "abc", NewID: "def"}),
		})

	tokenMock := new(accesstoken.Mock)
	tokenMock.On("Get").Return("mock-token")

	a := api.API{AccessToken: tokenMock}
	res := a.CreateColumn("606a34fc20904527edaf3790", "test-col", id)
	assert.Equal(t, "abc", res[0].OldID)
	assert.Equal(t, "def", res[0].NewID)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, true, gock.IsDone())
}

func TestAPI_CreateColumn_WithOtherColumnOnBoard(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	id := uuid.New()

	gock.New("https://api.yokanban.io").
		MatchHeader("Authorization", "Bearer mock-token").
		Get("/boards/606a34fc20904527edaf3790").
		Reply(200).
		JSON(api.GetBoardResponse{
			Success: true,
			Data: api.BoardDetails{
				Avatars: []api.AvatarDTO{},
				Elements: append([]api.BoardElement{}, api.BoardElement{
					Type: "COLUMN",
					Shape: api.ShapeDTO{
						X:      700,
						Y:      100,
						Width:  400,
						Height: 1000,
					},
					Title: "Column 1",
				}),
			},
		})

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
				  "type": "COLUMN",
				  "title": "test-col",
				  "shape": {
					"x": 1100,
					"y": 100,
					"width": 400,
					"height": 1000
				  },
				  "color": "white",
				  "wipLimit": 0,
				  "isLocked": false,
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
		JSON(api.CreateColumnResponse{
			Success: true,
			Data:    append([]api.CreateColumnResponseDetails{}, api.CreateColumnResponseDetails{OldID: "abc", NewID: "def"}),
		})

	tokenMock := new(accesstoken.Mock)
	tokenMock.On("Get").Return("mock-token")

	a := api.API{AccessToken: tokenMock}
	res := a.CreateColumn("606a34fc20904527edaf3790", "test-col", id)
	assert.Equal(t, "abc", res[0].OldID)
	assert.Equal(t, "def", res[0].NewID)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, true, gock.IsDone())
}

func TestAPI_CreateColumn_WithMultipleColumnsOnBoard(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	id := uuid.New()
	elements := append([]api.BoardElement{}, api.BoardElement{
		Type: "COLUMN",
		Shape: api.ShapeDTO{
			X:      700,
			Y:      100,
			Width:  400,
			Height: 1000,
		},
		Title: "Column 1",
	}, api.BoardElement{
		Type: "COLUMN",
		Shape: api.ShapeDTO{
			X:      1100,
			Y:      100,
			Width:  400,
			Height: 500,
		},
		Title: "Column 2",
	}, api.BoardElement{
		Type: "COLUMN",
		Shape: api.ShapeDTO{
			X:      1,
			Y:      200,
			Width:  400,
			Height: 100,
		},
		Title: "Column",
	})

	gock.New("https://api.yokanban.io").
		MatchHeader("Authorization", "Bearer mock-token").
		Get("/boards/606a34fc20904527edaf3790").
		Reply(200).
		JSON(api.GetBoardResponse{
			Success: true,
			Data: api.BoardDetails{
				Avatars:  []api.AvatarDTO{},
				Elements: elements,
			},
		})

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
				  "type": "COLUMN",
				  "title": "test-col",
				  "shape": {
					"x": 1500,
					"y": 100,
					"width": 400,
					"height": 500
				  },
				  "color": "white",
				  "wipLimit": 0,
				  "isLocked": false,
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
		JSON(api.CreateColumnResponse{
			Success: true,
			Data:    append([]api.CreateColumnResponseDetails{}, api.CreateColumnResponseDetails{OldID: "abc", NewID: "def"}),
		})

	tokenMock := new(accesstoken.Mock)
	tokenMock.On("Get").Return("mock-token")

	a := api.API{AccessToken: tokenMock}
	res := a.CreateColumn("606a34fc20904527edaf3790", "test-col", id)
	assert.Equal(t, "abc", res[0].OldID)
	assert.Equal(t, "def", res[0].NewID)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, true, gock.IsDone())
}
