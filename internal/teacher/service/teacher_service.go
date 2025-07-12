package service

import (
	"context"

	"github.com/karanparmar2o/student-dashboard/internal/teacher/model"
	"github.com/karanparmar2o/student-dashboard/internal/teacher/repository"
)

type TeacherService struct {
	repo *repository.TeacherRepo
}

func NewTeacherService(repo *repository.TeacherRepo) *TeacherService {
	return &TeacherService{repo: repo}
}

func (s *TeacherService) RegisterTeacher(ctx context.Context, teacher *model.Teacher) (*model.Teacher, error) {
	return s.repo.Create(teacher)
}

func (s *TeacherService) ListTeachers(ctx context.Context) ([]*model.Teacher, error) {
	return s.repo.List()
}
