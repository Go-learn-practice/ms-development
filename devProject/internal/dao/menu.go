package dao

import (
	"context"
	"test.com/devProject/internal/data/menu"
	"test.com/devProject/internal/database/gorms"
)

type MenuDao struct {
	conn *gorms.GormConn
}

func (m *MenuDao) FindMenus(ctx context.Context) (pms []*menu.MsProjectMenu, err error) {
	session := m.conn.Session(ctx)
	err = session.Order("pid,sort asc, id asc").Find(&pms).Error
	return
}

func NewMenuDao() *MenuDao {
	return &MenuDao{
		conn: gorms.New(),
	}
}
