package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pica.go/utils"
)

func ValidateToken(ctx *gin.Context) {
	if ctx.Request.URL.Path != "/pica/loginCheck" {
		_, err := utils.GetToken(ctx)

		// 没有 token 就去登录
		if err != nil {
			ctx.HTML(http.StatusOK, "login.html", nil)
			ctx.Abort()
			return
		}
	}
	ctx.Next()
}
