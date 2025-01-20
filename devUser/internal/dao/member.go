package dao

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"test.com/devUser/internal/data/member"
	"test.com/devUser/internal/database"
	"test.com/devUser/internal/database/gorms"
)

type MemberDao struct {
	conn *gorms.GormConn
}

func NewMemberDao() *MemberDao {
	return &MemberDao{conn: gorms.New()}
}

func (m *MemberDao) GetMemberByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&member.Member{}).Where("email=?", email).Count(&count).Error
	return count > 0, err
}

func (m *MemberDao) GetMemberByMobile(ctx context.Context, mobile string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&member.Member{}).Where("mobile=?", mobile).Count(&count).Error
	return count > 0, err
}

func (m *MemberDao) GetMemberByAccount(ctx context.Context, account string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&member.Member{}).Where("account=?", account).Count(&count).Error
	return count > 0, err
}

func (m *MemberDao) SaveMember(ctx context.Context, conn database.DbConn, member *member.Member) error {
	//interface 类型断言
	m.conn = conn.(*gorms.GormConn)
	return m.conn.Tx(ctx).Create(member).Error
}

func (m *MemberDao) FindMember(ctx context.Context, account string, pwd string) (*member.Member, error) {
	var mem *member.Member
	err := m.conn.Session(ctx).Where("account=? and password=?", account, pwd).First(&mem).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return mem, err
}
