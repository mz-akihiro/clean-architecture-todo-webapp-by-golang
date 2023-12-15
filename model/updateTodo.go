package model

type UpdateTodo struct {
	TodoId int `json:"TodoId"`
	UserId int
	Todo   string `json:"Todo"`
}
