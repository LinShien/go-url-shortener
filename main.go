package main

import (
	"fmt"

	"github.com/LinShien/go-url-shortener/handler"
	"github.com/LinShien/go-url-shortener/store"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	store.InitializeStore()

	router.POST("/create-short-url", handler.CreateShortUrlHandler)
	router.GET("/:shortUrl", handler.ShortUrlRedirectHandler)

	err := router.Run(":9808")

	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
