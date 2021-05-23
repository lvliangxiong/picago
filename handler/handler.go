package handler

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"net/http"
	"pica.go/api"
	"strconv"
	"strings"
)

func LoginCheck(ctx *gin.Context) {
	username, password := ctx.Request.PostFormValue("username"), ctx.Request.PostFormValue("password")
	result := api.Login(username, password)

	if result["code"] == 200 {
		// store token to the cookie
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
	comicId, order := ctx.Param("comicId"), ctx.Param("order")

	// Fetch the first page of images
	page1 := api.EpisodeDetail(token.Value, comicId, order, "1")

	ep := page1["data"].(*simplejson.Json).MustMap()["ep"].(map[string]interface{})["title"] // episode title
	pages := page1["data"].(*simplejson.Json).MustMap()["pages"].(map[string]interface{})    // images of page 1

	imagesInPage1 := pages["docs"].([]interface{}) // A variable to store all images in this episode

	pageCount, _ := pages["pages"].(json.Number).Int64()  // Total page count
	imageCount, _ := pages["total"].(json.Number).Int64() // Total page count

	images := make([]interface{}, 0, imageCount)

	images = addImages(images, imagesInPage1)

	if pageCount > 1 {
		// if there are more than one page, fetch other pages of images
		for i := 2; i <= int(pageCount); i++ {
			pageI := api.EpisodeDetail(token.Value, comicId, order, strconv.Itoa(i))
			imagesInPageI := pageI["data"].(*simplejson.Json).MustMap()["pages"].(map[string]interface{})["docs"].([]interface{})
			images = addImages(images, imagesInPageI)
		}
	}

	ctx.HTML(http.StatusOK, "episode.html", map[string]interface{}{"ep": ep, "images": images})
}

func addImages(images []interface{}, imagesN []interface{}) []interface{} {
	for _, img := range imagesN {
		images = append(images, img)
	}
	return images
}
