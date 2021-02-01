package database

//NewPagerResult ...
func NewPagerResult(q PagerQuery) (r PagerResult) {
	return PagerResult{
		PageIndex: q.PageIndex,
		PageSize:  q.PageSize,
	}
}

//PagerResult 分页查询返回值
type PagerResult struct {
	Total     int64 `json:"total,string"`
	PageSize  int   `json:"pagesize"`
	PageIndex int   `json:"pageindex"`
}

//PagerQuery 分页查询参数
type PagerQuery struct {
	PageSize  int `json:"pagesize"`
	PageIndex int `json:"pageindex"`
}
