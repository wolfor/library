package FileTransferHelper

import (
	"github.com/dutchcoders/goftp"
)

type GoFTPHelper struct {
	//用户名
	UserName string
	//密码
	Password   string
	FtpAddress string
}

func NewGoFTPHelper(ftpAddress string, userName string, password string) *GoFTPHelper {
	helper := new(GoFTPHelper)

	helper.FtpAddress = ftpAddress

	helper.UserName = userName
	helper.Password = password

	return helper
}

func (this *GoFTPHelper) connect() *goftp.FTP {
	client, err := goftp.Connect(this.FtpAddress)

	if err != nil {
		panic(err)
	}
	defer client.Close()

	err = client.Login(this.UserName, this.Password)

	if err != nil {
		panic(err)
	}

	return client
}

func (this *GoFTPHelper) List(remoteDir string) ([]string, error) {
	client := this.connect()

	defer client.Close()

	files, err := client.List(remoteDir)

	return files, err
}

func (this *GoFTPHelper) close(client *goftp.FTP) {
	if client != nil {
		client.Quit()
		client.Close()
	}
}
