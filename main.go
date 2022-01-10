package main

import (
	"fmt"
	"net/http"

	"github.com/LinShien/go-url-shortener/handler"
	"github.com/LinShien/go-url-shortener/storage"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")
	{
		api.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	storage.InitializeStorage()

	api.POST("/create-short-url", handler.CreateShortUrlHandler)
	api.GET("/:shortUrl", handler.ShortUrlRedirectHandler)

	err := router.Run(":9808")

	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
