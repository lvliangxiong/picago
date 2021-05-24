package conf

import (
	"errors"
	"fmt"
	"github.com/pelletier/go-toml"
)

var (
	SecretKey                     string
	ServerPort                    string
	AllowRememberTokenForAllUsers bool
	PublicToken                   string
)

var Headers = map[string]string{}

var picaRequired = []string{
	"secretKey",
}

var headersRequired = []string{
	"apiKey",
	"appVersion", "appChannel", "buildVersion", "appPlatform",
	"accept", "userAgent", "contentType", "imageQuality",
}

func Init() {
	// Load conf file
	config, _ := toml.LoadFile("conf/pica.toml")

	// Check conf required
	for _, key := range picaRequired {
		if !config.Has("pica." + key) {
			panic(errors.New(fmt.Sprintf("configuration '%s' missing", key)))
		}
	}

	for _, key := range headersRequired {
		if !config.Has("pica.headers." + key) {
			panic(errors.New(fmt.Sprintf("configuration '%s' missing", key)))
		}
	}

	// Assign or Update global variables
	SecretKey = config.Get("pica.secretKey").(string)

	Headers["Api-Key"] = config.Get("pica.headers.apiKey").(string)

	Headers["App-Version"] = config.Get("pica.headers.appVersion").(string)
	Headers["App-Channel"] = config.Get("pica.headers.appChannel").(string)
	Headers["App-Build-Version"] = config.Get("pica.headers.buildVersion").(string)
	Headers["App-Platform"] = config.Get("pica.headers.appPlatform").(string)

	Headers["Accept"] = config.Get("pica.headers.accept").(string)
	Headers["User-Agent"] = config.Get("pica.headers.userAgent").(string)
	Headers["Content-Type"] = config.Get("pica.headers.contentType").(string)

	Headers["Image-Quality"] = config.Get("pica.headers.imageQuality").(string)

	ServerPort = config.Get("server.port").(string)
	AllowRememberTokenForAllUsers = config.Get("server.allowRememberToken").(bool)
}
