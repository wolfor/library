package FileTransferHelper

import (
	"errors"
	"strings"

	"github.com/shenshouer/ftp4go"
)

type FTPHelper struct {
	//用户名
	UserName string
	//密码
	Password string
	//远端IP
	RemoteIP string
	//远端端口
	RemortePort int
}

func NewFTPHelper(userName string, password string, remoteIP string, remotePort int) *FTPHelper {
	helper := new(FTPHelper)
	helper.UserName = userName
	helper.Password = password
	helper.RemoteIP = remoteIP
	helper.RemortePort = remotePort

	return helper
}

func (this *FTPHelper) connect() (*ftp4go.FTP, error) {
	var err error

	ftpClient := ftp4go.NewFTP(0) // 1 for debugging

	defer func() {
		er := recover()

		if er != nil {
			ftpClient = nil

			err = errors.New("unknow error")
		}
	}()

	//connect
	_, err = ftpClient.Connect(this.RemoteIP, this.RemortePort, "")
	if err != nil {
		return nil, err
	}

	_, err = ftpClient.Login(this.UserName, this.Password, "")
	if err != nil {
		return nil, err
	}

	return ftpClient, err
}

//ftp服务端指定目录下文件列表
func (this *FTPHelper) List(remoteDir string) ([]string, error) {
	ftpClient, err := this.connect()

	if err != nil {
		return nil, err
	}

	defer ftpClient.Quit()

	return ftpClient.Dir(remoteDir)
}

//ftp服务端指定目录下文件列表
func (this *FTPHelper) ListFilter(remoteDir string, filters ...string) (resourceList []string) {
	ftpClient, err := this.connect()

	if err != nil {
		return nil
	}

	defer ftpClient.Quit()

	arr, err := ftpClient.Dir(remoteDir)

	if err != nil {
		return nil
	}

	for _, r := range arr {
		if len(filters) > 0 && this.contain(r, filters) {
			resourceList = append(resourceList, r)
		} else {
			resourceList = append(resourceList, r)
		}
	}

	return resourceList
}

func (this *FTPHelper) contain(filename string, filters []string) (isContain bool) {
	isContain = false

	for _, filter := range filters {
		if strings.Contains(filename, filter) {
			isContain = true
			break
		}
	}

	return isContain
}

//推送文件至ftp服务端
func (this *FTPHelper) PushFile(localFilePath string, remoteDir string) error {
	ftpClient, err := this.connect()

	if err != nil {
		return err
	}

	defer ftpClient.Quit()

	err = ftpClient.UploadFile(remoteDir, localFilePath, false, nil)

	return err
}

//拉拽文件至ftp客户端
func (this *FTPHelper) PullFile(remoteFilePath string, localDir string) error {
	defer func() {

	}()

	ftpClient, err := this.connect()

	if err != nil {
		return err
	}

	defer ftpClient.Quit()

	err = ftpClient.DownloadFile(remoteFilePath, localDir, false)

	return err
}
