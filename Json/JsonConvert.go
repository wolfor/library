// Json project jsonConvert.go
package Json

import (
	json "encoding/json"
	"errors"
	//	"fmt"
	io "io/ioutil"
	"library/Path"
)

type JsonConvert struct {
}

func (this *JsonConvert) JsonToStruct(filePath string, targetStruct interface{}) error {
	var isExist bool
	var err error

	if isExist, err = Path.PathExists(filePath); !isExist || err != nil {
		return err
	}

	var jsonData []byte

	jsonData, err = io.ReadFile(filePath)

	if err != nil {
		return err
	}

	err = this.JsonToStruct2(jsonData, targetStruct)

	if err != nil {
		return err
	}

	return nil
}

//json字节数组转换结构体
func (this *JsonConvert) JsonToStruct2(buffer []byte, targetStruct interface{}) error {
	err := json.Unmarshal(buffer, targetStruct)

	return err
}

//json format string convert to struct
//call example:
func (this *JsonConvert) JsonToStruct3(jsonString string, targetStruct interface{}) error {
	if jsonString == "" {
		return errors.New("the jsonString input parameter is empty or null.")
	}

	buffer := []byte(jsonString)

	err := this.JsonToStruct2(buffer, targetStruct)

	return err
}

func (this *JsonConvert) StructToJson(sourceStruct interface{}) (string, error) {
	jsons, err := this.StructToBytes(sourceStruct)

	return string(jsons), err //byte[]转换成string 输出
}

func (this *JsonConvert) StructToBytes(sourceStruct interface{}) ([]byte, error) {
	jsons, err := json.Marshal(sourceStruct) //转换成JSON返回的是byte[]

	return jsons, err
}
