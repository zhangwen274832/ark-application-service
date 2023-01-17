package config

import (
	"context"

	"gitlab.ftsview.com/OpenPlatform/open-go-lib/amango"
	"gitlab.ftsview.com/fotoable-go/glog"
	"gitlab.ftsview.com/fotoable-go/gmysql"
)

type Config struct {
	Logger LoggerConf    `ftV:"config.logger"` //普通日志配置
	Redis  RedisConf     `ftV:"config.redis"`  //redis 配置
	Cache  Cache         `ftV:"config.cache"`  //本次缓存配置
	Mysql  gmysql.Config `ftV:"config.mysql"`  //mysql 配置
}

type LoggerConf struct {
	StdOut   bool   `ftV:"stdout"`
	Level    string `ftV:"level"`
	Path     string `ftV:"path"`
	FileName string `ftV:"file_name"`
}

type RedisConf struct {
	UserName string   `ftV:"user_name"`
	Password string   `ftV:"password"`
	Host     []string `ftV:"host"`
	PoolSize int      `ftV:"poolsize"`
	MinIdle  int      `ftV:"minidle"`
}
type Cache struct {
	RedisInterval int `ftV:"redisinterval"`
}

var GlobConfig = &Config{}
var ImageVersion string

func MustInit(projectID, envID, addr, version string) {
	ImageVersion = version
	todo := context.Background()
	//初始化客户端
	client, err := amango.NewConfigClient(todo, &amango.ClientOptions{
		ProjectID: projectID, //项目ID
		EvnID:     envID,     //环境ID
		Host:      addr,      //配置中心地址
	}, GlobConfig)
	if err != nil {
		glog.Errorf(todo, "amgngo NewConfigClient error: %v", err)
		panic(err)
	}
	//拉去远端配置
	if err = client.ListenConfig(); err != nil {
		glog.Errorf(todo, "init config error: %v", err)
		panic(err)
	}
	//注册监听变化方法
	amango.ListenChangeEvent(todo, configChange)
}

// configChange 配置变化回调函数
func configChange(m map[string]*amango.Row) {
	// 处理日志级别变动
	if row, ok := m["config.logger.level"]; ok {
		if newLevel, ok := row.NewValue.(string); ok {
			glog.ChangeFileStdoutLevel(glog.Level(newLevel))
		}
	}
}
