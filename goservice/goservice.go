// goservice project goservice.go
package goservice

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//初始化服务
type InitService func() error

//开始服务
type StartService func() error

//停止服务
type StopService func() error

type GoService struct {
	initFunc  InitService
	startFunc StartService
	stopFunc  StopService
}

func NewGoService(initFunc InitService, startFunc StartService, stopFunc StopService) *GoService {
	service := new(GoService)
	service.initFunc = initFunc
	service.startFunc = startFunc
	service.stopFunc = stopFunc

	return service
}

func (this *GoService) Run() {
	if this.initFunc != nil {
		this.initFunc()
	}

	//创建监听退出chan
	c := make(chan os.Signal)

	//监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)

	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				this.exitFunc()
			case syscall.SIGUSR1:
				fmt.Println("usr1", s)
			case syscall.SIGUSR2:
				fmt.Println("usr2", s)
			default:
				fmt.Println("other", s)
			}
		}
	}()

	if this.startFunc != nil {
		this.startFunc()
	}

	for {
		time.Sleep(50 * time.Millisecond)
	}
}

func (this *GoService) exitFunc() {

	//自定义退出
	if this.stopFunc != nil {
		this.stopFunc()
	}

	os.Exit(0)
}
