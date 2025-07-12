package handler

import (
	"context"
	"strconv"

	studentpb "github.com/karanparmar2o/student-dashboard/api/student"
	"github.com/karanparmar2o/student-dashboard/internal/student/model"
	"github.com/karanparmar2o/student-dashboard/internal/student/service"
)

type StudentHandler struct {
	studentpb.UnimplementedStudentServiceServer
	svc *service.StudentService
}

func NewStudentHandler(svc *service.StudentService) *StudentHandler {
	return &StudentHandler{svc: svc}
}

func (h *StudentHandler) RegisterStudent(ctx context.Context, req *studentpb.RegisterStudentRequest) (*studentpb.RegisterStudentResponse, error) {
	student := &model.Student{
		Name:    req.GetName(),
		Age:     req.GetAge(),
		RollNo:  req.GetRollNumber(),
		Class:   req.GetClass(),
		Section: req.GetSection(),
	}

	_, err := h.svc.RegisterStudent(student)
	if err != nil {
		return &studentpb.RegisterStudentResponse{
			Success: false,
			Message: "Failed to register student",
		}, err
	}

	return &studentpb.RegisterStudentResponse{
		Success: true,
		Message: "Student registered successfully",
	}, nil
}

func (h *StudentHandler) GetStudentList(ctx context.Context, req *studentpb.GetStudentListRequest) (*studentpb.GetStudentListResponse, error) {
	students, err := h.svc.GetStudents()
	if err != nil {
		return nil, err
	}
	resp := &studentpb.GetStudentListResponse{}
	for _, s := range students {
		classInt, err := strconv.Atoi(s.Class)
		if err != nil {

		}
		resp.Students = append(resp.Students, &studentpb.Student{

			Name:    s.Name,
			Age:     s.Age,
			RollNo:  s.RollNo,
			Section: s.Section,
			Class:   int32(classInt),
		})
	}

	return resp, nil
}
