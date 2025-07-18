package config

import (
	"github.com/karanparmar2o/student-dashboard/internal/student/handler"
	repository "github.com/karanparmar2o/student-dashboard/internal/student/repo"
	"github.com/karanparmar2o/student-dashboard/internal/student/service"
)

type Config struct {
	StudentHandler *handler.StudentHandler
}

func Load() (*Config, error) {
	// Initialize in order: repo → service → handler
	repo := repository.NewMemoryStudentRepo()
	svc := service.NewStudentService(repo)
	h := handler.NewStudentHandler(svc)

	return &Config{
		StudentHandler: h,
	}, nil
}
