package handler

import (
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"net/http"
	"pica.go/api"
	"pica.go/conf"
)

func LoginCheck(ctx *gin.Context) {
	username, password := ctx.Request.PostFormValue("username"), ctx.Request.PostFormValue("password")
	result := api.Login(username, password)

	if result["code"] == 200 {
		// store token to the cookie
		token := result["data"].(*simplejson.Json).MustString()
		ctx.SetCookie("token", token,
			3*24*60*60, /*seconds*/
			"/",
			"localhost",
			false,
			true,
		)
		conf.PublicToken = token
	}

	// 重定向到首页
	ctx.Redirect(http.StatusMovedPermanently, "/pica/index")
}
