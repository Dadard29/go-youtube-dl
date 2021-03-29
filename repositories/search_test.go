package repositories

import (
	"testing"
)

func TestSearch(t *testing.T) {
	token := "token"
	title := "my home"
	artist := "igit"
	album := "album"
	genre := "genre"
	publishedAt := "2020-03-03"
	_, err := VideoSearch(token, title, artist, album, genre, publishedAt)
	if err != nil {
		t.Error(err.Error())
	}
}
