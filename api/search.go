package api

import (
	"encoding/json"
	"fmt"
	"github.com/lvliangxiong/picago/api/model"
)

// SearchByKeywordAndCategory fetch comics information from pica server according to token, keyword, category,
// page and sort information provided, returns a model.Comic slice and additional pages formation.
func SearchByKeywordAndCategory(
	token string, keyword string, categories []string, pageStr string, sort string,
) (code int, message interface{}, comics []model.Comic, page int, pages int, limit int, total int) {
	resultMap := send(
		fmt.Sprintf(`/comics/advanced-search?page=%s`, pageStr),
		"POST", token,
		fmt.Sprintf(`{"categories":%s, "sort":"%s", "keyword":"%s"}`, fmt.Sprintf("%q", categories), sort, keyword),
	)

	if statusCode := int(resultMap["code"].(float64)); statusCode != 200 {
		return statusCode, resultMap["message"], nil, 0, 0, 0, 0
	}

	data := resultMap["data"].(map[string]interface{})["comics"].(map[string]interface{})

	code, message = 200, "success"
	page, pages = int(data["page"].(float64)), int(data["pages"].(float64))
	limit, total = int(data["limit"].(float64)), int(data["total"].(float64))

	comicsJsonString, _ := json.Marshal(data["docs"])
	comics = make([]model.Comic, limit, limit)
	json.Unmarshal(comicsJsonString, &comics)

	return
}