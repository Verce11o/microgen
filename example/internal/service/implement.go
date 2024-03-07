package service

import (
	"context"
	"go.uber.org/zap"
)

type ExampleService struct {
	log *zap.SugaredLogger
}

func NewExampleService(log *zap.SugaredLogger) *ExampleService {
	return &ExampleService{log: log}
}

func (s *ExampleService) CreateExample(ctx context.Context) error {
	return nil
}

func (s *ExampleService) GetExample(ctx context.Context) error {
	return nil
}

func (s *ExampleService) UpdateExample(ctx context.Context) error {
	return nil
}

func (s *ExampleService) DeleteExample(ctx context.Context) error {
	return nil
}
