package models

var SendUser struct {
	Username  string
	Password  string
	Hierarchy string `json:"hierarchy"`
}
