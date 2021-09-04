package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lvliangxiong/picago/api"
	"github.com/lvliangxiong/picago/conf"
	"net/http"
)

func LoginCheck(ctx *gin.Context) {
	username, password := ctx.Request.PostFormValue("username"), ctx.Request.PostFormValue("password")

	code, _, token := api.Login(username, password)

	if code != 200 {
		ctx.Redirect(http.StatusMovedPermanently, "/pica/index")
		return
	}

	// store token to the cookie
	ctx.SetCookie(
		"token", token,
		3*24*60*60, /*seconds*/
		"/",
		"localhost",
		false,
		true,
	)
	conf.PublicToken = token

	ctx.Redirect(http.StatusMovedPermanently, "/pica/index")
}
