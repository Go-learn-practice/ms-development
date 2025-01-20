package repo

import (
	"context"
	"test.com/devUser/internal/data/member"
	"test.com/devUser/internal/database"
)

type MemberRepo interface {
	GetMemberByEmail(ctx context.Context, email string) (bool, error)
	GetMemberByAccount(ctx context.Context, account string) (bool, error)
	GetMemberByMobile(ctx context.Context, mobile string) (bool, error)
	SaveMember(ctx context.Context, conn database.DbConn, member *member.Member) error
	FindMember(ctx context.Context, account, pwd string) (member *member.Member, err error)
}
