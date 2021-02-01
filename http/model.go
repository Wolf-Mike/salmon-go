package sgchttp

//ResponseVto 返回消息的格式
type ResponseVto struct {
	// 代码
	Code int `json:"code" example:"200"`
	// 数据集
	Data interface{} `json:"data"`
	// 消息
	Msg       string `json:"msg"`
	RequestID string `json:"requestId"`
}

//PageVto 返回分页格式
type PageVto struct {
	List      interface{} `json:"list"`
	Count     int         `json:"count"`
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
}

//ReturnOk 返回ok
func (res *ResponseVto) ReturnOk() *ResponseVto {
	res.Code = 200
	return res
}

//ReturnError 返回错误
func (res *ResponseVto) ReturnError(code int) *ResponseVto {
	res.Code = code
	return res
}
