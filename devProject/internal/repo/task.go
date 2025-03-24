package repo

import (
	"context"
	"test.com/devProject/internal/data/task"
)

type TaskStagesTemplateRepo interface {
	FindInProTemIds(ctx context.Context, ids []int64) ([]task.MsTaskStagesTemplate, error)
}
