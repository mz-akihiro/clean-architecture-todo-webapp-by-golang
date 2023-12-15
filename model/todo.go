package model

type Todo struct {
	UserId int
	Todo   string `json:"task"`
}
