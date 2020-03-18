package repositories

import (
	"errors"
	"fmt"
	"github.com/Dadard29/go-youtube-dl/api"
	"github.com/Dadard29/go-youtube-dl/models"
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

