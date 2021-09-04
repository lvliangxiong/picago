package router

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lvliangxiong/picago/conf"
	"github.com/lvliangxiong/picago/handler"
	"github.com/lvliangxiong/picago/middleware"
)

func Init(staticResources embed.FS, templates embed.FS) {
	sr, err := fs.Sub(staticResources, "static")
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}

	r := gin.Default()

	t, _ := template.ParseFS(templates, "template/*.html")
	r.SetHTMLTemplate(t)

	pica := r.Group("/pica")
	pica.Use(middleware.ValidateToken)

	{
		pica.POST("/loginCheck", handler.LoginCheck)
		pica.GET("/index", handler.ShowCategory)
		pica.GET("/comics", handler.GetComics)
		pica.GET("/comic/:comicId", handler.GetComic)
		pica.GET("/comic/:comicId/episode/:order", handler.ReadComic)
		pica.GET("/img", handler.GetImage)
		pica.StaticFS("/static", http.FS(sr))
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
