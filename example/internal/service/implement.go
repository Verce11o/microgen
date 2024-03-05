package service

import (
	"go.uber.org/zap"
)

type ExampleService struct {
	log *zap.SugaredLogger
}

func NewExampleService(log *zap.SugaredLogger) *ExampleService {
	return &ExampleService{log: log}
}
