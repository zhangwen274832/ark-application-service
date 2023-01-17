package logic

import (
	"context"
	"sync"

	"gitlab.ftsview.com/aircraft/ark-application-service/internal/service"
)

var (
	_helloOnce  sync.Once
	_helloLogic *HelloLogic
)

type HelloLogic struct {
	helloService *service.HelloService
}

func SingletonHelloLogic() *HelloLogic {
	_helloOnce.Do(func() {
		_helloLogic = &HelloLogic{
			helloService: service.SingletonHelloService(),
		}
	})
	return _helloLogic
}

func (l *HelloLogic) Hello(ctx context.Context, id int64) (string, error) {
	data, err := l.helloService.Hello(ctx, id)
	if err != nil {
		return "", err
	}
	return data, nil
}
