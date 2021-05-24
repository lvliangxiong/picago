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
		pica.POST("/loginCheck", handler.LoginCheck)
		pica.GET("/index", handler.ShowCategory)
		pica.GET("/comics", handler.GetComics)
		pica.GET("/comic/:comicId", handler.GetComic)
		pica.GET("/comic/:comicId/episode/:order", handler.ReadComic)
		pica.GET("/img", handler.GetImage)
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

	go func() {
		for {
			if conf.PublicToken != "" {
				// Clear public token every two hours
				time.Sleep(2 * time.Hour)
				conf.PublicToken = ""
			}
			// Check every 2 minutes
			time.Sleep(2 * time.Minute)
		}
	}()

	// Start Server
	r.Run(conf.ServerPort)
}
