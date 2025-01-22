package dao

import (
	"context"
	"test.com/devProject/internal/data/pro"
	"test.com/devProject/internal/database/gorms"
)

type ProjectDao struct {
	conn *gorms.GormConn
}

func NewProjectDao() *ProjectDao {
	return &ProjectDao{
		conn: gorms.New(),
	}
}

func (p *ProjectDao) FindProjectByMemId(ctx context.Context, memId, page, size int64) ([]*pro.ProjectAndMember, int64, error) {
	var pms []*pro.ProjectAndMember
	session := p.conn.Session(ctx)
	index := (page - 1) * size
	// TODO 这里需要优化
	raw := session.Raw("select * from ms_projects a, ms_project_members b where a.id = b.project_code and b.member_code = ? limit ?, ?", memId, index, size)
	raw.Scan(&pms)
	var total int64
	err := session.Model(&pro.MsProjectMember{}).Where("member_code = ?", memId).Count(&total).Error
	return pms, total, err
}
