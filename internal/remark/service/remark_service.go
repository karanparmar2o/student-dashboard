package service

import (
	"github.com/karanparmar2o/student-dashboard/api/remark"
	"github.com/karanparmar2o/student-dashboard/internal/remark/repository"
)

type RemarkService struct {
	repo repository.RemarkRepo
}

func NewRemarkService(r repository.RemarkRepo) *RemarkService {
	return &RemarkService{repo: r}
}

func (s *RemarkService) AddRemark(req *remark.AddRemarkRequest) (*remark.AddRemarkResponse, error) {
	r := &remark.Remark{
		StudentId: req.StudentId,
		TeacherId: req.AddedByTeacherId,
		Text:      req.Text,
		Level:     req.Level,
		Approved:  false,
	}
	id, _ := s.repo.AddRemark(r)
	return &remark.AddRemarkResponse{
		Success: true,
		Message: "Remark added",
		Id:      id,
	}, nil
}

func (s *RemarkService) GetRemarksForStudent(req *remark.GetRemarksForStudentRequest) (*remark.GetRemarksForStudentResponse, error) {
	list, _ := s.repo.GetRemarksForStudent(req.StudentId)
	return &remark.GetRemarksForStudentResponse{
		Remarks: list,
	}, nil
}

func (s *RemarkService) ApproveRemark(req *remark.ApproveRemarkRequest) (*remark.ApproveRemarkResponse, error) {
	_ = s.repo.ApproveRemark(req.RemarkId)
	return &remark.ApproveRemarkResponse{
		Success: true,
		Message: "Remark approved",
	}, nil
}
