// internal/teacher/repository/teacher_repo.go
package repository

import (
	"context"
	"errors"
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

func (r *TeacherRepo) Create(ctx context.Context, t *model.Teacher) (*model.Teacher, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	t.ID = r.nextID
	r.nextID++
	r.teachers[t.ID] = t
	return t, nil
}

func (r *TeacherRepo) Update(ctx context.Context, id int64, t *model.Teacher) (*model.Teacher, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.teachers[id]
	if !exists {
		return nil, errors.New("teacher not found")
	}
	t.ID = id
	r.teachers[id] = t
	return t, nil
}

func (r *TeacherRepo) Delete(ctx context.Context, id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.teachers[id]; !ok {
		return errors.New("teacher not found")
	}
	delete(r.teachers, id)
	return nil
}

func (r *TeacherRepo) List(ctx context.Context) ([]*model.Teacher, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var list []*model.Teacher
	for _, t := range r.teachers {
		list = append(list, t)
	}
	return list, nil
}

func (r *TeacherRepo) GetByID(ctx context.Context, id int64) (*model.Teacher, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	t, ok := r.teachers[id]
	if !ok {
		return nil, errors.New("teacher not found")
	}
	return t, nil
}
