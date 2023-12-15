package usecase

import (
	"clean-architecture-todo-webapp-by-golang/model"
	"clean-architecture-todo-webapp-by-golang/repository"
	"errors"
	"net/http"
)

type ITaskUsecase interface {
	CreateTodo(todo model.Todo) (int, int64, error)
	ReadTodo(userId int) (int, []model.ReadTodo, error)
	UpdateTodo(updateTodo model.UpdateTodo) (int, error)
	DeleteTodo(deleteTask model.DeleteTodo) (int, error)
}

type taskUsecase struct {
	tr repository.ITaskRepository
}

func NewTaskUsecase(tr repository.ITaskRepository) ITaskUsecase {
	return &taskUsecase{tr}
}

func (tu *taskUsecase) CreateTodo(todo model.Todo) (int, int64, error) {
	if todo.UserId == 0 || todo.Todo == "" {
		return http.StatusBadRequest, 0, errors.New("invalid JSON format")
	}
	if taskId, err := tu.tr.CreateTodo(todo); err != nil {
		return http.StatusInternalServerError, 0, err
	} else {
		return http.StatusOK, taskId, nil
	}
}

func (tu *taskUsecase) ReadTodo(userId int) (int, []model.ReadTodo, error) {
	var taskData []model.ReadTodo
	if err := tu.tr.ReadTodo(userId, &taskData); err != nil { // ここにdb関連の技術は持ち込んではいけない(rows)
		return http.StatusInternalServerError, taskData, err
	} else {
		return http.StatusOK, taskData, nil
	}
}

func (tu *taskUsecase) UpdateTodo(updateTodo model.UpdateTodo) (int, error) {
	if updateTodo.UserId == 0 || updateTodo.TodoId == 0 || updateTodo.Todo == "" {
		return http.StatusBadRequest, errors.New("invalid JSON format")
	}
	if err := tu.tr.UpdateTodo(updateTodo); err != nil {
		return http.StatusInternalServerError, err
	} else {
		return http.StatusOK, nil
	}
}

func (tu *taskUsecase) DeleteTodo(deleteTask model.DeleteTodo) (int, error) {
	if deleteTask.UserId == 0 || deleteTask.DeleteId == 0 {
		return http.StatusBadRequest, errors.New("invalid JSON format")
	}
	if err := tu.tr.DeleteTodo(deleteTask); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
