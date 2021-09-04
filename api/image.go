package api

import "github.com/parnurzeal/gorequest"

// GetImage fetch image data from pica server and return a gorequest.Response object.
func GetImage(token string, fileServer string, path string) gorequest.Response {
	return sendImageRequest(fileServer, path, token)
}
