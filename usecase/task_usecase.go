package usecase

import (
	"clean-architecture-todo-webapp-by-golang/model"
	"clean-architecture-todo-webapp-by-golang/repository"
	"errors"
	"net/http"
)

type ITaskUsecase interface {
	CreateTask(task model.Task) (int, int64, error)
	ReloadTask(userId int) (int, []model.ReloadTask, error)
	DeleteTask(deleteTask model.DeleteTask) (int, error)
}

type taskUsecase struct {
	tr repository.ITaskRepository
}

func NewTaskUsecase(tr repository.ITaskRepository) ITaskUsecase {
	return &taskUsecase{tr}
}

func (tu *taskUsecase) CreateTask(task model.Task) (int, int64, error) {
	if task.UserId == 0 || task.Task == "" {
		return http.StatusBadRequest, 0, errors.New("invalid JSON format")
	}
	if taskId, err := tu.tr.CreateTask(task); err != nil {
		return http.StatusInternalServerError, 0, err
	} else {
		return http.StatusOK, taskId, nil
	}
}

func (tu *taskUsecase) ReloadTask(userId int) (int, []model.ReloadTask, error) {
	var taskData []model.ReloadTask
	if err := tu.tr.ReloadTask(userId, &taskData); err != nil { // ここにdb関連の技術は持ち込んではいけない(rows)
		return http.StatusInternalServerError, taskData, err
	} else {
		return http.StatusOK, taskData, nil
	}
}

func (tu *taskUsecase) DeleteTask(deleteTask model.DeleteTask) (int, error) {
	if deleteTask.UserId == 0 || deleteTask.DeleteId == 0 {
		return http.StatusBadRequest, errors.New("invalid JSON format")
	}
	if err := tu.tr.DeleteTask(deleteTask); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
