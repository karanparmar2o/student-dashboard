package handler

import (
	"context"

	pb "github.com/karanparmar2o/student-dashboard/api/student"
	"github.com/karanparmar2o/student-dashboard/internal/student/model"
	"github.com/karanparmar2o/student-dashboard/internal/student/service"
)

type StudentHandler struct {
	pb.UnimplementedStudentServiceServer
	service *service.StudentService
}

func NewStudentHandler(s *service.StudentService) *StudentHandler {
	return &StudentHandler{service: s}
}

func (h *StudentHandler) RegisterStudent(ctx context.Context, req *pb.RegisterStudentRequest) (*pb.RegisterStudentResponse, error) {
	id, err := h.service.RegisterStudent(model.Student{
		Name:        req.Name,
		Age:         req.Age,
		Class:       req.Class,
		Section:     req.Section,
		ParentName:  req.ParentName,
		ParentPhone: req.ParentPhone,
		Gender:      req.Gender,
		Address:     req.Address,
	})
	if err != nil {
		return &pb.RegisterStudentResponse{Success: false, Message: err.Error()}, nil
	}
	return &pb.RegisterStudentResponse{Success: true, Message: "student registered", Id: id}, nil
}

func (h *StudentHandler) GetStudentList(ctx context.Context, req *pb.GetStudentListRequest) (*pb.GetStudentListResponse, error) {
	students := h.service.GetAllStudents()
	var pbStudents []*pb.Student
	for _, s := range students {
		pbStudents = append(pbStudents, &pb.Student{
			Id:      s.ID,
			Name:    s.Name,
			Gender:  s.Gender,
			Profile: "",
		})
	}
	return &pb.GetStudentListResponse{Students: pbStudents}, nil
}

func (h *StudentHandler) GetStudentById(ctx context.Context, req *pb.GetStudentByIdRequest) (*pb.GetStudentByIdResponse, error) {
	s, err := h.service.GetStudent(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetStudentByIdResponse{
		Student: &pb.Student{
			Id:      s.ID,
			Name:    s.Name,
			Gender:  s.Gender,
			Profile: "",
		},
	}, nil
}

func (h *StudentHandler) UpdateStudent(ctx context.Context, req *pb.UpdateStudentRequest) (*pb.UpdateStudentResponse, error) {
	err := h.service.UpdateStudent(req.Id, model.Student{
		Name:        req.Name,
		Age:         req.Age,
		Class:       req.Class,
		Section:     req.Section,
		ParentName:  req.ParentName,
		ParentPhone: req.ParentPhone,
		Gender:      req.Gender,
		Address:     req.Address,
	})
	if err != nil {
		return &pb.UpdateStudentResponse{Success: false, Message: err.Error()}, nil
	}
	return &pb.UpdateStudentResponse{Success: true, Message: "student updated"}, nil
}

func (h *StudentHandler) DeleteStudent(ctx context.Context, req *pb.DeleteStudentRequest) (*pb.DeleteStudentResponse, error) {
	err := h.service.DeleteStudent(req.Id)
	if err != nil {
		return &pb.DeleteStudentResponse{Success: false, Message: err.Error()}, nil
	}
	return &pb.DeleteStudentResponse{Success: true, Message: "student deleted"}, nil
}

func (h *StudentHandler) GetStudentsForClass(ctx context.Context, req *pb.GetStudentsForClassRequest) (*pb.GetStudentsForClassResponse, error) {

	students := h.service.GetStudentsForClass(req.ClassSection)
	var pbStudents []*pb.Student
	for _, s := range students {
		pbStudents = append(pbStudents, &pb.Student{
			Id:      s.ID,
			Name:    s.Name,
			Gender:  s.Gender,
			Profile: "",
		})
	}
	return &pb.GetStudentsForClassResponse{Students: pbStudents}, nil
}

func (h *StudentHandler) AddMarks(ctx context.Context, req *pb.AddMarksRequest) (*pb.AddMarksResponse, error) {
	var marks []model.SubjectMark
	for _, m := range req.Marks {
		marks = append(marks, model.SubjectMark{
			SubjectName:      m.SubjectName,
			Mark:             m.Mark,
			GivenByTeacherID: m.GivenByTeacherId,
		})
	}
	err := h.service.AddMarks(req.StudentId, marks)
	if err != nil {
		return &pb.AddMarksResponse{Success: false, Message: err.Error()}, nil
	}
	return &pb.AddMarksResponse{Success: true, Message: "marks added"}, nil
}

func (h *StudentHandler) UpdateMarks(ctx context.Context, req *pb.UpdateMarksRequest) (*pb.UpdateMarksResponse, error) {
	var marks []model.SubjectMark
	for _, m := range req.Marks {
		marks = append(marks, model.SubjectMark{
			SubjectName:      m.SubjectName,
			Mark:             m.Mark,
			GivenByTeacherID: m.GivenByTeacherId,
		})
	}
	err := h.service.UpdateMarks(req.StudentId, marks)
	if err != nil {
		return &pb.UpdateMarksResponse{Success: false, Message: err.Error()}, nil
	}
	return &pb.UpdateMarksResponse{Success: true, Message: "marks updated"}, nil
}

func (h *StudentHandler) GetMarksForStudent(ctx context.Context, req *pb.GetMarksForStudentRequest) (*pb.GetMarksForStudentResponse, error) {
	marks, err := h.service.GetMarks(req.StudentId)
	if err != nil {
		return nil, err
	}
	var pbMarks []*pb.SubjectMark
	for _, m := range marks {
		pbMarks = append(pbMarks, &pb.SubjectMark{
			SubjectName:      m.SubjectName,
			Mark:             m.Mark,
			GivenByTeacherId: m.GivenByTeacherID,
		})
	}
	return &pb.GetMarksForStudentResponse{Marks: pbMarks}, nil
}
