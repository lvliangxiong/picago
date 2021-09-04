package api

import (
	"encoding/json"
	"github.com/lvliangxiong/picago/api/model"
)

// FetchCategories tries to fetch category information from pica server, return a model.Category slice if successful.
func FetchCategories(token string) (int, interface{}, []model.Category) {
	resultMap := send("/categories", "GET", token, "")

	if statusCode := int(resultMap["code"].(float64)); statusCode != 200 {
		return statusCode, resultMap["message"], nil
	}

	catArrJsonString, _ := json.Marshal(resultMap["data"].(map[string]interface{})["categories"])
	categories := make([]model.Category, 48, 48)
	json.Unmarshal(catArrJsonString, &categories)
	return 200, "success", categories
}
