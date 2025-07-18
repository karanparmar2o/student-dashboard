package handler

import (
	"context"

	teacherpb "github.com/karanparmar2o/student-dashboard/api/teacher"
	"github.com/karanparmar2o/student-dashboard/internal/teacher/repository"
	"github.com/karanparmar2o/student-dashboard/internal/teacher/service"
)

type TeacherHandler struct {
	teacherpb.UnimplementedTeacherServiceServer
	svc *service.TeacherService
}

func NewTeacherHandler(s *service.TeacherService) *TeacherHandler {
	return &TeacherHandler{svc: s}
}

func (h *TeacherHandler) RegisterTeacher(ctx context.Context, req *teacherpb.RegisterTeacherRequest) (*teacherpb.RegisterTeacherResponse, error) {
	t := repository.Teacher{
		Name:          req.Name,
		Gender:        req.Gender,
		Subjects:      req.Subject,
		ClassSections: req.ClassSections,
	}
	id, err := h.svc.RegisterTeacher(t)
	if err != nil {
		return &teacherpb.RegisterTeacherResponse{Success: false, Message: err.Error()}, nil
	}
	return &teacherpb.RegisterTeacherResponse{Success: true, Id: id, Message: "Teacher registered"}, nil
}

func (h *TeacherHandler) GetTeacherList(ctx context.Context, req *teacherpb.GetTeacherListRequest) (*teacherpb.GetTeacherListResponse, error) {
	teachers := h.svc.GetAllTeachers()
	var pbTeachers []*teacherpb.Teacher
	for _, t := range teachers {
		pbTeachers = append(pbTeachers, &teacherpb.Teacher{
			Id:            t.ID,
			Name:          t.Name,
			Gender:        t.Gender,
			Subjects:      t.Subjects,
			ClassSections: t.ClassSections,
		})
	}
	return &teacherpb.GetTeacherListResponse{Teachers: pbTeachers}, nil
}

func (h *TeacherHandler) GetTeacherById(ctx context.Context, req *teacherpb.GetTeacherByIdRequest) (*teacherpb.GetTeacherByIdResponse, error) {
	t, err := h.svc.GetTeacher(req.Id)
	if err != nil {
		return nil, err
	}
	return &teacherpb.GetTeacherByIdResponse{
		Teacher: &teacherpb.Teacher{
			Id:            t.ID,
			Name:          t.Name,
			Gender:        t.Gender,
			Subjects:      t.Subjects,
			ClassSections: t.ClassSections,
		},
	}, nil
}

func (h *TeacherHandler) UpdateTeacher(ctx context.Context, req *teacherpb.UpdateTeacherRequest) (*teacherpb.UpdateTeacherResponse, error) {
	t := repository.Teacher{
		ID:            req.Id,
		Name:          req.Name,
		Gender:        req.Gender,
		Subjects:      req.Subjects,
		ClassSections: req.ClassSections,
	}
	err := h.svc.UpdateTeacher(t)
	if err != nil {
		return &teacherpb.UpdateTeacherResponse{Success: false, Message: err.Error()}, nil
	}
	return &teacherpb.UpdateTeacherResponse{Success: true, Message: "Teacher updated"}, nil
}

func (h *TeacherHandler) DeleteTeacher(ctx context.Context, req *teacherpb.DeleteTeacherRequest) (*teacherpb.DeleteTeacherResponse, error) {
	err := h.svc.DeleteTeacher(req.Id)
	if err != nil {
		return &teacherpb.DeleteTeacherResponse{Success: false, Message: err.Error()}, nil
	}
	return &teacherpb.DeleteTeacherResponse{Success: true, Message: "Teacher deleted"}, nil
}

func (h *TeacherHandler) GetClassesForTeacher(ctx context.Context, req *teacherpb.GetClassesForTeacherRequest) (*teacherpb.GetClassesForTeacherResponse, error) {
	classes, err := h.svc.GetClassesForTeacher(req.TeacherId)
	if err != nil {
		return nil, err
	}
	return &teacherpb.GetClassesForTeacherResponse{ClassSections: classes}, nil
}
