package repository

import (
	"errors"
	"fmt"
	"sync"

	"github.com/karanparmar2o/student-dashboard/internal/student/model"
)

type StudentRepository interface {
	AddStudent(s model.Student) (int64, error)
	GetStudentByID(id int64) (*model.Student, error)
	GetAllStudents() []model.Student
	UpdateStudent(id int64, s model.Student) error
	DeleteStudent(id int64) error
	GetStudentsForClass(classSection string) []model.Student
	AddMarks(studentID int64, marks []model.SubjectMark) error
	UpdateMarks(studentID int64, marks []model.SubjectMark) error
	GetMarks(studentID int64) ([]model.SubjectMark, error)
}

type memoryStudentRepo struct {
	mu       sync.RWMutex
	students map[int64]model.Student
	nextID   int64
}

func NewMemoryStudentRepo() StudentRepository {
	return &memoryStudentRepo{
		students: make(map[int64]model.Student),
		nextID:   1,
	}
}

func (r *memoryStudentRepo) AddStudent(s model.Student) (int64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	s.ID = r.nextID
	r.students[s.ID] = s
	//fmt.Printf("Adding student:", r.students[s.ID])
	r.nextID++
	return s.ID, nil
}

func (r *memoryStudentRepo) GetStudentByID(id int64) (*model.Student, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	student, ok := r.students[id]
	if !ok {
		return nil, errors.New("student not found")
	}
	return &student, nil
}

func (r *memoryStudentRepo) GetAllStudents() []model.Student {
	r.mu.RLock()
	defer r.mu.RUnlock()
	students := make([]model.Student, 0, len(r.students))
	for _, s := range r.students {
		students = append(students, s)
	}
	return students
}

func (r *memoryStudentRepo) UpdateStudent(id int64, s model.Student) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.students[id]; !exists {
		return errors.New("student not found")
	}
	s.ID = id // ensure ID consistency
	r.students[id] = s
	return nil
}

func (r *memoryStudentRepo) DeleteStudent(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.students[id]; !exists {
		return errors.New("student not found")
	}
	delete(r.students, id)
	return nil
}

func (r *memoryStudentRepo) GetStudentsForClass(classSection string) []model.Student {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var students []model.Student
	for _, s := range r.students {
		//fmt.Printf("Checking student: %s_%s\n", s.Class, s.Section)
		if s.Class+"_"+s.Section == classSection {
			students = append(students, s)
		} else {
			fmt.Printf("no students found for class section: %s\n", classSection)
		}
	}
	return students
}

func (r *memoryStudentRepo) AddMarks(studentID int64, marks []model.SubjectMark) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	student, exists := r.students[studentID]
	if !exists {
		return errors.New("student not found")
	}
	student.Marks = append(student.Marks, marks...) // add new marks
	r.students[studentID] = student
	return nil
}

func (r *memoryStudentRepo) UpdateMarks(studentID int64, marks []model.SubjectMark) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	student, exists := r.students[studentID]
	if !exists {
		return errors.New("student not found")
	}
	student.Marks = marks // replace marks
	r.students[studentID] = student
	return nil
}

func (r *memoryStudentRepo) GetMarks(studentID int64) ([]model.SubjectMark, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	student, exists := r.students[studentID]
	if !exists {
		return nil, errors.New("student not found")
	}
	return student.Marks, nil
}
