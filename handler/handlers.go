package handler

import (
	"net/http"

	"github.com/LinShien/go-url-shortener/shortener"
	"github.com/LinShien/go-url-shortener/storage"
	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrlHandler(ctx *gin.Context) {
	var creationRequest UrlCreationRequest

	if err := ctx.ShouldBindJSON(&creationRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortUrl(creationRequest.LongUrl, creationRequest.UserId)
	storage.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	host := "http://localhost:9808/"

	ctx.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func ShortUrlRedirectHandler(ctx *gin.Context) {
	shortUrl := ctx.Param("shortUrl")
	initialLink := storage.RetrieveInitialUrl(shortUrl)
	ctx.Redirect(302, initialLink)
}
