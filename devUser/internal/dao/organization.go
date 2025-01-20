package dao

import (
	"context"
	"test.com/devUser/internal/data/organization"
	"test.com/devUser/internal/database"
	"test.com/devUser/internal/database/gorms"
)

type OrganizationDao struct {
	conn *gorms.GormConn
}

func NewOrganizationDao() *OrganizationDao {
	return &OrganizationDao{
		conn: gorms.New(),
	}
}

func (o *OrganizationDao) SaveOrganization(ctx context.Context, conn database.DbConn, org *organization.Organization) error {
	//interface 类型断言
	o.conn = conn.(*gorms.GormConn)
	err := o.conn.Tx(ctx).Create(org).Error
	return err
}

func (o *OrganizationDao) FindOrganizationByMemId(ctx context.Context, memId int64) ([]*organization.Organization, error) {
	var orgs []*organization.Organization
	err := o.conn.Session(ctx).Where("member_id=?", memId).Find(&orgs).Error
	return orgs, err
}
