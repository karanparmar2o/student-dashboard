package repository

import (
	"sync"

	"github.com/karanparmar2o/student-dashboard/internal/student/model"
)

type StudentRepo struct {
	mu       sync.Mutex
	students map[int64]*model.Student
	nextID   int64
}

func NewStudentRepo() *StudentRepo {
	return &StudentRepo{
		students: make(map[int64]*model.Student),
		nextID:   1,
	}
}

func (r *StudentRepo) Create(student *model.Student) (*model.Student, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	student.ID = r.nextID
	r.nextID++
	student.ClassSection = student.Class + student.Section
	r.students[student.ID] = student
	return student, nil
}

func (r *StudentRepo) List() ([]*model.Student, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	list := make([]*model.Student, 0, len(r.students))
	for _, s := range r.students {
		list = append(list, s)
	}
	return list, nil
}
