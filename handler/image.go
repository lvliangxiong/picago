package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lvliangxiong/pica.go/api"
	"github.com/lvliangxiong/pica.go/utils"
	"net/http"
)

func GetImage(ctx *gin.Context) {
	token, err := utils.GetToken(ctx)
	if err == nil {
		fileServer, path := ctx.Query("fileServer"), ctx.Query("path")
		image := api.ComicImage(token, fileServer, path)
		if image != nil {
			ctx.DataFromReader(http.StatusOK, image.ContentLength, "image/png", image.Body, map[string]string{})
		}
	}
}
