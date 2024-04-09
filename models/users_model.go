package models

type Users struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Hierarchy string `json:"hierarchy"`
}
