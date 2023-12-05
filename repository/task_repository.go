package repository

import (
	"clean-architecture-todo-webapp-by-golang/model"
	"database/sql"
)

type ITaskRepository interface {
	CreateTask(task model.Task) (int64, error)
	ReloadTask(userId int, taskData *[]model.ReloadTask) error
	DeleteTask(deleteTask model.DeleteTask) error
}

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) ITaskRepository {
	return &taskRepository{db}
}

func (tr *taskRepository) CreateTask(task model.Task) (int64, error) {
	if result, err := tr.db.Exec("INSERT INTO todos (userId, todo) VALUES (?, ?)", task.UserId, task.Task); err != nil {
		return 0, err
	} else {
		taskId, _ := result.LastInsertId()
		return taskId, nil
	}
}

func (tr *taskRepository) ReloadTask(userId int, taskData *[]model.ReloadTask) error {
	if rows, err := tr.db.Query("SELECT id, todo FROM todos WHERE userId = ? AND deleted = false", userId); err != nil {
		return err
	} else {
		defer rows.Close()
		for rows.Next() {
			var td model.ReloadTask
			if err := rows.Scan(&td.TaskId, &td.Task); err != nil {
				return err
			}
			*taskData = append(*taskData, td)
		}
		return nil
	}
}

func (tr *taskRepository) DeleteTask(deleteTask model.DeleteTask) error {
	if _, err := tr.db.Exec("UPDATE todos set deleted = true where id = ? and userId = ?;", deleteTask.DeleteId, deleteTask.UserId); err != nil {
		return err
	}
	return nil
}
