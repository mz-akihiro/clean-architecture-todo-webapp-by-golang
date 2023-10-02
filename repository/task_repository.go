package repository

import (
	"clean-architecture-todo-webapp-by-golang/model"
	"database/sql"
)

type ITaskRepository interface {
	CreateTask(task model.Task) (int64, error)
	ReloadTask(userId int) (*sql.Rows, error)
	DeleteTask(deleteTask model.DeleteTask) error
}

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) ITaskRepository {
	return &taskRepository{db}
}

func (tr *taskRepository) CreateTask(task model.Task) (int64, error) {
	if result, err := tr.db.Exec("INSERT INTO task_data (userId, memo) VALUES (?, ?)", task.UserId, task.Task); err != nil {
		return 0, err
	} else {
		taskId, _ := result.LastInsertId()
		return taskId, nil
	}
}

func (tr *taskRepository) ReloadTask(userId int) (*sql.Rows, error) {
	if rows, err := tr.db.Query("SELECT id,memo FROM task_data WHERE userId = ?", userId); err != nil {
		return rows, err
	} else {
		return rows, nil
	}
}

func (tr *taskRepository) DeleteTask(deleteTask model.DeleteTask) error {
	if _, err := tr.db.Exec("DELETE FROM task_data WHERE id = ? AND userId = ?", deleteTask.DeleteId, deleteTask.UserId); err != nil {
		return err
	}
	return nil
}
