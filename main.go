package main

import (
	"fmt"

	"github.com/LinShien/go-url-shortener/handler"
	"github.com/LinShien/go-url-shortener/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API ! " + ctx.ClientIP(),
		})
	})

	storage.InitializeStorage()

	router.POST("/create-short-url", handler.CreateShortUrlHandler)
	router.GET("/:shortUrl", handler.ShortUrlRedirectHandler)

	err := router.Run(":9808")

	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
