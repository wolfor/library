package FileTransferHelper

import (
	"errors"
<<<<<<< HEAD
	"log"
=======
>>>>>>> 305541849dce4fc3a17980b5e29e48ca5323a688
	//	"time"
	//	"log"
	"strings"
	"unicode"

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
	//连接超时时长(毫秒)
	ConnTimeout int64
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

	ftpClient := ftp4go.NewFTP2(0, this.ConnTimeout) // 1 for debugging

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

func (this *FTPHelper) Connect() *ftp4go.FTP {
	client, err := this.connect()

	if err != nil {
		return nil
	}

	return client
}

func (this *FTPHelper) isFile(resourceType string) bool {
	if resourceType == "1" {
		return true
	} else {
		return false
	}
}

//根据远端资源信息获取资源类型
func (this *FTPHelper) getResourceType(resourceInfo string) string {
	ri := strings.FieldsFunc(resourceInfo, unicode.IsSpace)
	resourceType := ri[1]

	return resourceType
}

//根据远端资源信息获取资源名
func (this *FTPHelper) getResourceName(resourceInfo string) string {
	ri := strings.FieldsFunc(resourceInfo, unicode.IsSpace)
	resourceName := ri[8]

	return resourceName
}

func (this *FTPHelper) ListDir(remoteDir string) []string {
	resourceList, err := this.List(remoteDir)

	if err != nil {
		return nil
	}

	folderList := make([]string, 0)

	for _, resourceInfo := range resourceList {

		resourceType := this.getResourceType(resourceInfo)

		if this.isFile(resourceType) {
			continue
		}

		resourceName := this.getResourceName(resourceInfo)

		folderList = append(folderList, resourceName)
	}

	return folderList
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

//get file name array
func (this *FTPHelper) ListFile(remoteDir string, filters ...string) (resourceList []string) {
	fileInfoList := this.ListFilter(remoteDir, filters...)

	//	log.Println("ftpHelper.ListFile fileInfoList count=", len(fileInfoList))

	if fileInfoList == nil || len(fileInfoList) <= 0 {
		return nil
	}

	for _, fileInfo := range fileInfoList {
		fileName := this.getResourceName(fileInfo)

		resourceList = append(resourceList, fileName)
	}

	return resourceList
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

	//	log.Println("ftpHelper.ListFilter arr count=", len(arr))

	for _, r := range arr {
		if len(filters) > 0 {
			if this.contain(r, filters) {
				resourceList = append(resourceList, r)
			}
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
		//		log.Println("ftp connect error=", err)
		return err
	}

	defer ftpClient.Quit()

	err = ftpClient.UploadFile(remoteDir, localFilePath, false, nil)

	if err != nil {
		log.Println("ftp push file error=", err)
	}

	return err
}

//拉拽文件至ftp客户端
func (this *FTPHelper) PullFile(remoteFilePath string, localDir string) error {

	ftpClient, err := this.connect()

	if err != nil {
		return err
	}

	defer ftpClient.Quit()

	err = ftpClient.DownloadFile(remoteFilePath, localDir, false)

	return err
}

//拉拽文件至ftp客户端,duan dian xu chuan
func (this *FTPHelper) PullResumeFile(remoteFilePath string, localDir string) error {

	ftpClient, err := this.connect()

	if err != nil {
		return err
	}

	defer ftpClient.Quit()

	err = ftpClient.DownloadResumeFile(remoteFilePath, localDir, false)

	return err
}

//download ftp file
func (this *FTPHelper) DownFile(ftpClient *ftp4go.FTP, remoteFilePath string, localDir string) error {

	err := ftpClient.DownloadResumeFile(remoteFilePath, localDir, false)

	return err
}

//create remote directory
func (this *FTPHelper) Mkd(remoteTargetPath string) (string, error) {
	ftpClient, err := this.connect()

	if err != nil {
		return "", err
	}
	defer ftpClient.Quit()

	dname, err := ftpClient.Mkd(remoteTargetPath)

	return dname, err
}

//rename or move file
func (this *FTPHelper) Rename(fromName, toName string) (*ftp4go.Response, error) {
	ftpClient, err := this.connect()

	if err != nil {
		return nil, err
	}
	defer ftpClient.Quit()

	response, err := ftpClient.Rename(fromName, toName)

	return response, err
}
