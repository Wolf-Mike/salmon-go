package database

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

//sgcTime 解决gorm中对mysql中datetime类型字段无法scan的问题
//参考https://blog.csdn.net/qq_37493556/article/details/105082422
type sgcTime time.Time

//UnmarshalJSON ...
func (t *sgcTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = sgcTime(t1)
	return err
}

//MarshalJSON ...
func (t sgcTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

//Value ...
func (t sgcTime) Value() (driver.Value, error) {
	// sgcTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format("2006-01-02 15:04:05"), nil
}

//Scan ...
func (t *sgcTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = sgcTime(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

//String ...
func (t *sgcTime) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
}
