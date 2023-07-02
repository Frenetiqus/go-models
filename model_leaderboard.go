package gomodels

type Leaderboard struct {
	UserID      string `json:"userId"`
	Username    string `json:"username"`
	Season      int    `json:"season"`
	TotalPoints int    `json:"totalPoints"`
	Rank        int    `json:"rank"`
}
