package repo

import (
	"context"
	"test.com/devProject/internal/data/menu"
)

type MenuRepo interface {
	FindMenus(ctx context.Context) ([]*menu.MsProjectMenu, error)
}
