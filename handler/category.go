package handler

import (
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"github.com/lvliangxiong/pica.go/api"
	"github.com/lvliangxiong/pica.go/utils"
	"net/http"
)

func ShowCategory(ctx *gin.Context) {
	token, err := utils.GetToken(ctx)
	if err == nil {
		result := api.Categories(token)
		if result["code"] == 200 {
			categories := result["data"].(*simplejson.Json).MustArray()
			ctx.HTML(http.StatusOK, "category.html", categories)
		}
	}
}
