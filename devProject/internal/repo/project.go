package repo

import (
	"context"
	"test.com/devProject/internal/data/pro"
	"test.com/devProject/internal/database"
)

type ProjectRepo interface {
	FindProjectByMemId(ctx context.Context, memId int64, condition string, page int64, size int64) ([]*pro.ProjectAndMember, int64, error)
	FindCollectProjectByMemId(ctx context.Context, member int64, page int64, size int64) ([]*pro.ProjectAndMember, int64, error)
	SaveProject(conn database.DbConn, ctx context.Context, project *pro.MsProject) error
	SaveProjectMember(conn database.DbConn, ctx context.Context, pm *pro.MsProjectMember) error
	FindProjectByPidAndMemId(ctx context.Context, projectCode int64, memberId int64) (*pro.ProjectAndMember, error)
	FindCollectByPidAndMemId(ctx context.Context, projectCode int64, memberId int64) (bool, error)
	UpdateDeletedProject(ctx context.Context, code int64, deleted bool) error
	SaveProjectCollect(ctx context.Context, pc *pro.MsProjectCollection) error
	DeleteProjectCollect(ctx context.Context, memId int64, projectCode int64) error
	UpdateProject(ctx context.Context, project *pro.MsProject) error
}

type ProjectTemplateRepo interface {
	FindProjectTemplateSystem(ctx context.Context, page int64, size int64) ([]pro.MsProjectTemplate, int64, error)
	FindProjectTemplateCustom(ctx context.Context, memId int64, organizationCode int64, page int64, size int64) ([]pro.MsProjectTemplate, int64, error)
	FindProjectTemplateAll(ctx context.Context, organizationCode int64, page int64, size int64) ([]pro.MsProjectTemplate, int64, error)
}
