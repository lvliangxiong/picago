package handler

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"net/http"
	"pica.go/api"
	"pica.go/utils"
	"strconv"
)

func GetComics(ctx *gin.Context) {
	token, err := utils.GetToken(ctx)
	if err == nil {
		cat, page := ctx.Query("category"), ctx.Query("page")
		if page == "" {
			page = "1"
		}
		result := api.SearchByKeywordAndCategory(token, "", []string{cat}, page, api.Newest)
		if result["code"] == 200 {
			comics := result["data"].(*simplejson.Json).MustMap()["docs"]
			ctx.HTML(http.StatusOK, "comics.html", comics)
		}
	}
}

// Show comic information including a list containing all episodes
func GetComic(ctx *gin.Context) {
	token, err := utils.GetToken(ctx)
	if err == nil {
		comicId := ctx.Param("comicId")

		comicInfo := api.ComicInfo(token, comicId)
		epInfoPage1 := api.EpisodeInfo(token, comicId, "1")

		pages := epInfoPage1["data"].(*simplejson.Json).MustMap()

		pagesCount, _ := pages["pages"].(json.Number).Int64()
		epsCount, _ := pages["total"].(json.Number).Int64()

		epsInPage1 := epInfoPage1["data"].(*simplejson.Json).MustMap()["docs"].([]interface{})

		eps := make([]interface{}, 0, epsCount)

		eps = utils.MergeSlices(eps, epsInPage1)

		if pagesCount > 1 {
			for i := 2; i <= int(pagesCount); i++ {
				epInfoInPageI := api.EpisodeInfo(token, comicId, strconv.Itoa(i))
				epsInPageI := epInfoInPageI["data"].(*simplejson.Json).MustMap()["docs"].([]interface{})
				eps = utils.MergeSlices(eps, epsInPageI)
			}
		}

		comic := map[string]interface{}{}

		comic["info"] = comicInfo["data"].(*simplejson.Json).MustMap()
		comic["episodes"] = eps

		ctx.HTML(http.StatusOK, "comic.html", comic)
	}
}

// Show all images in one episode
func ReadComic(ctx *gin.Context) {
	token, err := utils.GetToken(ctx)
	if err == nil {
		comicId, order := ctx.Param("comicId"), ctx.Param("order")

		// Fetch the first page of images
		page1 := api.EpisodeDetail(token, comicId, order, "1")

		ep := page1["data"].(*simplejson.Json).MustMap()["ep"].(map[string]interface{})["title"] // episode title
		pages := page1["data"].(*simplejson.Json).MustMap()["pages"].(map[string]interface{})    // images of page 1

		imagesInPage1 := pages["docs"].([]interface{}) // A variable to store all images in this episode

		pageCount, _ := pages["pages"].(json.Number).Int64()  // Total page count
		imageCount, _ := pages["total"].(json.Number).Int64() // Total images count

		images := make([]interface{}, 0, imageCount)

		images = utils.MergeSlices(images, imagesInPage1)

		if pageCount > 1 {
			// if there are more than one page, fetch other pages of images
			for i := 2; i <= int(pageCount); i++ {
				pageI := api.EpisodeDetail(token, comicId, order, strconv.Itoa(i))
				imagesInPageI := pageI["data"].(*simplejson.Json).MustMap()["pages"].(map[string]interface{})["docs"].([]interface{})
				images = utils.MergeSlices(images, imagesInPageI)
			}
		}

		ctx.HTML(http.StatusOK, "episode.html", map[string]interface{}{"ep": ep, "images": images})
	}
}
