package member

// Member 员工信息表
type Member struct {
	Id              int64  `gorm:"primaryKey" json:"id"`
	Account         string `gorm:"column:account" json:"account"`
	Password        string `gorm:"column:password" json:"password"`
	Name            string `gorm:"column:name" json:"name"`
	Mobile          string `gorm:"column:mobile" json:"mobile"`
	RealName        string `gorm:"column:real_name" json:"realName"`
	CreateTime      int64  `gorm:"column:create_time" json:"createTime"`
	Status          int8   `gorm:"column:status" json:"status"`
	LastLoginTime   int64  `gorm:"column:last_login_time" json:"lastLoginTime"`
	Sex             int8   `gorm:"column:sex" json:"sex"`
	Avatar          string `gorm:"column:avatar" json:"avatar"`
	Idcard          string `gorm:"column:idcard" json:"idcard"`
	Province        int    `gorm:"column:province" json:"province"`
	City            int    `gorm:"column:city" json:"city"`
	Area            int    `gorm:"column:area" json:"area"`
	Address         string `gorm:"column:address" json:"address"`
	Description     string `gorm:"column:description" json:"description"`
	Email           string `gorm:"column:email" json:"email"`
	DingtalkOpenId  string `gorm:"column:dingtalk_openid" json:"dingtalkOpenId"`
	DingtalkUnionId string `gorm:"column:dingtalk_unionid" json:"dingtalkUnionId"`
	DingtalkUserId  string `gorm:"column:dingtalk_userid" json:"dingtalkUserId"`
}

func (member *Member) TableName() string {
	return "ms_member"
}
