package config

import (
	"github.com/karanparmar2o/student-dashboard/internal/teacher/handler"
	"github.com/karanparmar2o/student-dashboard/internal/teacher/repository"
	"github.com/karanparmar2o/student-dashboard/internal/teacher/service"
)

type TeacherConfig struct {
	GRPCPort string
	HTTPPort string
	Handler  *handler.TeacherHandler
}

// LoadTeacherConfig initializes repo, service, handler & ports
func LoadTeacherConfig() *TeacherConfig {
	// Initialize repo
	repo := repository.NewMemoryTeacherRepo()

	// Initialize service
	svc := service.NewTeacherService(repo)

	// Initialize handler
	h := handler.NewTeacherHandler(svc)

	return &TeacherConfig{
		GRPCPort: ":50052",
		HTTPPort: ":8081",
		Handler:  h,
	}
}
