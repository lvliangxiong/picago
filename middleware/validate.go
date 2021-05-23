package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidateToken(ctx *gin.Context) {
	if ctx.Request.URL.Path != "/pica/loginCheck" {
		_, err := ctx.Cookie("token")

		if err != nil {
			// 没有 token 就去登录
			ctx.HTML(http.StatusOK, "login.html", nil)
			ctx.Abort()
			return
		}
	}
	ctx.Next()
}
