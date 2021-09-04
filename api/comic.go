package api

import (
	"encoding/json"
	"fmt"
	"github.com/lvliangxiong/picago/api/model"
)

// GetComicDetail fetch comic information in detail, needs token and comicId, return a model.ComicDetail if successful.
func GetComicDetail(token string, comicId string) (int, interface{}, model.ComicDetail) {
	resultMap := send("/comics/"+comicId, "GET", token, "")

	comicDetail := model.ComicDetail{}

	if statusCode := int(resultMap["code"].(float64)); statusCode != 200 {
		return statusCode, resultMap["message"], comicDetail
	}

	comicJsonString, _ := json.Marshal(resultMap["data"].(map[string]interface{})["comic"])
	json.Unmarshal(comicJsonString, &comicDetail)

	return 200, "success", comicDetail
}

// EpisodeInfo fetch comic's episodes information, needs token, comicId and an optional page no,
// return a model.Episode slice if successful and additional information about pages.
func EpisodeInfo(token string, comicId string, pageStr string) (
	code int, message interface{}, episodes []model.Episode, page int, pages int, limit int, total int,
) {
	resultMap := send(fmt.Sprintf("/comics/%s/eps?page=%s", comicId, pageStr), "GET", token, "")

	if statusCode := int(resultMap["code"].(float64)); statusCode != 200 {
		return statusCode, resultMap["message"], nil, 0, 0, 0, 0
	}

	data := resultMap["data"].(map[string]interface{})["eps"].(map[string]interface{})

	code, message = 200, "success"
	page, pages = int(data["page"].(float64)), int(data["pages"].(float64))
	limit, total = int(data["limit"].(float64)), int(data["total"].(float64))

	episodesJsonString, _ := json.Marshal(data["docs"])
	episodes = make([]model.Episode, 0, total)
	json.Unmarshal(episodesJsonString, &episodes)

	return
}

// EpisodeDetail fetch information of a comic episode, needs token, comicId, episodeOrder, and an optional page no,
// return a model.ComicImage slice, model.Episode, and some additional information about pages.
func EpisodeDetail(token string, comicId string, order string, pageStr string) (
	code int, message interface{},
	images []model.ComicImage, episode model.Episode,
	page int, pages int, limit int, total int,
) {
	resultMap := send(
		fmt.Sprintf("/comics/%s/order/%s/pages?page=%s", comicId, order, pageStr),
		"GET", token, "",
	)

	if statusCode := int(resultMap["code"].(float64)); statusCode != 200 {
		return statusCode, resultMap["message"], nil, episode, 0, 0, 0, 0
	}

	data := resultMap["data"].(map[string]interface{})["pages"].(map[string]interface{})

	// sample: {"_id":"60996d2d59ef3308ba59c4c6","title":"第56話"}
	episodeJson, _ := json.Marshal(resultMap["data"].(map[string]interface{})["ep"])
	json.Unmarshal(episodeJson, &episode)

	code, message = 200, "success"
	page, pages = int(data["page"].(float64)), int(data["pages"].(float64))
	limit, total = int(data["limit"].(float64)), int(data["total"].(float64))

	imagesJsonString, _ := json.Marshal(data["docs"])
	images = make([]model.ComicImage, 0, total)
	json.Unmarshal(imagesJsonString, &images)

	return
}
