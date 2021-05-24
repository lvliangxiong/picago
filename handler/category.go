package handler

import (
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"net/http"
	"pica.go/api"
	"pica.go/utils"
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
