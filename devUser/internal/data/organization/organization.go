package organization

// Organization 组织表
type Organization struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	Avatar      string `gorm:"column:avatar" json:"avatar"`
	Description string `gorm:"column:description" json:"description"`
	MemberId    int64  `gorm:"column:member_id" json:"memberId"`
	CreateTime  int64  `gorm:"column:create_time" json:"createTime"`
	Personal    int32  `gorm:"column:personal" json:"personal"`
	Address     string `gorm:"column:address" json:"address"`
	Province    int32  `gorm:"column:province" json:"province"`
	City        int32  `gorm:"column:city" json:"city"`
	Area        int32  `gorm:"column:area" json:"area"`
}

func (organization *Organization) TableName() string {
	return "ms_organization"
}

func ToMap(orgs []*Organization) map[int64]*Organization {
	m := make(map[int64]*Organization)

	for _, org := range orgs {
		m[org.Id] = org
	}
	return m
}
