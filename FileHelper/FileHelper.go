// FileHelper project FileHelper.go
package FileHelper

import (
	"fmt"
	"io/ioutil"
	"library/Path"
<<<<<<< HEAD
	"log"
=======
>>>>>>> 305541849dce4fc3a17980b5e29e48ca5323a688
	"os"
)

//写文件，若文件已存在，则覆盖
func WriteFileCover(filePath string, content string) {
	var d1 = []byte(content)

	ioutil.WriteFile(filePath, d1, 0666) //写入文件(字节数
}

//写文件
func WriteFile(filePath, fileContent string) bool {
	var (
		fi  *os.File
		err error
	)

	if Path.IsExists(filePath) {
		//use "os.OpenFile(filePath)" code throw error message "bad file descriptor"
		fi, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0666)
	} else {
		fi, err = os.Create(filePath)
	}

	if err != nil {
<<<<<<< HEAD
		log.Println("open or create file failure ,error =", err)
=======
>>>>>>> 305541849dce4fc3a17980b5e29e48ca5323a688
		return false
	}
	defer fi.Close()

	_, er := fi.WriteString(fileContent)

	if er != nil {
		fmt.Println("write string error=", er)
		return false
	}

	return true
}

//读文本文件
func ReadTextFile(filePath string) string {
	file, err := os.Open(filePath)

	if err != nil {
		return ""
	}

	defer file.Close()

	data, er := ioutil.ReadAll(file)

	if er != nil {
		return ""
	}

	return string(data)
}

//删除指定文件
func RemoveFile(localFilePath string) bool {
	err := os.Remove(localFilePath)

	if err != nil {
		return false
	}

	return true
}

//删除指定文件夹
func RemovePath(localPath string) bool {
	err := os.RemoveAll(localPath)

	if err != nil {
		return false
	}

	return true
}

//删除指定文件夹下所有文件
func RemoveAllFile(localDir string) {
<<<<<<< HEAD
	filePathArr := GetAllFile(localDir)
=======
	filePathArr := getAllFile(localDir)
>>>>>>> 305541849dce4fc3a17980b5e29e48ca5323a688

	for _, filePath := range filePathArr {
		RemoveFile(filePath)
	}
}

//第归获取目录下含子目录所有文件
<<<<<<< HEAD
func GetAllFile(currPath string) (filePathArrary []string) {
=======
func getAllFile(currPath string) (filePathArrary []string) {
>>>>>>> 305541849dce4fc3a17980b5e29e48ca5323a688
	rd, err := ioutil.ReadDir(currPath)

	if err != nil {
		return nil
	}

	for _, fi := range rd {
		if fi.IsDir() {
<<<<<<< HEAD
			filePathArrary = append(filePathArrary, GetAllFile(currPath+"/"+fi.Name())...)
=======
			filePathArrary = append(filePathArrary, getAllFile(currPath+"/"+fi.Name())...)
>>>>>>> 305541849dce4fc3a17980b5e29e48ca5323a688
		} else {
			filePathArrary = append(filePathArrary, currPath+"/"+fi.Name())
		}
	}

	return filePathArrary
}
