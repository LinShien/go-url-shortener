package handler

import (
	"net/http"
	"net/url"

	"github.com/LinShien/go-url-shortener/shortener"
	"github.com/LinShien/go-url-shortener/store"
	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required,max=2048"`
	UserId  string `json:"user_id" binding:"required"`
}

func isValidUrl(link string) bool {
	u, err := url.Parse(link)

	if u.Scheme == "" {
		link = "http://" + link
	}

	u, err = url.Parse(link)

	return err == nil && u.Scheme != "" && u.Host != ""
}

func CreateShortUrlHandler(ctx *gin.Context) {
	var creationRequest UrlCreationRequest

	if err := ctx.ShouldBindJSON(&creationRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !isValidUrl(creationRequest.LongUrl) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "The provided url is not valid"})
		return
	}

	shortUrl := shortener.GenerateShortUrl(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	host := "http://localhost:9808/"

	ctx.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func ShortUrlRedirectHandler(ctx *gin.Context) {
	shortUrl := ctx.Param("shortUrl")
	initialLink := store.RetrieveInitialUrl(shortUrl)
	ctx.Redirect(302, initialLink)
}
