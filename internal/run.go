package internal

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"gitlab.ftsview.com/aircraft/ark-application-service/internal/api"
)

func Run() {
	go api.StartServerGin()
	exit()
}

func exit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			//终端主动退出。（Ctrl+C）、（Ctrl+/）、（KILL + PID）
			time.Sleep(time.Second * 1)
			return
		case syscall.SIGHUP:
			//终端控制进程结束（终端连接断开）
			time.Sleep(time.Second * 1)
			return
		}
	}
}
