package pro

// MsProject TODO: 数据库映射表，后面使用 gorm gen代替
type MsProject struct {
	Id                 int64
	Cover              string
	Name               string
	Description        string
	AccessControlType  int
	WhiteList          string
	Order              int
	Deleted            int
	TemplateCode       string
	Schedule           float64
	CreateTime         string
	OrganizationCode   int64
	DeletedTime        string
	Private            int
	Prefix             string
	OpenPrefix         int
	Archive            int
	ArchiveTime        int64
	OpenBeginTime      int
	OpenTaskPrivate    int
	TaskBoardTheme     string
	BeginTime          int64
	EndTime            int64
	AutoUpdateSchedule int
}

func (*MsProject) TableName() string {
	return "ms_projects"
}

type MsProjectMember struct {
	Id          int64
	ProjectCode int64
	MemberCode  int64
	JoinTime    int64
	IsOwner     int64
	Authorize   string
}

func (*MsProjectMember) TableName() string {
	return "ms_project_members"
}

type ProjectAndMember struct {
	MsProject
	ProjectCode int64
	MemberCode  int64
	JoinTime    int64
	IsOwner     int64
	Authorize   string
}
