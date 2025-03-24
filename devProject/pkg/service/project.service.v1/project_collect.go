package project_service_v1

import (
	"context"
	"go.uber.org/zap"
	"strconv"
	"test.com/devCommon/encrypts"
	"test.com/devCommon/errs"
	"test.com/devGrpc/project"
	"test.com/devProject/internal/data/pro"
	"test.com/devProject/pkg/model"
	"time"
)

// UpdateCollectProject 收藏项目、取消收藏项目
func (p *ProjectService) UpdateCollectProject(ctx context.Context, msg *project.ProjectRpcRequest) (*project.CollectProjectResponse, error) {
	projectCodeStr, _ := encrypts.Decrypt(msg.ProjectCode, model.AESKey)
	projectCode, _ := strconv.ParseInt(projectCodeStr, 10, 64)
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var err error

	// 1. 收藏
	if "collect" == msg.CollectType {
		pc := &pro.MsProjectCollection{
			ProjectCode: projectCode,
			MemberCode:  msg.MemberId,
			CreateTime:  time.Now().UnixMilli(), // 毫秒
		}
		err = p.projectRepo.SaveProjectCollect(c, pc)
	}
	// 2. 取消收藏
	if "cancel" == msg.CollectType {
		err = p.projectRepo.DeleteProjectCollect(c, projectCode, msg.MemberId)
	}
	if err != nil {
		zap.L().Error("project UpdateCollectProject SaveProjectCollect error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	return &project.CollectProjectResponse{}, nil
}
