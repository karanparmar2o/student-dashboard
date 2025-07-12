package handler

import (
	"context"

	teacherpb "github.com/karanparmar2o/student-dashboard/api/teacher"
	"github.com/karanparmar2o/student-dashboard/internal/teacher/model"
	"github.com/karanparmar2o/student-dashboard/internal/teacher/service"
)

type TeacherHandler struct {
	service *service.TeacherService
}

func NewTeacherHandler(s *service.TeacherService) *TeacherHandler {
	return &TeacherHandler{service: s}
}

// Register a new teacher
func (h *TeacherHandler) RegisterTeacher(ctx context.Context, req *teacherpb.RegisterTeacherRequest) (*teacherpb.RegisterTeacherResponse, error) {
	teacher := &model.Teacher{
		Name:          req.GetName(),
		Gender:        req.GetGender(),
		Subject:       req.GetSubjects(),
		ClassSections: req.GetClassSections(),
	}

	created, err := h.service.RegisterTeacher(ctx, teacher)
	if err != nil {
		return nil, err
	}

	return &teacherpb.RegisterTeacherResponse{
		Id:      created.ID,
		Success: true,
		Message: "Teacher registered successfully",
	}, nil
}

// Update existing teacher
func (h *TeacherHandler) UpdateTeacher(ctx context.Context, req *teacherpb.UpdateTeacherRequest) (*teacherpb.UpdateTeacherResponse, error) {
	teacher := &model.Teacher{
		Name:          req.GetName(),
		Gender:        req.GetGender(),
		Subject:       req.GetSubjects(),
		ClassSections: req.GetClassSections(),
	}

	_, err := h.service.UpdateTeacher(ctx, req.Id, teacher)
	if err != nil {
		return nil, err
	}

	return &teacherpb.UpdateTeacherResponse{
		Success: true,
		Message: "Teacher registered successfully",
	}, nil
}

// Delete teacher by ID
func (h *TeacherHandler) DeleteTeacher(ctx context.Context, req *teacherpb.DeleteTeacherRequest) (*teacherpb.DeleteTeacherResponse, error) {
	err := h.service.DeleteTeacher(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &teacherpb.DeleteTeacherResponse{Success: true}, nil
}

// List all teachers
func (h *TeacherHandler) ListTeachers(ctx context.Context, req *teacherpb.GetTeacherListRequest) (*teacherpb.GetTeacherListResponse, error) {
	teachers, err := h.service.ListTeachers(ctx)
	if err != nil {
		return nil, err
	}

	resp := &teacherpb.GetTeacherListResponse{}
	for _, t := range teachers {
		resp.Teachers = append(resp.Teachers, &teacherpb.Teacher{
			Id:            t.ID,
			Name:          t.Name,
			Gender:        t.Gender,
			ClassSections: t.ClassSections,
		})
	}
	return resp, nil
}

// Get teacher by name
func (h *TeacherHandler) GetTeacherByName(ctx context.Context, req *teacherpb.GetTeacherByIdRequest) (*teacherpb.GetTeacherByIdResponse, error) {
	t, err := h.service.GetTeacherByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &teacherpb.GetTeacherByIdResponse{
		Teacher: &teacherpb.Teacher{
			Id:            t.ID,
			Name:          t.Name,
			Gender:        t.Gender,
			ClassSections: t.ClassSections,
		},
	}, nil
}
