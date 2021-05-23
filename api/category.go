package api

func Categories(token string) map[string]interface{} {
	/*
		{
		  "code": 200,
		  "message": "success",
		  "data": {
		    "categories": [
		      {
		        "title": "援助嗶咔",
		        "thumb": {
		          "originalName": "help.jpg",
		          "path": "help.jpg",
		          "fileServer": "https://wikawika.xyz/static/"
		        },
		        "isWeb": true,
		        "active": true,
		        "link": "https://donate.wikawika.xyz"
		      },
		      {
		        "title": "嗶咔小禮物",
		        "thumb": {
		          "originalName": "picacomic-gift.jpg",
		          "path": "picacomic-gift.jpg",
		          "fileServer": "https://wikawika.xyz/static/"
		        },
		        "isWeb": true,
		        "link": "https://gift-web.wikawika.xyz",
		        "active": true
		      },
		      ...
		    ]
		  }
		}
	*/
	result := send("/categories", "GET", token, "")

	if code := result.Get("code").MustInt(); code != 200 {
		return errorOutput(code, result.Get("error").MustString(), result.Get("message").MustString())
	}

	return successOutput(result.Get("data").Get("categories"))
}
