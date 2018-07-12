package FileHelper

import (
	"io"
	"os"

	"github.com/jszwec/csvutil"
)

type CSVHelper struct {
	filePath string
	CSVFile  *os.File
}

func NewCSVHelper(filePath string) *CSVHelper {
	helper := new(CSVHelper)
	helper.filePath = filePath
	helper.CSVFile = helper.open(filePath)

	return helper
}

func (this *CSVHelper) open(filePath string) *os.File {
	fi, err := os.Open(filePath)

	if err != nil {
		return nil
	}

	return fi
}

//读csv文件
func (this *CSVHelper) ReadFile() []byte {

	if this.CSVFile == nil {
		return nil
	}

	defer this.CSVFile.Close()

	chunks := make([]byte, 0)
	buff := make([]byte, 1024)

	for {
		count, er := this.CSVFile.Read(buff)

		if er != nil && er != io.EOF {
			panic(er)
		}

		if count == 0 {
			break
		}

		chunks = append(chunks, buff[:count]...)
	}

	return chunks
}

//csv文件二进制数组转换为结构体
func (this *CSVHelper) UnmarshalBytes(fileBuff []byte, object interface{}) error {
	err := csvutil.Unmarshal(fileBuff, object)

	return err
}

//csv文件转换为结构体
func (this *CSVHelper) UnmarshalFile(object interface{}) error {
	buff := this.ReadFile()

	return this.UnmarshalBytes(buff, object)
}
