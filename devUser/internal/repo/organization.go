package repo

import (
	"context"
	"test.com/devUser/internal/data/organization"
	"test.com/devUser/internal/database"
)

type OrganizationRepo interface {
	SaveOrganization(ctx context.Context, conn database.DbConn, org *organization.Organization) error
	FindOrganizationByMemId(ctx context.Context, memId int64) ([]*organization.Organization, error)
}
