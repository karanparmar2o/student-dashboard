package service

import (
	"github.com/karanparmar2o/student-dashboard/internal/student/model"
	repository "github.com/karanparmar2o/student-dashboard/internal/student/repo"
)

type StudentService struct {
	repo *repository.StudentRepo
}

func NewStudentService(repo *repository.StudentRepo) *StudentService {
	return &StudentService{repo: repo}
}

func (s *StudentService) RegisterStudent(student *model.Student) (*model.Student, error) {
	return s.repo.Create(student)
}

func (s *StudentService) GetStudents() ([]*model.Student, error) {
	return s.repo.List()
}
