package dto

type TokenDto struct {
	Token  string `json:"token"`
	Expiry int64   `json:"expiry"`
}
