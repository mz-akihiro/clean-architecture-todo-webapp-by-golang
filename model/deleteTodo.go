package model

type DeleteTodo struct {
	UserId   int
	DeleteId int `json:"TaskId"`
}
