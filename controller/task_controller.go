package controller

import (
	"clean-architecture-todo-webapp-by-golang/model"
	"clean-architecture-todo-webapp-by-golang/usecase"
	"encoding/json"
	"fmt"
	"net/http"
)

type ITaskController interface {
	CreateTodo(w http.ResponseWriter, r *http.Request)
	ReadTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
}

type taskController struct {
	tu usecase.ITaskUsecase
}

func NewTaskController(tu usecase.ITaskUsecase) ITaskController {
	return &taskController{tu}
}

func (tc *taskController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	todo := model.Todo{}
	//userId, ok := r.Context().Value("testId").(int)
	userId, ok := r.Context().Value(model.ContextKey{UserId: "user_id"}).(int)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Context value is incorrect")
		return
	}
	todo.UserId = userId
	// ここでtaskがからでも検知できていない
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid JSON format")
		return
	}

	if statusCode, taskId, err := tc.tu.CreateTodo(todo); err != nil {
		w.WriteHeader(statusCode)
		fmt.Fprintln(w, err)
		return
	} else {
		jsonData := model.TodoIdReturn{}
		jsonData.Id = taskId
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(jsonData); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Faild json encode")
			return
		}
	}
}

func (tc *taskController) ReadTodo(w http.ResponseWriter, r *http.Request) {
	//userId, ok := r.Context().Value("testId").(int)
	userId, ok := r.Context().Value(model.ContextKey{UserId: "user_id"}).(int)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Context value is incorrect")
		return
	}
	if _, taskData, err := tc.tu.ReadTodo(userId); err != nil {
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

func (tc *taskController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	var deleteTask model.DeleteTodo
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
	if statusCode, err := tc.tu.DeleteTodo(deleteTask); err != nil {
		w.WriteHeader(statusCode)
		fmt.Fprintln(w, err)
		return
	}
}
