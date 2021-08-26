package models

type ErrRes struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
	Source  string `json:"source"`
}
