package middleware

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"time"

	"gitlab.ftsview.com/aircraft/ark-application-service/internal/constants"

	"gitlab.ftsview.com/fotoable-go/glog"
	"gitlab.ftsview.com/fotoable-go/gutil"

	"github.com/gin-gonic/gin"
)

//responseBodyWriter 拦截返回的Body
type responseBodyWriter struct {
	gin.ResponseWriter
	//body *bytes.Buffer
	ctx context.Context
}

//Write 缓存返回数据
func (r responseBodyWriter) Write(b []byte) (int, error) {
	//r.body.Write(b)
	glog.C(r.ctx).Infof("Response Data: %s", string(b))
	return r.ResponseWriter.Write(b)
}

// RequestID 获取/添加请求ID中间件
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取请求的数据
		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		body, _ := ioutil.ReadAll(tee)
		c.Request.Body = ioutil.NopCloser(&buf)

		//生成或者设置请求ID
		st := time.Now()
		uuID := c.GetHeader(constants.HeaderXRequestID)
		if len(uuID) == 0 {
			uuID = gutil.UUID()
		}
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), glog.TrackKey, uuID))
		c.Header(constants.HeaderXRequestID, uuID)
		ctx := c.Request.Context()
		//请求数据
		glog.C(ctx).Infof("Request Data: %s", string(body))
		w := &responseBodyWriter{ResponseWriter: c.Writer, ctx: ctx}
		c.Writer = w
		c.Next()
		glog.C(ctx).Infof("Method: %s, URL: %s, Status: %d, Latency: %dms", c.Request.Method, c.Request.URL, c.Writer.Status(), time.Since(st).Milliseconds())
	}
}
