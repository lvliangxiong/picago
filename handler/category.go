package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lvliangxiong/picago/api"
	"github.com/lvliangxiong/picago/utils"
)

func ShowCategory(ctx *gin.Context) {
	token, err := utils.GetToken(ctx)
	if err != nil {
		return
	}
	code, _, categories := api.FetchCategories(token)
	if code != 200 {
		return
	}
	ctx.HTML(http.StatusOK, "category.html", categories)
}
