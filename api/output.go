package api

func errorOutput(code int, err string, msg string) map[string]interface{} {
	return map[string]interface{}{"code": code, "error": err, "message": msg}
}

func successOutput(data interface{}) map[string]interface{} {
	return map[string]interface{}{"code": 200, "message": "success", "data": data}
}
