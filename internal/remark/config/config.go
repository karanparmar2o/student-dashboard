package config

import (
	"github.com/karanparmar2o/student-dashboard/internal/remark/handler"
	"github.com/karanparmar2o/student-dashboard/internal/remark/repository"
	"github.com/karanparmar2o/student-dashboard/internal/remark/service"
)

type RemarkConfig struct {
	GRPCPort string
	HTTPPort string
	Handler  *handler.RemarkHandler
}

// LoadRemarkConfig initializes repo, service, handler & ports
func LoadRemarkConfig() *RemarkConfig {
	repo := repository.NewMemoryRemarkRepo()
	svc := service.NewRemarkService(repo)
	h := handler.NewRemarkHandler(svc)

	return &RemarkConfig{
		GRPCPort: ":50053",
		HTTPPort: ":8082",
		Handler:  h,
	}
}

// LoadRemarkConfigWithRepo allows custom repository initialization
