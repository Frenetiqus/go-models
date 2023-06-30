package model

type CreatedUserConfiguration struct {
	UserConfiguration UserConfiguration `json:"userConfiguration"`
	AuthToken         *string           `json:"authToken"`
}
