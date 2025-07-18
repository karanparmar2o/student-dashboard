package handler

import (
	"context"

	"github.com/karanparmar2o/student-dashboard/api/remark"
	"github.com/karanparmar2o/student-dashboard/internal/remark/service"
)

type RemarkHandler struct {
	remark.UnimplementedRemarkServiceServer
	svc *service.RemarkService
}

func NewRemarkHandler(s *service.RemarkService) *RemarkHandler {
	return &RemarkHandler{svc: s}
}

func (h *RemarkHandler) AddRemark(ctx context.Context, req *remark.AddRemarkRequest) (*remark.AddRemarkResponse, error) {
	return h.svc.AddRemark(req)
}

func (h *RemarkHandler) GetRemarksForStudent(ctx context.Context, req *remark.GetRemarksForStudentRequest) (*remark.GetRemarksForStudentResponse, error) {
	return h.svc.GetRemarksForStudent(req)
}

func (h *RemarkHandler) ApproveRemark(ctx context.Context, req *remark.ApproveRemarkRequest) (*remark.ApproveRemarkResponse, error) {
	return h.svc.ApproveRemark(req)
}
