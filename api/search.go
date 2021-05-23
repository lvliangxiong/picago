package api

import (
	"fmt"
	"net/url"
)

// Fetch all hot searched keywords
func HotSearchKeywords(token string) map[string]interface{} {
	result := send("/keywords", "GET", token, "")

	if code := result.Get("code").MustInt(); code != 200 {
		return errorOutput(code, result.Get("error").MustString(), result.Get("message").MustString())
	}

	return successOutput(result.Get("data").Get("keywords"))
}

// Search comics according to keyword and page no
func SearchByKeyword(token string, keyword string, page string) map[string]interface{} {
	result := send(fmt.Sprintf("/comics/search?page=%s&q=%s", page, url.QueryEscape(keyword)),
		"GET", token, "")

	if code := result.Get("code").MustInt(); code != 200 {
		return errorOutput(code, result.Get("error").MustString(), result.Get("message").MustString())
	}

	return successOutput(result.Get("data").Get("comics"))
}

func SearchByKeywordAndCategory(token string, keyword string, categories []string, page string, sort string) map[string]interface{} {
	result := send(fmt.Sprintf(`/comics/advanced-search?page=%s`, page),
		"POST", token,
		fmt.Sprintf(`{"categories":%s, "sort":"%s", "keyword":"%s"}`, strArrToString(categories), sort, ""))

	if code := result.Get("code").MustInt(); code != 200 {
		return errorOutput(code, result.Get("error").MustString(), result.Get("message").MustString())
	}

	return successOutput(result.Get("data").Get("comics"))
}

func strArrToString(strs []string) string {
	for i, str := range strs {
		strs[i] = `"` + str + `"`
	}
	return fmt.Sprint(strs)
}
