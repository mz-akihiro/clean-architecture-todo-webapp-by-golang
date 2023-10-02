package controller

import (
	"clean-architecture-todo-webapp-by-golang/model"
	"clean-architecture-todo-webapp-by-golang/usecase"
	"encoding/json"
	"fmt"
	"net/http"
)

type ITaskController interface {
	CreateTask(w http.ResponseWriter, r *http.Request)
	ReloadTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
}

type taskController struct {
	tu usecase.ITaskUsecase
}

func NewTaskController(tu usecase.ITaskUsecase) ITaskController {
	return &taskController{tu}
}

func (tc *taskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	task := model.Task{}
	//userId, ok := r.Context().Value("testId").(int)
	userId, ok := r.Context().Value(model.ContextKey{UserId: "user_id"}).(int)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Context value is incorrect")
		return
	}
	task.UserId = userId
	// ここでtaskがからでも検知できていない
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid JSON format")
		return
	}

	if statusCode, taskId, err := tc.tu.CreateTask(task); err != nil {
		w.WriteHeader(statusCode)
		fmt.Fprintln(w, err)
		return
	} else {
		jsonData := model.TaskIdReturn{}
		jsonData.Id = taskId
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(jsonData); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Faild json encode")
			return
		}
	}
}

func (tc *taskController) ReloadTask(w http.ResponseWriter, r *http.Request) {
	//userId, ok := r.Context().Value("testId").(int)
	userId, ok := r.Context().Value(model.ContextKey{UserId: "user_id"}).(int)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Context value is incorrect")
		return
	}
	if _, taskData, err := tc.tu.ReloadTask(userId); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Context value is incorrect")
		return
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(taskData); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Faild json encode")
			return
		}
	}
}

func (tc *taskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	var deleteTask model.DeleteTask
	userId, ok := r.Context().Value(model.ContextKey{UserId: "user_id"}).(int)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Context value is incorrect")
		return
	}
	deleteTask.UserId = userId
	if err := json.NewDecoder(r.Body).Decode(&deleteTask); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid JSON format")
		return
	}
	fmt.Println(deleteTask)
	if statusCode, err := tc.tu.DeleteTask(deleteTask); err != nil {
		w.WriteHeader(statusCode)
		fmt.Fprintln(w, err)
		return
	}
}
