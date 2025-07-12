package repository

import (
	"sync"

	"github.com/karanparmar2o/student-dashboard/internal/teacher/model"
)

type TeacherRepo struct {
	mu       sync.Mutex
	teachers map[int64]*model.Teacher
	nextID   int64
}

func NewTeacherRepo() *TeacherRepo {
	return &TeacherRepo{
		teachers: make(map[int64]*model.Teacher),
		nextID:   1,
	}
}

func (r *TeacherRepo) Create(teacher *model.Teacher) (*model.Teacher, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	teacher.ID = r.nextID
	r.nextID++
	r.teachers[teacher.ID] = teacher
	return teacher, nil
}

func (r *TeacherRepo) List() ([]*model.Teacher, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	list := make([]*model.Teacher, 0, len(r.teachers))
	for _, t := range r.teachers {
		list = append(list, t)
	}
	return list, nil
}
