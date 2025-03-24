package task

type MsTaskStagesTemplate struct {
	Id                  int64  `gorm:"primaryKey;column:id" json:"id"`
	Name                string `gorm:"column:name;type:varchar(255)" json:"name"`
	ProjectTemplateCode int    `gorm:"column:project_template_code;default:0" json:"projectTemplateCode"`
	CreateTime          int64  `gorm:"column:create_time" json:"createTime"`
	Sort                int    `gorm:"column:sort;default:0" json:"sort"`
}

func (*MsTaskStagesTemplate) TableName() string {
	return "ms_task_stages_templates"
}

type TaskStagesOnlyName struct {
	Name string `json:"name"`
}

// CovertProjectMap
func CoverProjectMap(tasks []MsTaskStagesTemplate) map[int64][]*TaskStagesOnlyName {
	var tss = make(map[int64][]*TaskStagesOnlyName)
	for _, v := range tasks {
		ts := &TaskStagesOnlyName{}
		ts.Name = v.Name
		tss[int64(v.ProjectTemplateCode)] = append(tss[int64(v.ProjectTemplateCode)], ts)
	}
	return tss
}
