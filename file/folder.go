package file

import "os"

//PathExist 判断文件或文件夹是否存在
func PathExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

//FolderCreate 创建文件夹
func FolderCreate(path string) (err error) {
	if !PathExist(path) {
		err = os.Mkdir(path, os.ModePerm)
	}
	return
}
