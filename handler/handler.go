package handler

import (
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"net/http"
	"pica.go/api"
	"strings"
)

func LoginCheck(ctx *gin.Context) {
	username, password := ctx.Request.PostFormValue("username"), ctx.Request.PostFormValue("password")
	result := api.Login(username, password)

	if result["code"] == 200 {
		// store token to the session
		ctx.SetCookie("token", result["data"].(*simplejson.Json).MustString(),
			3*24*60*60, /*seconds*/
			"/",
			"localhost",
			false,
			true,
		)
	}

	// 重定向到首页
	ctx.Redirect(http.StatusMovedPermanently, "/pica/index")
}

func ShowCategory(ctx *gin.Context) {
	token, _ := ctx.Request.Cookie("token")
	result := api.Categories(token.Value)
	if result["code"] == 200 {
		categories := result["data"].(*simplejson.Json).MustArray()
		ctx.HTML(http.StatusOK, "category.html", categories)
	}
}

func GetImage(ctx *gin.Context) {
	token, _ := ctx.Request.Cookie("token")
	fileServer, path := ctx.Query("fileServer"), ctx.Query("path")

	if !strings.Contains(fileServer, "static") {
		image := api.ComicImage(token.Value, fileServer, path)
		if image != nil {
			ctx.DataFromReader(http.StatusOK, image.ContentLength, "image/png", image.Body, map[string]string{})
		}
	}
}

func GetComics(ctx *gin.Context) {
	token, _ := ctx.Request.Cookie("token")
	cat, page := ctx.Query("category"), ctx.Query("page")
	if page == "" {
		page = "1"
	}
	result := api.SearchByKeywordAndCategory(token.Value, "", []string{cat}, page, api.Newest)
	if result["code"] == 200 {
		comics := result["data"].(*simplejson.Json).MustMap()["docs"]
		ctx.HTML(http.StatusOK, "comics.html", comics)
	}
}

func GetComic(ctx *gin.Context) {
	token, _ := ctx.Request.Cookie("token")
	comicId, page := ctx.Param("comicId"), ctx.Query("page")
	if page == "" {
		page = "1"
	}
	comicInfo := api.ComicInfo(token.Value, comicId)
	epInfo := api.EpisodeInfo(token.Value, comicId, page)

	comic := map[string]interface{}{}

	comic["episodes"] = epInfo["data"].(*simplejson.Json).MustMap()["docs"]
	comic["info"] = comicInfo["data"].(*simplejson.Json).MustMap()

	ctx.HTML(http.StatusOK, "comic.html", comic)
}

func ReadComic(ctx *gin.Context) {
	token, _ := ctx.Request.Cookie("token")
	comicId, order, page := ctx.Param("comicId"), ctx.Param("order"), ctx.Query("page")

	if page == "" {
		page = "1"
	}

	images := api.EpisodeDetail(token.Value, comicId, order, page)

	pages := images["data"].(*simplejson.Json).MustMap()["pages"].(map[string]interface{})["docs"]

	ep := images["data"].(*simplejson.Json).MustMap()["ep"].(map[string]interface{})["title"]

	ctx.HTML(http.StatusOK, "episode.html", map[string]interface{}{"pages": pages, "ep": ep})
}
