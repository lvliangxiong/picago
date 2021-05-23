package api

import (
	"fmt"
)

func Login(username string, password string) map[string]interface{} {
	if username == "" || password == "" {
		return errorOutput(400, "110", "Please provide email and password to login!")
	}

	/*
		{
		  "code": 200,
		  "message": "success",
		  "data": {
			"token": "..."
		  }
		}

		or

		{"code":400,"error":"1004","message":"invalid email or password","detail":":("}
	*/
	result := send("/auth/sign-in", "POST", "",
		fmt.Sprintf(`{"email":"%s", "password":"%s"}`, username, password))

	if code := result.Get("code").MustInt(); code != 200 {
		return errorOutput(code, result.Get("error").MustString(), result.Get("message").MustString())
	}

	return successOutput(result.Get("data").Get("token"))
}
