package api

import (
	"context"
	"net/http"

	"gitlab.ftsview.com/aircraft/ark-application-service/internal/config"
	"gitlab.ftsview.com/aircraft/ark-application-service/internal/handler"

	gginutil "gitlab.ftsview.com/fotoable-go/ggin-util"
	"gitlab.ftsview.com/fotoable-go/ginprom"
	"gitlab.ftsview.com/fotoable-go/glog"

	"github.com/gin-gonic/gin"
)

//StartServerGin 启动Gin服务
func StartServerGin() {
	addr := ":40200"
	glog.Infof(context.TODO(), "server port : %s", addr)
	router := generateRouter()
	if err := router.Run(addr); err != nil {
		glog.Errorf(context.TODO(), "StartServerGin err: %v", err)
		panic(err)
	}
}

func generateRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	//初始化 gin Engine
	router := gin.New()
	// 监控接口
	router.GET("/metrics", ginprom.PromHandler())
	router.GET("/version", func(context *gin.Context) {
		context.JSON(http.StatusOK, map[string]interface{}{"version": config.ImageVersion})
	})
	router.GET("/heartbeat", func(context *gin.Context) {})
	router.StaticFS("/static", http.Dir("./static"))

	//设置中间件
	router.Use(
		gginutil.Recovery(),
		gginutil.RequestID(gginutil.AllPrint),
		gginutil.SignToC(false), //默认True，True:标识全局开启加解密，False:不开启， 且不能与EncryptionV2共用
		ginprom.PromMiddleware(
			ginprom.WithExcludeRegexEndpoint(`^/metrics|/version|/heartbeat$`),
		),
	)

	//********** Hello 接口为业务接口示例，可删除 ***********
	router.Any("/hello", handler.SingletonHelloHandler().Hello)

	return router
}
