package models

type AuthDTO struct {
	UserID string `json:"userId"`
	Token  string `json:"token"`
}
