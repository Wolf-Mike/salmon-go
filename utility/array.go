package utility

//Contains 字符串切片s是否包含字符串i
func Contains(s []string, i string) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}
	return false
}
