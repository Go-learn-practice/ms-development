package project_service_v1

import (
	"context"
	"test.com/devGrpc/project"
	"test.com/devProject/internal/dao"
	"test.com/devProject/internal/database/tran"
	"test.com/devProject/internal/repo"
)

type ProjectService struct {
	project.UnimplementedProjectServiceServer
	cache       repo.Cache
	transaction tran.Transaction
}

func New() *ProjectService {
	return &ProjectService{
		cache:       dao.RedisCacheInstance,
		transaction: dao.NewTransaction(),
	}
}

func (p *ProjectService) Index(context.Context, *project.IndexRequest) (*project.IndexResponse, error) {
	return &project.IndexResponse{}, nil
}
