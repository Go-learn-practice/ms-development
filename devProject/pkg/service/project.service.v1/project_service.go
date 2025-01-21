package project_service_v1

import (
	"context"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"test.com/devCommon/errs"
	"test.com/devGrpc/project"
	"test.com/devProject/internal/dao"
	"test.com/devProject/internal/data/menu"
	"test.com/devProject/internal/database/tran"
	"test.com/devProject/internal/repo"
	"test.com/devUser/pkg/model"
)

type ProjectService struct {
	project.UnimplementedProjectServiceServer
	cache       repo.Cache
	transaction tran.Transaction
	menuRepo    repo.MenuRepo
}

func New() *ProjectService {
	return &ProjectService{
		cache:       dao.RedisCacheInstance,
		transaction: dao.NewTransaction(),
		menuRepo:    dao.NewMenuDao(),
	}
}

func (p *ProjectService) Index(context.Context, *project.IndexRequest) (*project.IndexResponse, error) {
	pms, err := p.menuRepo.FindMenus(context.Background())
	zap.L().Info("menus数据获取成功")
	if err != nil {
		zap.L().Error("Index db FindMenus error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	childs := menu.CovertChild(pms)
	var mms []*project.MenuMessage
	_ = copier.Copy(&mms, childs)
	return &project.IndexResponse{Menus: mms}, nil
}
