package gomodels

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameInformationUnmarshallSucces(t *testing.T) {
	assert := assert.New(t)
	recvJSON := `
    {
        "gameId":"123",
        "teamA":"Barca",
        "teamB":"Arsenal",
        "date":"2024-04-20T14:20:00+02:00"
    }
    `
	expectedStruct := GameInformation{GameID: "123", TeamA: "Barca", TeamB: "Arsenal", Date: "2024-04-20T14:20:00+02:00"}

	var unmarshalled GameInformation
	err := json.Unmarshal([]byte(recvJSON), &unmarshalled)
	assert.Nil(err)
	assert.Equal(expectedStruct, unmarshalled)

	recvJSON = `
    [
        {
            "gameId":"123",
            "teamA":"Barca",
            "teamB":"Arsenal",
            "date":"2024-04-20T14:20:00+02:00"
        },
        {
            "gameId":"345",
            "teamA":"Barca",
            "teamB":"City",
            "date":"2022-04-20T14:20:00+02:00",
            "outcome":"AWAY"
        }
    ]
    `
	outcome := AWAY
	expectedStructs := []GameInformation{
		{
			GameID: "123",
			TeamA:  "Barca",
			TeamB:  "Arsenal",
			Date:   "2024-04-20T14:20:00+02:00",
		},
		{
			GameID:  "345",
			TeamA:   "Barca",
			TeamB:   "City",
			Date:    "2022-04-20T14:20:00+02:00",
			Outcome: &outcome,
		},
	}

	var unmarshalleds []GameInformation
	err = json.Unmarshal([]byte(recvJSON), &unmarshalleds)
	assert.Nil(err)
	assert.Equal(expectedStructs, unmarshalleds)
}

func TestGameInformationUnmarshallFail(t *testing.T) {
	assert := assert.New(t)
	recvJSON := `
    {
        "gameId":"123",
        "teamA":"Barca",
        "teamB":"Arsenal",
        "date":"2024-04-20T14:20:00+02:00",
		"outcome":"NOT_HOME_OR_AWAY_OR_DRAW"
    }
    `

	var unmarshalled GameInformation
	err := json.Unmarshal([]byte(recvJSON), &unmarshalled)
	assert.NotNil(err)
}
