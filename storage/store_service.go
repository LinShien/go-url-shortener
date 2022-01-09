package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// wrapper around raw Redis client
type StorageService struct {
	redisClient *redis.Client
}

var (
	storageService = &StorageService{}
	ctx            = context.Background()
)

const CacheDuration = 6 * time.Hour

// Initializing the store service and return a store pointer
func InitializeStorage() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()

	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started sucessfully: pong message = {%s}", pong)
	storageService.redisClient = redisClient

	return storageService
}

/*
save the mapping between the original Url and the generated shortened Url url
*/
func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storageService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()

	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - short Url: %s - origianl Url: %s", err, shortUrl, originalUrl))
	}
}

/*
Let users use our short Url and redirect the original links for our users
*/
func RetrieveInitialUrl(shortUrl string) string {
	result, err := storageService.redisClient.Get(ctx, shortUrl).Result()

	if err != nil {
		panic(fmt.Sprintf("Failed retrieve original Url | Error: %v - short Url: %s", err, shortUrl))
	}

	return result
}
