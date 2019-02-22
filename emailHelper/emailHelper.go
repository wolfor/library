// emailHelper project emailHelper.go
package emailHelper

import (
	"strconv"

	"github.com/go-gomail/gomail"
)

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
	"strings"
)

//邮件发送
//parament:
//mailFrom value example
/*
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mailFrom := map[string]string{
		"user": "XXXXXXXX@qq.com",
		"pass": "XXXXXXX",
		"host": "smtp.qq.com",
		"port": "465",
	}
*/
func SendMail(aliasFrom string, mailFrom map[string]string, mailTo []string, subject string, body string) error {
	port, _ := strconv.Atoi(mailFrom["port"]) //转换端口类型为int

	m := gomail.NewMessage()
	m.SetHeader("From", aliasFrom+"<"+mailFrom["user"]+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo...)                            //发送给多个用户
	m.SetHeader("Subject", subject)                         //设置邮件主题
	m.SetBody("text/html", body)                            //设置邮件正文

	d := gomail.NewDialer(mailFrom["host"], port, mailFrom["user"], mailFrom["pass"])

	err := d.DialAndSend(m)
	return err
}

//////////////////////////////////////////////////////

//Client simple email client support ssl
type Client struct {
	user     string
	addr     string
	nickName string
	isSSL    bool
	auth     smtp.Auth
}

//New new email client
func New(user, password, nickName, host string, port int, isSsl bool) *Client {
	ec := &Client{
		user:  user,
		addr:  fmt.Sprintf("%s:%d", host, port),
		isSSL: isSsl,
		auth:  smtp.PlainAuth("", user, password, host),
	}
	if nickName == "" {
		ec.nickName = user
	} else {
		ec.nickName = nickName
	}
	return ec
}

func (ec *Client) generateEmailMsg(toUser []string, subject, content string) []byte {
	return ec.generateEmailMsgByte(toUser, subject, []byte(content))
}

func (ec *Client) generateEmailMsgByte(toUser []string, subject string, body []byte) []byte {
	msgStr := fmt.Sprintf("To: %s\r\nFrom: %s<%s>\r\nSubject: %s\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\n",
		strings.Join(toUser, ","), ec.nickName, ec.user, subject)
	return append([]byte(msgStr), body...)
}

func (ec *Client) sendMailTLS(toUser []string, msg []byte) error {
	host, _, _ := net.SplitHostPort(ec.addr)
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}
	conn, err := tls.Dial("tcp", ec.addr, tlsconfig)
	if err != nil {
		return fmt.Errorf("DialConn:%v", err)
	}
	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return fmt.Errorf("Client:generateClient:%v", err)
	}
	defer client.Close()
	if ec.auth != nil {
		if ok, _ := client.Extension("AUTH"); ok {
			if err = client.Auth(ec.auth); err != nil {
				return fmt.Errorf("Client:clientAuth:%v", err)
			}
		}
	}
	if err = client.Mail(ec.user); err != nil {
		return fmt.Errorf("Client:clientMail:%v", err)
	}

	for _, addr := range toUser {
		if err = client.Rcpt(addr); err != nil {
			return fmt.Errorf("Client:Rcpt:%v", err)
		}
	}
	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("Client:%v", err)
	}
	_, err = w.Write(msg)
	if err != nil {
		return fmt.Errorf("Client:WriterBody:%v", err)
	}
	err = w.Close()
	if err != nil {
		return fmt.Errorf("Client:CloseBody:%v", err)
	}
	return client.Quit()
}

func (ec *Client) sendMail(toUser []string, msg []byte) error {
	return smtp.SendMail(ec.addr, ec.auth, ec.user, toUser, msg)
}

//SendEmail send email by string content
func (ec *Client) SendEmail(toUser []string, subject string, content string) error {
	msg := ec.generateEmailMsg(toUser, subject, content)
	if ec.isSSL {
		return ec.sendMailTLS(toUser, msg)
	}
	return ec.sendMail(toUser, msg)
}

//SendEmailByte send email by byte body
func (ec *Client) SendEmailByte(toUser []string, subject string, body []byte) error {
	msg := ec.generateEmailMsgByte(toUser, subject, body)
	if ec.isSSL {
		return ec.sendMailTLS(toUser, msg)
	}
	return ec.sendMail(toUser, msg)
}
