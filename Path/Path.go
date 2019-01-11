// Path project Path.go
package Path

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

//获取指定目录父级目录
func GetParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

//获取当前执行目录
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

//创建目录
func CreateDir(dir string) bool {
	err := os.MkdirAll(dir, os.ModePerm)

	if err != nil {
		return false
	} else {
		return true
	}
}

//创建目录
func CreateLocalDir(dir string) bool {
	isExist, _ := PathExists(dir)

	if isExist {
		return true
	} else {
		return CreateDir(dir)
	}
}

//获取当前程序EXE所在目录
func GetCurrentPath() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return ""
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return ""
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return ""
	}
	return string(path[0 : i+1])

}

//文件或目录是否存在。true表存在，false表不存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//文件或目录是否存在。true表存在，false表不存在
func IsExists(path string) bool {
	isExists, err := PathExists(path)

	if err != nil {
		return false
	}

	return isExists
}

//根据文件路径字符串获取文件路径
//例如:filePath="/usr/local/test.txt"
//返回结果:/usr/local
func GetFileDir(filePath string) string {
	if strings.TrimSpace(filePath) == "" {
		return ""
	}

	dir := filePath[:strings.LastIndex(filePath, "/")]

	return dir
}

//根据文件路径字符串获取文件名
//例如:filePath="/usr/local/test.txt"
//返回结果:test.txt
func GetFileName(filePath string) string {
	if strings.TrimSpace(filePath) == "" {
		return ""
	}

	fileName := filePath[strings.LastIndex(filePath, "/")+1:]

	return fileName
}

//根据文件路径字符串获取文件名
//例如:filePath="/usr/local/test.txt"
//返回结果:test
func GetFileName2(filePath string) string {
	if strings.TrimSpace(filePath) == "" {
		return ""
	}

	fileName := GetFileName(filePath)

	fileName = fileName[:strings.LastIndex(fileName, ".")]

	return fileName
}
