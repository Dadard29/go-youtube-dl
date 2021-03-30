package repositories

import (
	"github.com/Dadard29/go-youtube-dl/models"
	"testing"
	"time"
)

func TestDownload(t *testing.T) {
	m := models.VideoModel{
		Id:       1,
		VideoId:  "83f8Qe6zo3A",
		Token:    "token",
		Title:    "toitsu",
		Album:    "album",
		Artist:   "senbei",
		Date:     time.Now(),
		Genre:    "rap",
		ImageUrl: "url",
	}
	Download(m, ".", "file.mp3")
}
