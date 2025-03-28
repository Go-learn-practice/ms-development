package dao

import (
	"context"
	"test.com/devProject/internal/data/pro"
	"test.com/devProject/internal/database/gorms"
)

type ProjectTemplateDao struct {
	conn *gorms.GormConn
}

func NewProjectTemplateDao() *ProjectTemplateDao {
	return &ProjectTemplateDao{
		conn: gorms.New(),
	}
}

func (p *ProjectTemplateDao) FindProjectTemplateSystem(ctx context.Context, page int64, size int64) (pts []pro.MsProjectTemplate, total int64, err error) {
	session := p.conn.Session(ctx)
	err = session.
		Model(&pro.MsProjectTemplate{}).
		Where("is_system=?", 1).
		Limit(int(size)).
		Offset(int((page - 1) * size)).
		Find(&pts).
		Error
	if err != nil {
		return pts, total, err
	}
	err = session.
		Model(&pro.MsProjectTemplate{}).
		Where("is_system=?", 1).
		Count(&total).Error
	return pts, total, nil
}

func (p *ProjectTemplateDao) FindProjectTemplateCustom(ctx context.Context, memId int64, organizationCode int64, page int64, size int64) (pts []pro.MsProjectTemplate, total int64, err error) {
	session := p.conn.Session(ctx)
	err = session.
		Model(&pro.MsProjectTemplate{}).
		Where("is_system=? and member_code=? and organization_code=?", 0, memId, organizationCode).
		Limit(int(size)).
		Offset(int((page - 1) * size)).
		Find(&pts).
		Error
	if err != nil {
		return pts, total, err
	}
	err = session.
		Model(&pro.MsProjectTemplate{}).
		Where("is_system=? and member_code=? and organization_code=?", 0, memId, organizationCode).
		Count(&total).Error
	return pts, total, err
}

func (p *ProjectTemplateDao) FindProjectTemplateAll(ctx context.Context, organizationCode int64, page int64, size int64) (pts []pro.MsProjectTemplate, total int64, err error) {
	session := p.conn.Session(ctx)
	err = session.
		Model(&pro.MsProjectTemplate{}).
		Where("organization_code=?", organizationCode).
		Limit(int(size)).
		Offset(int((page - 1) * size)).
		Find(&pts).
		Error
	if err != nil {
		return pts, total, err
	}
	err = session.
		Model(&pro.MsProjectTemplate{}).
		Where("organization_code=?", organizationCode).
		Count(&total).Error
	return pts, total, err
}
