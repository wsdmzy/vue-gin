package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ziogie.top/gin/response"
)
//错误处理
func RecoveryMiddleware() gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Fail(ctx,nil, fmt.Sprint(err))
			}
		}()
		ctx.Next()
	}
}
