package Json

import (
	"errors"
	io "io/ioutil"
	"library/Path"
)

type JsonHelper struct{}

func NewJsonHelper() *JsonHelper {
	helper := new(JsonHelper)

	return helper
}

//读json文件
func (this *JsonHelper) ReadFile(filePath string) (fileContent string) {
	isExist, er := Path.PathExists(filePath)

	if !isExist || er != nil {
		panic(errors.New("file not exist"))
	}

	data, err := io.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	datajson := []byte(data)

	fileContent = string(datajson)

	return fileContent
}
