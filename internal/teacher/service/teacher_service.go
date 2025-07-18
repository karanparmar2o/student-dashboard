package service

import "github.com/karanparmar2o/student-dashboard/internal/teacher/repository"

type TeacherService struct {
	repo repository.TeacherRepository
}

func NewTeacherService(r repository.TeacherRepository) *TeacherService {
	return &TeacherService{repo: r}
}

func (s *TeacherService) RegisterTeacher(t repository.Teacher) (int64, error) {
	return s.repo.AddTeacher(t)
}

func (s *TeacherService) GetTeacher(id int64) (*repository.Teacher, error) {
	return s.repo.GetTeacherByID(id)
}

func (s *TeacherService) GetAllTeachers() []repository.Teacher {
	return s.repo.GetAllTeachers()
}

func (s *TeacherService) UpdateTeacher(t repository.Teacher) error {
	return s.repo.UpdateTeacher(t)
}

func (s *TeacherService) DeleteTeacher(id int64) error {
	return s.repo.DeleteTeacher(id)
}

func (s *TeacherService) GetClassesForTeacher(teacherID int64) ([]string, error) {
	return s.repo.GetClassesForTeacher(teacherID)
}
