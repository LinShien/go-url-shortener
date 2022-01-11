package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StoreService{}

func init() {
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://www.youtube.com/watch?v=VafTMsrnSTU"
	userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortUrl := "Jsz4k57oAX"

	SaveUrlMapping(shortUrl, initialLink, userUUId)

	retrievedUrl := RetrieveInitialUrl(shortUrl)

	assert.Equal(t, retrievedUrl, initialLink)
}
