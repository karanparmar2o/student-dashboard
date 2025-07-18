package service

import (
	"github.com/karanparmar2o/student-dashboard/internal/student/model"
	repository "github.com/karanparmar2o/student-dashboard/internal/student/repo"
)

type StudentService struct {
	repo repository.StudentRepository
}

func NewStudentService(r repository.StudentRepository) *StudentService {
	return &StudentService{repo: r}
}

func (s *StudentService) RegisterStudent(student model.Student) (int64, error) {
	return s.repo.AddStudent(student)
}

func (s *StudentService) GetStudent(id int64) (*model.Student, error) {
	return s.repo.GetStudentByID(id)
}

func (s *StudentService) GetAllStudents() []model.Student {
	return s.repo.GetAllStudents()
}

func (s *StudentService) UpdateStudent(id int64, student model.Student) error {
	return s.repo.UpdateStudent(id, student)
}

func (s *StudentService) DeleteStudent(id int64) error {
	return s.repo.DeleteStudent(id)
}

func (s *StudentService) GetStudentsForClass(classSection string) []model.Student {
	return s.repo.GetStudentsForClass(classSection)
}

func (s *StudentService) AddMarks(studentID int64, marks []model.SubjectMark) error {
	return s.repo.AddMarks(studentID, marks)
}

func (s *StudentService) UpdateMarks(studentID int64, marks []model.SubjectMark) error {
	return s.repo.UpdateMarks(studentID, marks)
}

func (s *StudentService) GetMarks(studentID int64) ([]model.SubjectMark, error) {
	return s.repo.GetMarks(studentID)
}
