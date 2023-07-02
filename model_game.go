package gomodels

import (
	"encoding/json"
	"fmt"
)

type Outcome int8

const (
	HOME Outcome = iota
	AWAY
	DRAW
)

func (o Outcome) String() string {
	return [...]string{"HOME", "AWAY", "DRAW"}[o]
}

type GameInformation struct {
	GameID  string   `json:"gameId"`
	TeamA   string   `json:"teamA"`
	TeamB   string   `json:"teamB"`
	Date    string   `json:"date"`
	Outcome *Outcome `json:"outcome,omitempty"`
}

func (o *Outcome) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "HOME":
		*o = HOME
	case "AWAY":
		*o = AWAY
	case "DRAW":
		*o = DRAW
	default:
		return fmt.Errorf("invalid outcome value: %s", s)
	}
	return nil
}

func (o Outcome) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.String())
}
