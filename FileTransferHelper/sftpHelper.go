package FileTransferHelper

import (
	"log"
	"net"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/sftp"

	"golang.org/x/crypto/ssh"
)

type SFTPHelper struct {
	//用户名
	UserName string
	//密码
	Password string
	//远端IP
	RemoteIP string
	//远端端口
	RemortePort int
}

func NewSFTPHelper(userName string, password string, remoteIP string, remotePort int) *SFTPHelper {
	helper := new(SFTPHelper)
	helper.UserName = userName
	helper.Password = password
	helper.RemoteIP = remoteIP
	helper.RemortePort = remotePort

	return helper
}

func (this *SFTPHelper) connect() (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
		port         string
	)

	port = strconv.Itoa(this.RemortePort)

	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(this.Password))

	clientConfig = &ssh.ClientConfig{
		User:    this.UserName,
		Auth:    auth,
		Timeout: 30 * time.Second,
		//需要验证服务端，不做验证返回nil就可以，点击HostKeyCallback看源码就知道了
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = strings.Join([]string{this.RemoteIP, ":", port}, "") // fmt.Sprintf("%s:%d", host, port)

	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}

	return sftpClient, nil
}

//sftp服务端指定目录下文件列表
func (this *SFTPHelper) List(remoteDir string) ([]os.FileInfo, error) {
	sftpClient, err := this.connect()
	if err != nil {
		return nil, err
	}
	defer sftpClient.Close()

	return sftpClient.ReadDir(remoteDir)
}

//推送文件至sftp服务端
func (this *SFTPHelper) PushFile(localFilePath string, remoteDir string) {
	var (
		err        error
		sftpClient *sftp.Client
	)

	// 这里换成实际的 SSH 连接的 用户名，密码，主机名或IP，SSH端口
	sftpClient, err = this.connect()
	if err != nil {
		return
	}
	defer sftpClient.Close()

	// 用来测试的本地文件路径 和 远程机器上的文件夹
	srcFile, err := os.Open(localFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	var remoteFileName = path.Base(localFilePath)
	dstFile, err := sftpClient.Create(path.Join(remoteDir, remoteFileName))
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	srcfi, _ := srcFile.Stat()

	var (
		scrLen, quotient, endLen, cycIndex int64
	)

	scrLen = srcfi.Size()

	quotient = scrLen / 1024

	endLen = scrLen - quotient*1024

	cycIndex = 0
	var buf []byte

	for {
		if quotient > 0 {

			if endLen > 0 && quotient == cycIndex {
				buf = make([]byte, endLen)
			} else {
				buf = make([]byte, 1024)
			}

			cycIndex++

		} else {
			buf = make([]byte, scrLen)
		}

		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		dstFile.Write(buf)
	}
}

//拉拽文件至sftp客户端
func (this *SFTPHelper) PullFile(remoteFilePath string, localDir string) error {
	var (
		err        error
		sftpClient *sftp.Client
	)

	// 这里换成实际的 SSH 连接的 用户名，密码，主机名或IP，SSH端口
	sftpClient, err = this.connect()
	if err != nil {
		return err
	}
	defer sftpClient.Close()

	// 用来测试的远程文件路径 和 本地文件夹
	srcFile, err := sftpClient.Open(remoteFilePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	var localFileName = path.Base(remoteFilePath)
	dstFile, err := os.Create(path.Join(localDir, localFileName))
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err = srcFile.WriteTo(dstFile); err != nil {
		return err
	}

	return nil
}

//删除sftp服务端文件
func (this *SFTPHelper) Remove(remoteFilePath string) {
	var (
		err        error
		sftpClient *sftp.Client
	)

	// 这里换成实际的 SSH 连接的 用户名，密码，主机名或IP，SSH端口
	sftpClient, err = this.connect()
	if err != nil {
		log.Fatal(err)
	}
	defer sftpClient.Close()

	err = sftpClient.Remove(remoteFilePath)

	if err != nil {
		log.Fatal(err)
	}
}
