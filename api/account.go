package api

import (
	"fmt"
)

// Login tries to login the pica server, return token if successful.
func Login(username string, password string) (code int, message interface{}, token string) {
	if username == "" || password == "" {
		code = 400
		message = "Please provide email and password to login!"
		return
	}

	resultMap := send(
		"/auth/sign-in", "POST", "",
		fmt.Sprintf(`{"email":"%s", "password":"%s"}`, username, password),
	)

	if statusCode := int(resultMap["code"].(float64)); statusCode != 200 {
		code = statusCode
		message = resultMap["message"]
		return
	}

	return 200, "success", resultMap["data"].(map[string]interface{})["token"].(string)
}
