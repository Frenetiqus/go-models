package gomodels

// http headers: Location: /users/<user_id>, Authorization: Bearer 786598273423iuy4g
type UserConfiguration struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Email    *string `json:"email,omitempty"`
}
