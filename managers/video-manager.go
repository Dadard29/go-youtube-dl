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

func getVideoDetails(videoId string) (models.VideoYoutubeModel, error) {
	var v models.VideoYoutubeModel
	infosUrl := "https://www.googleapis.com/youtube/v3/videos?part=%s&id=%s&key=%s"

	part := "snippet"
	key := os.Getenv("YT_API_KEY")

	resp, err := http.Get(fmt.Sprintf(infosUrl, part, videoId, key))
	if err != nil {
		return v, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return v, err
	}

	var out models.VideoYoutubeModel
	err = json.Unmarshal(body, &out)
	if err != nil {
		return v, err
	}

	return out, nil
}

func VideoManagerGet(token string, videoId string) (models.VideoJson, error) {
	var v models.VideoJson

	videoDb, err := repositories.VideoGet(token, videoId)
	if err != nil {
		logger.Error(err.Error())
		return v, errors.New("error getting video")
	}

	return models.NewVideoJson(videoDb), err
}

func VideoManagerCreate(token string, videoId string) (models.VideoJson, error) {

	var v models.VideoJson

	videoInfos, err := getVideoDetails(videoId)
	if err != nil {
		return v, err
	}

	if len(videoInfos.Items) < 1 {
		return v, errors.New("this video ID is invalid")
	}

	videoYt := videoInfos.Items[0]
	title := unidecode(videoYt.Snippet.Title)
	channel := unidecode(videoYt.Snippet.ChannelTitle)
	publishedAt := videoYt.Snippet.PublishedAt

	video := models.VideoModel{
		Id:       0,
		VideoId:  videoId,
		Token:    token,
		Title:    title,
		Album:    channel,
		Artist:   channel,
		Date:     publishedAt,
		Genre:    "",
		ImageUrl: "",
	}

	videoDb, err := repositories.VideoCreate(video)
	if err != nil {
		return v, err
	}

	return models.NewVideoJson(videoDb), nil
}

func VideoManagerUpdate(token string, videoJson models.VideoJson) (models.VideoJson, error) {
	var v models.VideoJson

	t, err := time.Parse("2006-01-02", videoJson.Date)
	if err != nil {
		return v, err
	}

	videoDb, err := repositories.VideoUpdate(models.VideoModel{
		VideoId: videoJson.VideoId,
		Token:   token,
		Title:   videoJson.Title,
		Album:   videoJson.Album,
		Artist:  videoJson.Artist,
		Date:    t,
		Genre: videoJson.Genre,
		ImageUrl: videoJson.ImageUrl,
	})
	if err != nil {
		return v, err
	}

	return models.NewVideoJson(videoDb), nil
}

func VideoManagerRemove(token string, videoId string) (models.VideoJson, error) {
	var v models.VideoJson
	videoDb, err := repositories.VideoGet(token, videoId)
	if err != nil {
		return v, err
	}

	videoDeleted, err := repositories.VideoDelete(videoDb)
	if err != nil {
		return v, err
	}

	return models.NewVideoJson(videoDeleted), nil
}
