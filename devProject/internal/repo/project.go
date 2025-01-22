package repo

import (
	"context"
	"test.com/devProject/internal/data/pro"
)

type ProjectRepo interface {
	FindProjectByMemId(ctx context.Context, memId, page, size int64) ([]*pro.ProjectAndMember, int64, error)
}
