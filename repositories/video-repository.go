package repositories

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BrianAllred/goydl"
	"github.com/Dadard29/go-youtube-dl/api"
	"github.com/Dadard29/go-youtube-dl/models"
	"time"
)

func videoExists(videoId string, token string) bool {
	var v models.VideoModel
	api.Api.Database.Orm.Where(&models.VideoModel{
		VideoId: videoId,
		Token: token,
	}).First(&v)

	return v.VideoId == videoId
}

func VideoGet(token string, videoId string) (models.VideoModel, error) {
	var v models.VideoModel
	api.Api.Database.Orm.Where(&models.VideoModel{
		VideoId: videoId,
		Token: token,
	}).First(&v)

	if v.VideoId != videoId {
		msg := fmt.Sprintf("video with id %s not found", videoId)
		return v, errors.New(msg)
	}

	return v, nil
}

func VideoGetList(token string) ([]models.VideoModel, error) {
	var v []models.VideoModel
	api.Api.Database.Orm.Find(&v, &models.VideoModel{
		Token: token,
	})

	return v, nil
}

func VideoCreate(video models.VideoModel) (models.VideoModel, error) {
	var v models.VideoModel
	if videoExists(video.VideoId, video.Token) {
		msg := fmt.Sprintf("video with id %s already existing", video.VideoId)
		return v, errors.New(msg)
	}

	api.Api.Database.Orm.Create(&video)

	if !videoExists(video.VideoId, video.Token) {
		msg := fmt.Sprintf("video with id %s not found", video.VideoId)
		return v, errors.New(msg)
	}

	return video, nil
}

func VideoSearch(token string, title string, artist string, album string,
	genre string, publishedAt string) (models.VideoModel, error) {
	var f models.VideoModel

	youtubeDl := goydl.NewYoutubeDl()
	youtubeDl.Options.SkipDownload.Value = true
	youtubeDl.Options.PrintJSON.Value = true

	query := fmt.Sprintf("'%s %s'", artist, title)
	search := "ytsearch:" + query

	cmd, err := youtubeDl.Download(search)
	if err != nil {
		return f, err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(youtubeDl.Stdout)
	if err != nil {
		return f, err
	}

	cmd.Wait()

	data := buf.Bytes()
	var res models.VideoSearchModel
	err = json.Unmarshal(data, &res)
	if err != nil {
		return f, err
	}

	// parse time
	t, err := time.Parse("2006-01-02", publishedAt)
	if err != nil {
		return f, err
	}

	return models.VideoModel{
		VideoId:  res.ID,
		Token:    token,
		Title:    title,
		Album:    album,
		Artist:   artist,
		Date:     t,
		Genre:    genre,
		ImageUrl: res.Thumbnail,
	}, nil
}

func VideoUpdate(video models.VideoModel) (models.VideoModel, error) {
	var v models.VideoModel
	if !videoExists(video.VideoId, video.Token) {
		msg := fmt.Sprintf("video with id %s not found", video.VideoId)
		return v, errors.New(msg)
	}

	videoDb, err := VideoGet(video.Token, video.VideoId)
	if err != nil {
		return v, err
	}

	videoDb.Title = video.Title
	videoDb.Artist = video.Artist
	videoDb.Album = video.Album
	videoDb.Date = video.Date
	videoDb.Genre = video.Genre
	videoDb.ImageUrl = video.ImageUrl

	api.Api.Database.Orm.Save(&videoDb)


	return video, nil
}


func VideoDelete(video models.VideoModel) (models.VideoModel, error) {
	var v models.VideoModel
	if !videoExists(video.VideoId, video.Token) {
		msg := fmt.Sprintf("video with id %s not found", video.VideoId)
		return v, errors.New(msg)
	}

	api.Api.Database.Orm.Delete(&video)

	return video, nil
}

