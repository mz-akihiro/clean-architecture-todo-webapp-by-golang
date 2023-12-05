package repository

import (
	"clean-architecture-todo-webapp-by-golang/model"
	"database/sql"
)

type ITaskRepository interface {
	CreateTodo(todo model.Todo) (int64, error)
	ReadTodo(userId int, todoData *[]model.ReadTodo) error
	DeleteTodo(deleteTodo model.DeleteTodo) error
}

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) ITaskRepository {
	return &taskRepository{db}
}

func (tr *taskRepository) CreateTodo(todo model.Todo) (int64, error) {
	if result, err := tr.db.Exec("INSERT INTO todos (userId, todo) VALUES (?, ?)", todo.UserId, todo.Todo); err != nil {
		return 0, err
	} else {
		taskId, _ := result.LastInsertId()
		return taskId, nil
	}
}

func (tr *taskRepository) ReadTodo(userId int, todoData *[]model.ReadTodo) error {
	if rows, err := tr.db.Query("SELECT id, todo FROM todos WHERE userId = ? AND deleted = false", userId); err != nil {
		return err
	} else {
		defer rows.Close()
		for rows.Next() {
			var td model.ReadTodo
			if err := rows.Scan(&td.TaskId, &td.Todo); err != nil {
				return err
			}
			*todoData = append(*todoData, td)
		}
		return nil
	}
}

func (tr *taskRepository) DeleteTodo(deleteTodo model.DeleteTodo) error {
	if _, err := tr.db.Exec("UPDATE todos set deleted = true where id = ? and userId = ?;", deleteTodo.DeleteId, deleteTodo.UserId); err != nil {
		return err
	}
	return nil
}
