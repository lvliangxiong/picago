package api

import (
	"crypto/tls"
	"github.com/bitly/go-simplejson"
	"github.com/parnurzeal/gorequest"
	UUID "github.com/satori/go.uuid"
	"pica.go/conf"
	"pica.go/utils"
	"strconv"
	"strings"
	"time"
)

func send(url string, method string, authorization string, payload string) simplejson.Json {
	headers := utils.CopyStringStringMap(conf.Headers)

	url = "https://picaapi.picacomic.com" + url

	appUUID := UUID.NewV4().String()
	host := "picaapi.picacomic.com"
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := strings.Replace(UUID.NewV4().String(), "-", "", -1)

	signature := strings.Replace(url, "https://picaapi.picacomic.com/", "", -1)
	signature = strings.ToLower(signature + timestamp + nonce + method + headers["Api-Key"])
	signature = utils.ComputeHmacSha256(signature, conf.SecretKey)

	request := gorequest.New()

	if method == "GET" {
		request.Get(url)
	} else {
		request.Post(url)
	}

	// update headers
	headers["App-Uuid"] = appUUID
	headers["Host"] = host
	headers["Time"] = timestamp
	headers["Nonce"] = nonce
	headers["Signature"] = signature
	headers["Authorization"] = authorization

	setHeaders(request, headers)

	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	var body string

	if method == "POST" {
		request.Send(payload)
	}
	_, body, _ = request.End()

	json, _ := simplejson.NewJson([]byte(body))

	// return the result as a *Json pointer
	return *json
}

func sendImageRequest(fileServer string, path string, authorization string) gorequest.Response {
	headers := utils.CopyStringStringMap(conf.Headers)

	url := fileServer + "/static/" + path

	appUUID := UUID.NewV4().String()
	host := strings.Replace(fileServer, "https://", "", 1)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := strings.Replace(UUID.NewV4().String(), "-", "", -1)

	signature := "static/" + path
	signature = strings.ToLower(signature + timestamp + nonce + "GET" + headers["Api-Key"])
	signature = utils.ComputeHmacSha256(signature, conf.SecretKey)

	request := gorequest.New()
	request.Get(url)

	// update headers
	headers["App-Uuid"] = appUUID
	headers["Host"] = host
	headers["Time"] = timestamp
	headers["Nonce"] = nonce
	headers["Signature"] = signature
	headers["Authorization"] = authorization

	setHeaders(request, headers)

	resp, _, _ := request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true}).End()
	return resp
}

func setHeaders(request *gorequest.SuperAgent, headers map[string]string) {
	for k, v := range headers {
		request.Set(k, v)
	}
}
