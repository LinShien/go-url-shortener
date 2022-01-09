package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const UserId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortUrlGenerator(t *testing.T) {
	initialLink_1 := "https://www.youtube.com/watch?v=HcTaeonqWPw"
	shortLink_1 := GenerateShortUrl(initialLink_1, UserId)

	initialLink_2 := "https://forum.gamer.com.tw/Co.php?bsn=60076&sn=80889903"
	shortLink_2 := GenerateShortUrl(initialLink_2, UserId)

	assert.Equal(t, shortLink_1, "23juYH1B")
	assert.Equal(t, shortLink_2, "2z2PUUEY")
}
