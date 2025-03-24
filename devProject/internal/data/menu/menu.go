package menu

import "github.com/jinzhu/copier"

type MsProjectMenu struct {
	Id         int64  `gorm:"primaryKey;autoIncrement;column:id;comment:菜单id" json:"id"`
	Pid        int64  `gorm:"column:pid;not null;default:0;comment:父级id" json:"pid"`
	Title      string `gorm:"column:title;type:varchar(100);not null;default:'';comment:名称" json:"title"`
	Icon       string `gorm:"column:icon;type:varchar(100);not null;default:'';comment:菜单图标" json:"icon"`
	Url        string `gorm:"column:url;type:varchar(400);not null;default:'';comment:菜单链接" json:"url"`
	FlePath    string `gorm:"column:file_path;type:varchar(200);default:null;comment:文件路径" json:"filePath"`
	Params     string `gorm:"column:params;type:varchar(500);default:'';comment:连接参数" json:"params"`
	Node       string `gorm:"column:node;type:varchar(500);default:'#';comment:权限节点" json:"node"`
	Sort       uint32 `gorm:"column:sort;default:0;comment:菜单排序" json:"sort"`
	Status     int8   `gorm:"column:status;default:1;comment:状态(0:禁用,1:启用)" json:"status"`
	CreateBy   int64  `gorm:"column:create_by;not null;default:0;comment:创建人" json:"createBy"`
	IsInner    int8   `gorm:"column:is_inner;default:0;comment:是否内页" json:"isInner"`
	Values     string `gorm:"column:values;type:varchar(255);default:null;comment:参数默认值" json:"values"`
	ShowSlider int8   `gorm:"column:show_slider;default:1;comment:是否显示侧边栏" json:"show_slider"`
}

func (msProjectMenu *MsProjectMenu) TableName() string {
	return "ms_project_menus"
}

type MsProjectMenuChild struct {
	MsProjectMenu
	StatusText string `json:"statusText"`
	InnerText  string `json:"innerText"`
	FullUrl    string `json:"fullUrl"`
	Children   []*MsProjectMenuChild
}

func CovertChild(pms []*MsProjectMenu) []*MsProjectMenuChild {
	var pmcs []*MsProjectMenuChild

	_ = copier.Copy(&pmcs, pms)
	for _, v := range pmcs {
		v.StatusText = getStatus(v.Status)
		v.InnerText = getInnerText(v.IsInner)
		v.FullUrl = getFullUrl(v.Url, v.Params, v.Values)
	}

	var childPmcs []*MsProjectMenuChild
	//递归
	for _, v := range pmcs {
		// 存放根节点
		if v.Pid == 0 {
			pmc := &MsProjectMenuChild{}
			_ = copier.Copy(pmc, v)
			childPmcs = append(childPmcs, pmc)
		}
	}
	toChild(childPmcs, pmcs)
	return childPmcs
}

func toChild(childPmcs []*MsProjectMenuChild, pmcs []*MsProjectMenuChild) {
	for _, pmc := range childPmcs {
		for _, pm := range pmcs {
			if pmc.Id == pm.Pid {
				child := &MsProjectMenuChild{}
				_ = copier.Copy(child, pm)
				pmc.Children = append(pmc.Children, child)
			}
		}
		toChild(pmc.Children, pmcs)
	}
}

func getFullUrl(url string, params string, values string) string {
	if (params != "" && values != "") || values != "" {
		return url + "/" + values
	}
	return url
}

func getInnerText(inner int8) string {
	if inner == 0 {
		return "导航"
	}
	if inner == 1 {
		return "内页"
	}
	return ""
}

func getStatus(status int8) string {
	if status == 0 {
		return "禁用"
	}
	if status == 1 {
		return "使用中"
	}
	return ""
}
