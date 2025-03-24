package pro

type Project struct {
	Id                 int64   `json:"id"`
	Cover              string  `json:"cover"`
	Name               string  `json:"name"`
	Description        string  `json:"description"`
	AccessControlType  int     `json:"accessControlType"`
	WhiteList          string  `json:"whiteList"`
	Order              int     `json:"order"`
	Deleted            int     `json:"deleted"`
	TemplateCode       string  `json:"templateCode"`
	Schedule           float64 `json:"schedule"`
	CreateTime         string  `json:"createTime"`
	OrganizationCode   int64   `json:"organizationCode"`
	DeletedTime        string  `json:"deletedTime"`
	Private            int     `json:"private"`
	Prefix             string  `json:"prefix"`
	OpenPrefix         int     `json:"openPrefix"`
	Archive            int     `json:"archive"`
	ArchiveTime        int64   `json:"archiveTime"`
	OpenBeginTime      int     `json:"openBeginTime"`
	OpenTaskPrivate    int     `json:"openTaskPrivate"`
	TaskBoardTheme     string  `json:"taskBoardTheme"`
	BeginTime          int64   `json:"beginTime"`
	EndTime            int64   `json:"endTime"`
	AutoUpdateSchedule int     `json:"autoUpdateSchedule"`
	Code               string  `json:"code"`
}

type ProjectMember struct {
	Id          int64  `json:"id"`
	ProjectCode int64  `json:"projectCode"`
	MemberCode  int64  `json:"memberCode"`
	JoinTime    int64  `json:"joinTime"`
	IsOwner     int64  `json:"isOwner"`
	Authorize   string `json:"authorize"`
}

type ProjectAndMember struct {
	Project
	ProjectCode int64  `json:"projectCode"`
	MemberCode  int64  `json:"memberCode"`
	JoinTime    int64  `json:"joinTime"`
	IsOwner     int64  `json:"isOwner"`
	Authorize   string `json:"authorize"`
	OwnerName   string `json:"ownerName"`
	Collected   string `json:"collected"`
}

type ProjectDetail struct {
	Project
	OwnerName   string `json:"ownerName"`
	Collected   string `json:"collected"`
	OwnerAvatar string `json:"ownerAvatar"`
}

type ProjectTemplate struct {
	Id               int                   `json:"id"`
	Name             string                `json:"name"`
	Description      string                `json:"description"`
	Sort             int                   `json:"sort"`
	CreateTime       string                `json:"createTime"`
	OrganizationCode int64                 `json:"organizationCode"`
	Cover            string                `json:"cover"`
	MemberCode       string                `json:"memberCode"`
	IsSystem         int                   `json:"isSystem"`
	TaskStages       []*TaskStagesOnlyName `json:"taskStages"`
	Code             string                `json:"code"`
}

type TaskStagesOnlyName struct {
	Name string `json:"name"`
}

type SaveProjectRequest struct {
	Name         string `json:"name" form:"name"`
	TemplateCode string `json:"templateCode" form:"templateCode"`
	Description  string `json:"description" form:"description"`
	Id           int    `json:"id" form:"id"`
}

type SaveProject struct {
	Id               int64  `json:"id"`
	Cover            string `json:"cover"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Code             string `json:"code"`
	CreateTime       string `json:"createTime"`
	TaskBoardTheme   string `json:"taskBoardTheme"`
	OrganizationCode string `json:"organizationCode"`
}

type ProjectReq struct {
	ProjectCode        int64  `json:"projectCode" form:"projectCode"`
	Cover              string `json:"cover" form:"cover"`
	Name               string `json:"name" form:"name"`
	Description        string `json:"description" form:"description"`
	Schedule           int    `json:"schedule" form:"schedule"`
	Private            int    `json:"private" form:"private"`
	OpenPrefix         int    `json:"openPrefix" form:"openPrefix"`
	Prefix             string `json:"prefix" form:"prefix"`
	OpenTaskPrivate    int    `json:"openTaskPrivate" form:"openTaskPrivate"`
	TaskBoardTheme     string `json:"taskBoardTheme" form:"taskBoardTheme"`
	AutoUpdateSchedule int    `json:"autoUpdateSchedule" form:"autoUpdateSchedule"`
}
