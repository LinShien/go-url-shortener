package shortener

import (
	"crypto/sha256"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
)

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input)) // ascii bytes array to 64-bit hex array

	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoded := base58.Encode(bytes)

	return string(encoded)
}

func GenerateShortUrl(initialLink string, userId string) string {
	urlHashBytes := sha256Of(initialLink + userId)

	shortUrl := base58Encoded([]byte(fmt.Sprintf("%x", urlHashBytes)))

	return shortUrl[:8]
}
