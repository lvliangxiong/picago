package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lvliangxiong/picago/conf"
)

func CopyMap(m map[string]interface{}) map[string]interface{} {
	cp := make(map[string]interface{})
	for k, v := range m {
		vm, ok := v.(map[string]interface{})
		if ok {
			cp[k] = CopyMap(vm)
		} else {
			cp[k] = v
		}
	}
	return cp
}

func CopyStringStringMap(m map[string]string) map[string]string {
	cp := make(map[string]string)
	for k, v := range m {
		cp[k] = v
	}
	return cp
}

func ComputeHmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

func GetToken(ctx *gin.Context) (string, error) {
	token, err := ctx.Cookie("token")
	if err != nil || token == "" {
		// When cookie has no token, check the public token
		if conf.AllowRememberTokenForAllUsers && conf.PublicToken != "" {
			return conf.PublicToken, nil
		} else {
			return "", errors.New("no token found")
		}
	}
	return token, nil
}
