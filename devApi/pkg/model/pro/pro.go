package pro

type MenusResponse struct {
	Id         int64            `json:"id"`
	Pid        int64            `json:"pid"`
	Title      string           `json:"title"`
	Icon       string           `json:"icon"`
	Url        string           `json:"url"`
	FilePath   string           `json:"filePath"`
	Params     string           `json:"params"`
	Node       string           `json:"node"`
	Sort       int32            `json:"sort"`
	Status     int32            `json:"status"`
	CreateBy   int64            `json:"createBy"`
	IsInner    int32            `json:"isInner"`
	Values     string           `json:"values"`
	ShowSlider int32            `json:"showSlider"`
	Children   []*MenusResponse `json:"children"`
}

type Project struct {
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
}
