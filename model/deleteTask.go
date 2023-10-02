package model

type DeleteTask struct {
	UserId   int
	DeleteId int `json:"TaskId"`
}
