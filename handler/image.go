package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lvliangxiong/picago/api"
	"github.com/lvliangxiong/picago/utils"
	"net/http"
)

func GetImage(ctx *gin.Context) {
	token, err := utils.GetToken(ctx)

	if err != nil {
		return
	}

	fileServer, path := ctx.Query("fileServer"), ctx.Query("path")
	imageResp := api.GetImage(token, fileServer, path)

	if imageResp == nil || imageResp.StatusCode != 200 {
		return
	}

	ctx.DataFromReader(http.StatusOK, imageResp.ContentLength, "image/png", imageResp.Body, map[string]string{})
}
