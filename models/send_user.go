package models

type SendUser struct {
	Username  string
	Password  string
	Hierarchy string `json:"hierarchy"`
}
