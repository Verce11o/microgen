package internal

import (
	"context"
)

type Repository interface {
	CreateExample(ctx context.Context, in *pb.CreateExampleRequest) (string, error)
	GetExample(ctx context.Context, in *pb.GetExampleRequest) (models.Example, error)
	UpdateExample(ctx context.Context, in *pb.UpdateExampleRequest) (models.Example, error)
	DeleteExample(ctx context.Context, in *pb.DeleteExampleRequest) error
}
