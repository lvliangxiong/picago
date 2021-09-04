package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lvliangxiong/picago/api"
	"github.com/lvliangxiong/picago/api/model"
	"github.com/lvliangxiong/picago/utils"
	"net/http"
	"strconv"
)

func GetComics(ctx *gin.Context) {
	token, err := utils.GetToken(ctx)

	if err != nil {
		return
	}

	category, pageStr := ctx.Query("category"), ctx.Query("page")
	if pageStr == "" {
		pageStr = "1"
	}

	code, _, comics, curPage, pageCount, _, _ := api.SearchByKeywordAndCategory(
		token, "", []string{category}, pageStr, model.Newest,
	)

	if code != 200 {
		return
	}

	var previousPage, nextPage int
	previousPage = curPage - 1
	if curPage < pageCount {
		nextPage = curPage + 1
	} else {
		nextPage = 0
	}

	data := map[string]interface{}{
		"category":     category,
		"comics":       comics,
		"previousPage": previousPage,
		"nextPage":     nextPage,
	}
	ctx.HTML(http.StatusOK, "comics.html", data)
}

// GetComic show comic information including a list containing all episodes
func GetComic(ctx *gin.Context) {
	token, err := utils.GetToken(ctx)
	if err == nil {
		comicId := ctx.Param("comicId")

		code, _, comicDetail := api.GetComicDetail(token, comicId)

		if code != 200 {
			return
		}

		code, _, episodes, _, pageCount, _, _ := api.EpisodeInfo(token, comicId, "1")
		if code != 200 {
			return
		}

		if pageCount > 1 {
			for i := 2; i <= pageCount; i++ {
				_, _, eps, _, _, _, _ := api.EpisodeInfo(token, comicId, strconv.Itoa(i))
				episodes = append(episodes, eps...)
			}
		}

		comic := map[string]interface{}{}

		comic["info"] = comicDetail
		comic["episodes"] = episodes

		ctx.HTML(http.StatusOK, "comic.html", comic)
	}
}

// ReadComic show all images in one episode
func ReadComic(ctx *gin.Context) {
	token, err := utils.GetToken(ctx)
	if err == nil {
		comicId, order := ctx.Param("comicId"), ctx.Param("order")

		// Fetch the first page of images
		code, _, images, episode, _, pageCount, _, _ := api.EpisodeDetail(token, comicId, order, "1")
		if code != 200 {
			ctx.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/pica/comic/%s", comicId))
		}

		if pageCount > 1 {
			// if there are more than one page, fetch other pages of images
			for i := 2; i <= pageCount; i++ {
				_, _, imgs, _, _, _, _, _ := api.EpisodeDetail(token, comicId, order, strconv.Itoa(i))
				images = append(images, imgs...)
			}
		}

		orderNo, _ := strconv.Atoi(order)

		ctx.HTML(
			http.StatusOK, "episode.html",
			map[string]interface{}{
				"episode":  episode,
				"images":   images,
				"comicId":  comicId,
				"previous": orderNo - 1,
				"next":     orderNo + 1,
			},
		)
	}
}
