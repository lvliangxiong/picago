package api

import (
	"crypto/tls"
	"github.com/bitly/go-simplejson"
	"github.com/lvliangxiong/pica.go/conf"
	"github.com/lvliangxiong/pica.go/utils"
	"github.com/parnurzeal/gorequest"
	uuid "github.com/satori/go.uuid"
	"strconv"
	"strings"
	"time"
)

func send(url string, method string, authorization string, payload string) simplejson.Json {
	url = "https://picaapi.picacomic.com" + url
	headers := utils.CopyStringStringMap(conf.Headers)

	request := gorequest.New()

	// build the header
	appUUID := uuid.NewV4().String()
	host := "picaapi.picacomic.com"
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := strings.Replace(uuid.NewV4().String(), "-", "", -1)

	signature := strings.Replace(url, "https://picaapi.picacomic.com/", "", -1)
	signature = strings.ToLower(signature + timestamp + nonce + method + headers["Api-Key"])
	signature = utils.ComputeHmacSha256(signature, conf.SecretKey)

	// update headers
	headers["App-Uuid"] = appUUID
	headers["Host"] = host
	headers["Time"] = timestamp
	headers["Nonce"] = nonce
	headers["Signature"] = signature
	headers["Authorization"] = authorization

	if method == "GET" {
		request.Get(url)
	} else {
		request.Post(url)
	}

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

	var (
		headers map[string]string
		url     string
	)

	request := gorequest.New()

	if !strings.Contains(fileServer, "static") {
		url = fileServer + "/static/" + path
		headers = utils.CopyStringStringMap(conf.Headers)

		appUUID := uuid.NewV4().String()
		host := strings.Replace(fileServer, "https://", "", 1)
		timestamp := strconv.FormatInt(time.Now().Unix(), 10)
		nonce := strings.Replace(uuid.NewV4().String(), "-", "", -1)

		signature := "static/" + path
		signature = strings.ToLower(signature + timestamp + nonce + "GET" + headers["Api-Key"])
		signature = utils.ComputeHmacSha256(signature, conf.SecretKey)

		// update headers
		headers["App-Uuid"] = appUUID
		headers["Host"] = host
		headers["Time"] = timestamp
		headers["Nonce"] = nonce
		headers["Signature"] = signature
		headers["Authorization"] = authorization
	} else {
		url = fileServer + path
		url = strings.Replace(url, "wikawika.xyz", "storage.wikawika.xyz", 1)
		headers = map[string]string{}
	}

	request.Get(url)
	setHeaders(request, headers)

	resp, _, _ := request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true}).End()
	return resp
}

func setHeaders(request *gorequest.SuperAgent, headers map[string]string) {
	for k, v := range headers {
		request.Set(k, v)
	}
}
