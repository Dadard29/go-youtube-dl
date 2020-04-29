package managers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Dadard29/go-youtube-dl/models"
	"github.com/Dadard29/go-youtube-dl/repositories"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func getPlaylistDetails(playlistId string) ([]models.PlaylistYoutubeItem, error) {
	var v []models.PlaylistYoutubeItem
	playlistUrl :=
		"https://www.googleapis.com/youtube/v3/playlistItems?part=%s&playlistId=%s&key=%s&maxResults=%d&pageToken=%s"
	part := "snippet"
	key := os.Getenv("YT_API_KEY")
	maxResults := 25
	nextPageToken := ""

	var vList []models.PlaylistYoutubeItem
	for true {
		resp, err := http.Get(
			fmt.Sprintf(playlistUrl, part, playlistId, key, maxResults, nextPageToken))
		if err != nil {
			return v, err
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return v, err
		}

		var out models.PlaylistYoutubeModel
		err = json.Unmarshal(body, &out)
		if err != nil {
			return v, err
		}

		// add the items to the list
		for _, item := range out.Items {
			vList = append(vList, item)
		}

		// if no other page, break the loop
		nextPageToken = out.NextPageToken
		if nextPageToken == "" {
			break
		}
	}

	return vList, nil
}

func VideoListManagerCreateFromPlaylist(token string, playlistId string) ([]models.VideoJson, error) {
	var g = make([]models.VideoJson, 0)

	vList, err := getPlaylistDetails(playlistId)
	if err != nil {
		return g, err
	}

	var vJsonList = make([]models.VideoJson, 0)
	if len(vList) == 0 {
		return g, errors.New("error getting playlist info from Youtube API")
	}


	for _, v := range vList {

		videoId := v.Snippet.ResourceID.VideoID
		title := unidecode(v.Snippet.Title)
		channel := unidecode(v.Snippet.ChannelTitle)
		publishedAt := v.Snippet.PublishedAt

		vModel := models.VideoModel{
			VideoId: videoId,
			Token:   token,
			Title:   title,
			Album:   channel,
			Artist:  channel,
			Date:    publishedAt,
			Genre:   "",
		}

		videoDb, err := repositories.VideoCreate(vModel)
		if err != nil {
			logger.Error(err.Error())
			continue
		}

		vJsonList = append(vJsonList, models.NewVideoJson(videoDb))
	}

	return vJsonList, nil
}

func VideoListManagerDelete(token string, json models.VideoDeleteAllJson) ([]models.VideoJson, error) {
	var vList = make([]models.VideoJson, 0)
	for _, v := range json.VideoList {
		v, err := repositories.VideoGet(token, v)
		if err != nil {
			logger.Error(err.Error())
			continue
		}

		vDeleted, err := repositories.VideoDelete(v)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		vList = append(vList, models.NewVideoJson(vDeleted))
	}

	return vList, nil
}

func VideoListManagerUpdate(token string, json models.VideoUpdateAllJson) ([]models.VideoJson, error) {
	infos := json.Infos

	var vList = make([]models.VideoJson, 0)
	for _, videoId := range json.VideoList {
		v, err := repositories.VideoGet(token, videoId)
		if err != nil {
			logger.Error(err.Error())
			continue
		}

		if infos.Title != "" {
			v.Title = infos.Title
		}

		if infos.Album != "" {
			v.Album = infos.Album
		}

		if infos.Artist != "" {
			v.Artist = infos.Artist
		}

		if infos.Date != "" {
			d, err := time.Parse("2006-01-02", infos.Date)
			if err != nil {
				logger.Error(err.Error())
				continue
			}
			v.Date = d
		}

		if infos.Genre != "" {
			v.Genre = infos.Genre
		}

		vUpdated, err := repositories.VideoUpdate(v)
		if err != nil {
			logger.Error(err.Error())
			continue
		}

		vList = append(vList, models.NewVideoJson(vUpdated))
	}

	return vList, nil
}

func VideoListManagerGet(token string) ([]models.VideoJson, error) {
	var l = make([]models.VideoJson, 0)

	listDb, err := repositories.VideoGetList(token)
	if err != nil {
		return l, err
	}

	var vList = make([]models.VideoJson, 0)
	for _, v := range listDb {
		vList = append(vList, models.NewVideoJson(v))
	}

	return vList, nil
}
