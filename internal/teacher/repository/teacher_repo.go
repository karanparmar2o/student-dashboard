package repository

import (
	"errors"
	"sync"
)

type Teacher struct {
	ID            int64
	Name          string
	Gender        string
	Subjects      []string
	ClassSections []string
}

type TeacherRepository interface {
	AddTeacher(t Teacher) (int64, error)
	GetTeacherByID(id int64) (*Teacher, error)
	GetAllTeachers() []Teacher
	UpdateTeacher(t Teacher) error
	DeleteTeacher(id int64) error
	GetClassesForTeacher(teacherID int64) ([]string, error)
}

type memoryTeacherRepo struct {
	mu       sync.RWMutex
	teachers map[int64]Teacher
	nextID   int64
}

func NewMemoryTeacherRepo() TeacherRepository {
	return &memoryTeacherRepo{
		teachers: make(map[int64]Teacher),
		nextID:   1,
	}
}

func (r *memoryTeacherRepo) AddTeacher(t Teacher) (int64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	t.ID = r.nextID
	r.teachers[t.ID] = t
	r.nextID++
	return t.ID, nil
}

func (r *memoryTeacherRepo) GetTeacherByID(id int64) (*Teacher, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	teacher, exists := r.teachers[id]
	if !exists {
		return nil, errors.New("teacher not found")
	}
	return &teacher, nil
}

func (r *memoryTeacherRepo) GetAllTeachers() []Teacher {
	r.mu.RLock()
	defer r.mu.RUnlock()
	teachers := make([]Teacher, 0, len(r.teachers))
	for _, t := range r.teachers {
		teachers = append(teachers, t)
	}
	return teachers
}

func (r *memoryTeacherRepo) UpdateTeacher(t Teacher) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.teachers[t.ID]; !exists {
		return errors.New("teacher not found")
	}
	r.teachers[t.ID] = t
	return nil
}

func (r *memoryTeacherRepo) DeleteTeacher(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.teachers[id]; !exists {
		return errors.New("teacher not found")
	}
	delete(r.teachers, id)
	return nil
}

func (r *memoryTeacherRepo) GetClassesForTeacher(teacherID int64) ([]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	teacher, exists := r.teachers[teacherID]
	if !exists {
		return nil, errors.New("teacher not found")
	}
	return teacher.ClassSections, nil
}
