package task_manager

import "time"

type Task struct {
	title       string
	description string
	createdAt   time.Time
	expiresAt   time.Time
	softDeleted bool
}

type DB interface {
	SaveTask(task *Task) (*Task, error)
	SoftDeleteTask(task *Task) (*Task, error)
	UpdateTask(task *Task) (*Task, error)
}

type TaskManager struct {
	db DB
}

func NewTaskManager(db *DB) *TaskManager {
	return &TaskManager{db: db}
}

func (t *TaskManager) CreateTask(task *Task) (*Task, error) {
	if _, err := t.db.SaveTask(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (t *TaskManager) DeleteTask(task *Task) (*Task, error) {
	if _, err := t.db.SoftDeleteTask(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (t *TaskManager) UpdateTask(task *Task) (*Task, error) {
	if _, err := t.db.SaveTask(task); err != nil {
		return nil, err
	}
	return nil, nil
}
