package router

import (
	"github.com/gin-gonic/gin"
	"pica.go/conf"
	"pica.go/handler"
	"pica.go/middleware"
	"time"
)

func Init() {
	r := gin.Default()

	r.LoadHTMLGlob("template/*")

	pica := r.Group("/pica")
	pica.Use(middleware.ValidateToken)

	{
		pica.GET("/index", handler.ShowCategory)
		pica.POST("/loginCheck", handler.LoginCheck)
		pica.GET("/img", handler.GetImage)
		pica.GET("/comics", handler.GetComics)
	}

	// Init configuration
	conf.Init()
	// Update configurations every 60s
	go func() {
		for {
			time.Sleep(60 * time.Second)
			conf.Init()
		}
	}()

	// Start Server
	r.Run(conf.ServerPort)
}
