package task

import (
	"gorm.io/gorm"
)

type DB interface {
	SaveTask(task *Task) (*Task, error)
	SoftDeleteTask(task *Task) (*Task, error)
	UpdateTask(task *Task) (*Task, error)
}

type GormRepo struct {
	db *gorm.DB
}

func NewGormRepo(db *gorm.DB) *GormRepo {
	return &GormRepo{db: db}
}

func (r *GormRepo) CreateTask(task *Task) (*Task, error) {
	if err := r.db.Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (r *GormRepo) DeleteTask(task *Task) (*Task, error) {
	if err := r.db.Delete(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (r *GormRepo) UpdateTask(task *Task) (*Task, error) {
	if _, err := r.db.; err != nil {
		return nil, err
	}
	return nil, nil
}
