package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lvliangxiong/picago/utils"
	"net/http"
	"strings"
)

func ValidateToken(ctx *gin.Context) {
	// let login request and static request go
	if ctx.Request.URL.Path == "/pica/loginCheck" || strings.HasPrefix(ctx.Request.URL.Path, "/pica/static") {
		ctx.Next()
		return
	}

	// Check token
	_, err := utils.GetToken(ctx)

	// without token, you should login
	if err != nil {
		ctx.HTML(http.StatusOK, "login.html", nil)
		ctx.Abort()
		return
	}

	// with token, you should go
	ctx.Next()
}
