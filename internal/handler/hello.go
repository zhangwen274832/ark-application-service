package handler

import (
	"sync"

	"gitlab.ftsview.com/aircraft/ark-application-service/internal/handler/bean"
	"gitlab.ftsview.com/aircraft/ark-application-service/internal/logic"
	gginutil "gitlab.ftsview.com/fotoable-go/ggin-util"

	"github.com/gin-gonic/gin"
)

var (
	_helloOnce    sync.Once
	_helloHandler *HelloHandler
)

type HelloHandler struct {
	gginutil.BaseHandler
	helloLogic *logic.HelloLogic
}

func SingletonHelloHandler() *HelloHandler {
	_helloOnce.Do(func() {
		_helloHandler = &HelloHandler{
			helloLogic: logic.SingletonHelloLogic(),
		}
	})
	return _helloHandler
}

func (h *HelloHandler) Hello(c *gin.Context) {
	ctx := c.Request.Context()
	req := &bean.HelloReq{}
	//true：绑定Header，false：不绑定
	if !h.Bind(c, req) {
		return
	}
	resp, err := h.helloLogic.Hello(ctx, req.ID)
	if err != nil {
		h.Fail(c, err)
		return
	}
	h.Success(c, resp)
}
