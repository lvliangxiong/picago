package api

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

// Fetch comic information, needs token and comicId
func ComicInfo(token string, comicId string) map[string]interface{} {
	result := send("/comics/"+comicId, "GET", token, "")

	if code := result.Get("code").MustInt(); code != 200 {
		return errorOutput(code, result.Get("error").MustString(), result.Get("message").MustString())
	}

	return successOutput(result.Get("data").Get("comic"))
}

// Fetch comic's episodes information, needs token, comicId and an optional page no
func EpisodeInfo(token string, comicId string, page string) map[string]interface{} {
	result := send(fmt.Sprintf("/comics/%s/eps?page=%s", comicId, page), "GET", token, "")

	if code := result.Get("code").MustInt(); code != 200 {
		return errorOutput(code, result.Get("error").MustString(), result.Get("message").MustString())
	}

	return successOutput(result.Get("data").Get("eps"))
}

// Fetch information of a comic episode, needs token, comicId, episodeOrder, and an optional page no
func EpisodeDetail(token string, comicId string, order string, page string) map[string]interface{} {
	result := send(fmt.Sprintf("/comics/%s/order/%s/pages?page=%s", comicId, order, page),
		"GET", token, "")

	if code := result.Get("code").MustInt(); code != 200 {
		return errorOutput(code, result.Get("error").MustString(), result.Get("message").MustString())
	}

	return successOutput(result.Get("data"))
}

func ComicImage(token string, fileServer string, path string) gorequest.Response {
	return sendImageRequest(fileServer, path, token)
}
