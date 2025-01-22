package project_service_v1

import (
	"context"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"test.com/devCommon/encrypts"
	"test.com/devCommon/errs"
	"test.com/devGrpc/project"
	"test.com/devProject/internal/dao"
	"test.com/devProject/internal/data/menu"
	"test.com/devProject/internal/database/tran"
	"test.com/devProject/internal/repo"
	"test.com/devProject/pkg/model"
)

type ProjectService struct {
	project.UnimplementedProjectServiceServer
	cache       repo.Cache
	transaction tran.Transaction
	menuRepo    repo.MenuRepo
	projectRepo repo.ProjectRepo
}

func New() *ProjectService {
	return &ProjectService{
		cache:       dao.RedisCacheInstance,
		transaction: dao.NewTransaction(),
		menuRepo:    dao.NewMenuDao(),
		projectRepo: dao.NewProjectDao(),
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

func (p *ProjectService) FindProjectByMemId(ctx context.Context, msg *project.ProjectRpcRequest) (*project.MyProjectResponse, error) {
	memberId := msg.MemberId
	page := msg.Page
	pageSize := msg.PageSize
	pms, total, err := p.projectRepo.FindProjectByMemId(ctx, memberId, page, pageSize)
	if err != nil {
		zap.L().Error("project FindProjectByMemId err", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if pms == nil {
		return &project.MyProjectResponse{Pm: []*project.ProjectMessage{}, Total: total}, nil
	}
	var pmm []*project.ProjectMessage
	_ = copier.Copy(&pmm, pms)
	for _, v := range pmm {
		v.Code, _ = encrypts.EncryptInt64(v.Id, model.AESKey)
	}
	return &project.MyProjectResponse{Pm: pmm, Total: total}, nil
}
