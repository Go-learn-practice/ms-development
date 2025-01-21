package project

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
