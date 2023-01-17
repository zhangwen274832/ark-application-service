package service

import (
	"context"
	"sync"
)

var (
	_helloOnce    sync.Once
	_helloService *HelloService
)

type HelloService struct {
}

func SingletonHelloService() *HelloService {
	_helloOnce.Do(func() {
		_helloService = &HelloService{}
	})
	return _helloService
}

func (l *HelloService) Hello(ctx context.Context, id int64) (string, error) {
	//=====================MySql使用方法=====================
	//使用gorm-gen查询
	//获取数据库连接
	//db := gmysql.DB(ctx, constants.DBExperiment)
	//Table 为数据库中的表面
	//table := query.Use(db).Table
	//table.WithContext(ctx).Select(t.UserID, t.ExprID, t.VariantID).Where(t.IsDeleted.Eq(0)).Find()
	//更多操作参照文档：https://github.com/go-gorm/gen

	//=====================Redis使用方法=====================
	//tempRedis := gredis.Redis(constants.RedisClusterName)
	//tempRedis.HGetAll(ctx, "redis:hash:key")
	return "Hello World", nil
}
