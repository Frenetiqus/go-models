package gomodels

type BetConfiguration struct {
	UserID           string           `json:"userId"`
	GameID           string           `json:"gameId"`
	BetAmount        float32          `json:"betAmount"`
	PredictedOutcome PredictedOutcome `json:"predictedOutcome"`
}
