package pro

import (
	"test.com/devCommon/encrypts"
	"test.com/devCommon/tms"
	"test.com/devProject/internal/data/task"
	"test.com/devProject/pkg/model"
)

// MsProject TODO: 数据库映射表，后面使用 gorm gen代替
type MsProject struct {
	Id                 int64   `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Cover              string  `gorm:"column:cover" json:"cover"`
	Name               string  `gorm:"column:name" json:"name"`
	Description        string  `gorm:"column:description" json:"description"`
	AccessControlType  int     `gorm:"column:access_control_type;default:0" json:"accessControlType"`
	WhiteList          string  `gorm:"column:white_list" json:"whiteList"`
	Sort               int     `gorm:"column:sort;default:0" json:"sort"`
	Deleted            int     `gorm:"column:deleted;default:0" json:"deleted"`
	TemplateCode       int     `gorm:"column:template_code" json:"templateCode"`
	Schedule           float64 `gorm:"column:schedule;default:0.00" json:"schedule"`
	CreateTime         int64   `gorm:"column:create_time" json:"createTime"`
	OrganizationCode   int64   `gorm:"column:organization_code" json:"organizationCode"`
	DeletedTime        string  `gorm:"column:deleted_time" json:"deletedTime"`
	Private            int     `gorm:"column:private;default:1" json:"private"`
	Prefix             string  `gorm:"column:prefix" json:"prefix"`
	OpenPrefix         int     `gorm:"column:open_prefix;default:0" json:"openPrefix"`
	Archive            int     `gorm:"column:archive;default:0" json:"archive"`
	ArchiveTime        int64   `gorm:"column:archive_time" json:"archiveTime"`
	OpenBeginTime      int     `gorm:"column:open_begin_time;default:0" json:"openBeginTime"`
	OpenTaskPrivate    int     `gorm:"column:open_task_private;default:0" json:"openTaskPrivate"`
	TaskBoardTheme     string  `gorm:"column:task_board_theme;default:'default'" json:"taskBoardTheme"`
	BeginTime          int64   `gorm:"column:begin_time" json:"beginTime"`
	EndTime            int64   `gorm:"column:end_time" json:"endTime"`
	AutoUpdateSchedule int     `gorm:"column:auto_update_schedule;default:0" json:"autoUpdateSchedule"`
}

func (*MsProject) TableName() string {
	return "ms_projects"
}

type MsProjectMember struct {
	Id          int64  `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	ProjectCode int64  `gorm:"column:project_code" json:"projectCode"`
	MemberCode  int64  `gorm:"column:member_code" json:"memberCode"`
	JoinTime    int64  `gorm:"column:join_time" json:"joinTime"`
	IsOwner     int64  `gorm:"column:is_owner;default:0" json:"isOwner"`
	Authorize   string `gorm:"column:authorize;type:varchar(255)" json:"authorize"`
}

func (*MsProjectMember) TableName() string {
	return "ms_project_members"
}

type MsProjectCollection struct {
	Id          int64 `gorm:"primaryKey;column:id" json:"id"`
	ProjectCode int64 `gorm:"column:project_code" json:"projectCode"`
	MemberCode  int64 `gorm:"column:member_code" json:"memberCode"`
	CreateTime  int64 `gorm:"column:create_time" json:"createTime"`
}

func (*MsProjectCollection) TableName() string {
	return "ms_project_collections"
}

type ProjectAndMember struct {
	MsProject
	ProjectCode int64  `json:"projectCode"`
	MemberCode  int64  `json:"memberCode"`
	JoinTime    int64  `json:"joinTime"`
	IsOwner     int64  `json:"isOwner"`
	Authorize   string `json:"authorize"`
	OwnerName   string `json:"ownerName"`
	Collected   int    `json:"collected"`
}

func (m *ProjectAndMember) GetAccessControlType() string {
	if m.AccessControlType == 0 {
		return "open"
	}
	if m.AccessControlType == 1 {
		return "private"
	}
	if m.AccessControlType == 2 {
		return "custom"
	}
	return ""
}

func ToMap(orgs []*ProjectAndMember) map[int64]*ProjectAndMember {
	m := make(map[int64]*ProjectAndMember)
	for _, org := range orgs {
		m[org.Id] = org
	}
	return m
}

type MsProjectTemplate struct {
	Id               int64  `gorm:"primaryKey;column:id" json:"id"`
	Name             string `gorm:"column:name" json:"name"`
	Description      string `gorm:"column:description" json:"description"`
	Sort             int    `gorm:"column:sort" json:"sort"`
	CreateTime       int64  `gorm:"column:create_time" json:"createTime"`
	OrganizationCode int64  `gorm:"column:organization_code" json:"organizationCode"`
	Cover            string `gorm:"column:cover" json:"cover"`
	MemberCode       int64  `gorm:"column:member_code" json:"memberCode"`
	IsSystem         int    `gorm:"column:is_system" json:"isSystem"`
}

func (*MsProjectTemplate) TableName() string {
	return "ms_project_templates"
}

type ProjectTemplateAll struct {
	Id               int                        `json:"id"`
	Name             string                     `json:"name"`
	Description      string                     `json:"description"`
	Sort             int                        `json:"sort"`
	CreateTime       string                     `json:"createTime"`
	OrganizationCode string                     `json:"organizationCode"`
	Cover            string                     `json:"cover"`
	MemberCode       string                     `json:"memberCode"`
	IsSystem         int                        `json:"isSystem"`
	TaskStages       []*task.TaskStagesOnlyName `json:"taskStages"`
	Code             string                     `json:"code"`
}

func (pt *MsProjectTemplate) Covert(taskStages []*task.TaskStagesOnlyName) *ProjectTemplateAll {
	organizationCode, _ := encrypts.EncryptInt64(pt.OrganizationCode, model.AESKey)
	memberCode, _ := encrypts.EncryptInt64(pt.MemberCode, model.AESKey)
	code, _ := encrypts.EncryptInt64(pt.Id, model.AESKey)
	pta := &ProjectTemplateAll{
		Id:               int(pt.Id),
		Name:             pt.Name,
		Description:      pt.Description,
		Sort:             pt.Sort,
		CreateTime:       tms.FormatByMilli(pt.CreateTime),
		OrganizationCode: organizationCode,
		Cover:            pt.Cover,
		MemberCode:       memberCode,
		IsSystem:         pt.IsSystem,
		TaskStages:       taskStages,
		Code:             code,
	}
	return pta
}

func ToProjectTemplateIds(pts []MsProjectTemplate) []int64 {
	var ids []int64
	for _, v := range pts {
		ids = append(ids, v.Id)
	}
	return ids
}
