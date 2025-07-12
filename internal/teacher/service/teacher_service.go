// internal/teacher/service/teacher_service.go
package service

import (
	"context"

	"github.com/karanparmar2o/student-dashboard/internal/teacher/model"
	"github.com/karanparmar2o/student-dashboard/internal/teacher/repository"
)

type TeacherService struct {
	repo *repository.TeacherRepo
}

func NewTeacherService(r *repository.TeacherRepo) *TeacherService {
	return &TeacherService{repo: r}
}

func (s *TeacherService) RegisterTeacher(ctx context.Context, t *model.Teacher) (*model.Teacher, error) {
	return s.repo.Create(ctx, t)
}

func (s *TeacherService) UpdateTeacher(ctx context.Context, id int64, t *model.Teacher) (*model.Teacher, error) {
	return s.repo.Update(ctx, id, t)
}

func (s *TeacherService) DeleteTeacher(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *TeacherService) ListTeachers(ctx context.Context) ([]*model.Teacher, error) {
	return s.repo.List(ctx)
}

func (s *TeacherService) GetTeacherByID(ctx context.Context, id int64) (*model.Teacher, error) {
	return s.repo.GetByID(ctx, id)
}
