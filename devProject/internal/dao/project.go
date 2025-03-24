package dao

import (
	"context"
	"fmt"
	"test.com/devProject/internal/data/pro"
	"test.com/devProject/internal/database"
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

func (p *ProjectDao) UpdateProject(ctx context.Context, pro *pro.MsProject) error {
	return p.conn.Session(ctx).Updates(pro).Error
}

func (p *ProjectDao) DeleteProjectCollect(ctx context.Context, memId int64, projectCode int64) error {
	return p.conn.Session(ctx).Where("member_code = ? and project_code = ?", memId, projectCode).Delete(&pro.MsProjectCollection{}).Error
}

func (p *ProjectDao) SaveProjectCollect(ctx context.Context, pc *pro.MsProjectCollection) error {
	return p.conn.Session(ctx).Save(pc).Error
}

func (p *ProjectDao) UpdateDeletedProject(ctx context.Context, code int64, deleted bool) error {
	session := p.conn.Session(ctx)
	var err error
	if deleted {
		err = session.Model(&pro.MsProject{}).Where("id =?", code).Update("deleted", 1).Error
	} else {
		err = session.Model(&pro.MsProject{}).Where("id =?", code).Update("deleted", 0).Error
	}
	return err
}

func (p *ProjectDao) FindProjectByPidAndMemId(ctx context.Context, projectCode int64, memberId int64) (*pro.ProjectAndMember, error) {
	var pms pro.ProjectAndMember
	session := p.conn.Session(ctx)
	sql := "select a.*, b.project_code, b.member_code, b.join_time, b.is_owner, b.authorize from ms_projects a, ms_project_members b where a.id = b.project_code and b.member_code =? and b.project_code=? limit 1"
	raw := session.Raw(sql, memberId, projectCode)
	err := raw.Scan(&pms).Error

	return &pms, err
}

func (p *ProjectDao) FindCollectByPidAndMemId(ctx context.Context, projectCode int64, memberId int64) (bool, error) {
	var count int64
	session := p.conn.Session(ctx)
	sql := fmt.Sprintf("select count(*) from ms_project_collections where member_code=? and project_code=?")
	raw := session.Raw(sql, memberId, projectCode)
	err := raw.Scan(&count).Error
	return count > 0, err
}

func (p *ProjectDao) SaveProject(conn database.DbConn, ctx context.Context, pr *pro.MsProject) error {
	p.conn = conn.(*gorms.GormConn)
	return p.conn.Tx(ctx).Save(pr).Error
}

func (p *ProjectDao) SaveProjectMember(conn database.DbConn, ctx context.Context, pm *pro.MsProjectMember) error {
	p.conn = conn.(*gorms.GormConn)
	return p.conn.Tx(ctx).Save(pm).Error
}

func (p *ProjectDao) FindCollectProjectByMemId(ctx context.Context, memberId int64, page int64, size int64) ([]*pro.ProjectAndMember, int64, error) {
	var pms []*pro.ProjectAndMember
	session := p.conn.Session(ctx)
	index := (page - 1) * size
	sql := fmt.Sprintf("select * from ms_projects where id in (select project_code from ms_project_collections where member_code=?) order by sort limit?,?")
	raw := session.Raw(sql, memberId, index, size)
	raw.Scan(&pms)
	var count int64
	query := fmt.Sprintf("member_code=?")
	err := session.Model(&pro.MsProjectCollection{}).Where(query, memberId).Count(&count).Error
	return pms, count, err
}

func (p *ProjectDao) FindProjectByMemId(ctx context.Context, memId int64, condition string, page int64, size int64) ([]*pro.ProjectAndMember, int64, error) {
	var pms []*pro.ProjectAndMember
	session := p.conn.Session(ctx)
	index := (page - 1) * size
	sql := fmt.Sprintf("select * from ms_projects a, ms_project_members b where a.id = b.project_code and b.member_code =? %s order by sort limit ?,?", condition)
	// TODO 这里需要优化
	raw := session.Raw(sql, memId, index, size)
	raw.Scan(&pms)
	var total int64
	query := fmt.Sprintf("select count(*) from ms_projects a, ms_project_members b where a.id = b.project_code and b.member_code =? %s", condition)
	tx := session.Raw(query, memId)
	err := tx.Scan(&total).Error
	return pms, total, err
}
