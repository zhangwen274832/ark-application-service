package cmd

import (
	"context"
	"os"

	"gitlab.ftsview.com/aircraft/ark-application-service/internal"
	"gitlab.ftsview.com/aircraft/ark-application-service/internal/cache"
	"gitlab.ftsview.com/aircraft/ark-application-service/internal/config"
	"gitlab.ftsview.com/aircraft/ark-application-service/internal/constants"

	"gitlab.ftsview.com/fotoable-go/glog"

	"github.com/spf13/cobra"
)

var (
	// aman flag
	amanAddr      string
	amanEnvID     string
	amanProjectID string
	version       string

	rootCmd = cobra.Command{
		Use: "ark-application-service",
		Run: func(cmd *cobra.Command, args []string) {
			internal.Run()
		},
	}
)

func init() {
	cobra.OnInitialize(initDependency)
	rootCmd.PersistentFlags().StringVar(&amanAddr, "aman_addr", "https://aman-internal.akgoo.net", "aman 的请求地址")
	rootCmd.PersistentFlags().StringVar(&amanProjectID, "aman_project_id", "ark-application-service", "在 aman 配置的项目ID")
	rootCmd.PersistentFlags().StringVar(&amanEnvID, "aman_env_id", "1", "在 aman 配置的环境ID")
	rootCmd.PersistentFlags().StringVar(&version, "version", "latest.", "service version")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func initDependency() {
	// 初始化依赖，如，配置、数据库、缓存、日志等
	// 初始化配置，必须前置
	initConfig()
	loggerInit()
	redisInit()
	mysqlInit()
	cacheInit()
}

// nitConfig 初始化远程配置
func initConfig() {
	config.MustInit(amanProjectID, amanEnvID, amanAddr, version)
}

func loggerInit() {
	hostname, _ := os.Hostname()
	logFileName := config.GlobConfig.Logger.Path + hostname +
		constants.MiddleLine + config.GlobConfig.Logger.FileName
	logOptions := []glog.Option{
		glog.WithFileLocation(logFileName),
		glog.WithLevel(glog.Level(config.GlobConfig.Logger.Level)),
		glog.WithCustomizedGlobalField(map[string]interface{}{constants.LoggerServerCode: constants.ServiceCode}),
	}
	if config.GlobConfig.Logger.StdOut {
		logOptions = append(logOptions, glog.WithConsoleStdout())
	}
	if err := glog.Init(logOptions...); err != nil {
		glog.Errorf(context.TODO(), "init logger err: %v", err)
	}
}

func cacheInit() {
	cache.MustInit(config.GlobConfig.Cache.RedisInterval)
}

func mysqlInit() {
	//注意：使用时需要安装 go get gitlab.ftsview.com/fotoable-go/gmysql
	//且，生成了CRUD文件
	//gmysql.MustInitMysql(&config.GlobConfig.Mysql)
}

func redisInit() {
	//注意：使用时需要安装 go get gitlab.ftsview.com/fotoable-go/gredis
	//初始化Redis，使用不同的名称创建不同的Redis实例，Redis实例类型包括（哨兵、集群）
	//gredis.MustInit(constants.RedisClusterName,
	//	gredis.WithCluster(),
	//	gredis.WithHosts(config.GlobConfig.Redis.Host),
	//	gredis.WithPoolSize(config.GlobConfig.Redis.PoolSize),
	//	gredis.WithPassWord(config.GlobConfig.Redis.Password),
	//	gredis.WithMinIdleCons(config.GlobConfig.Redis.MinIdle))
}
