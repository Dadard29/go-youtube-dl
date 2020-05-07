package controllers

import (
	"encoding/json"
	"github.com/Dadard29/go-api-utils/auth"
	"github.com/Dadard29/go-youtube-dl/api"
	"github.com/Dadard29/go-youtube-dl/managers"
	"github.com/Dadard29/go-youtube-dl/models"
	"io/ioutil"
	"net/http"
)

// GET
// Authorization: 	token
// Params: 			videoId
// Body: 			None
func VideoGet(w http.ResponseWriter, r *http.Request) {

	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, w) {
		return
	}

	videoId := r.URL.Query().Get("videoId")
	if videoId == "" {
		api.Api.BuildErrorResponse(http.StatusBadRequest,
			"missing parameter", w)
		return
	}

	v, err := managers.VideoManagerGet(accessToken, videoId)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusInternalServerError,
			"error getting video", w)
		return
	}

	api.Api.BuildJsonResponse(true, "video retrieved", v, w)
}

// POST
// Authorization: 	token
// Params: 			videoId
// Body: 			None
func VideoPost(w http.ResponseWriter, r *http.Request) {

	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, w) {
		return
	}

	videoId := r.URL.Query().Get("videoId")
	if videoId == "" {
		api.Api.BuildErrorResponse(http.StatusBadRequest,
			"missing parameter", w)
		return
	}

	v, err := managers.VideoManagerCreate(accessToken, videoId)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusInternalServerError,
			"error creating video", w)
		return
	}

	api.Api.BuildJsonResponse(true, "video created", v, w)
}


// GET
// Authorization: 	token
// Params: 			queryParam
// Body: 			None

// add a video from keywords
func VideoSearchPost(w http.ResponseWriter, r *http.Request) {
	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, w) {
		return
	}

	title := r.URL.Query().Get(titleParam)
	artist := r.URL.Query().Get(artistParam)
	album := r.URL.Query().Get(albumParam)
	genre := r.URL.Query().Get(genreParam)
	publishedAt := r.URL.Query().Get(dateParam)
	if title == "" || artist == "" || album == "" || genre == "" || publishedAt == ""{
		api.Api.BuildMissingParameter(w)
		return
	}

	v, err := managers.VideoManagerSearch(accessToken, title, artist, album, genre, publishedAt)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusInternalServerError, "error creating video", w)
		return
	}

	api.Api.BuildJsonResponse(true, "video created", v, w)

}

// PUT
// Authorization: 	token
// Params: 			None
// Body: 			models.VideoJson
func VideoPut(w http.ResponseWriter, r *http.Request) {

	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, w) {
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusBadRequest, "invalid body", w)
		return
	}

	var j models.VideoJson
	err = json.Unmarshal(body, &j)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusBadRequest, "invalid json body", w)
		return
	}

	v, err := managers.VideoManagerUpdate(accessToken, j)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusInternalServerError,
			"error updating video", w)
		return
	}

	api.Api.BuildJsonResponse(true, "video updated", v, w)
}

// DELETE
// Authorization: 	token
// Params: 			videoId
// Body: 			None
func VideoDelete(w http.ResponseWriter, r *http.Request) {

	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, w) {
		return
	}

	videoId := r.URL.Query().Get("videoId")
	if videoId == "" {
		api.Api.BuildErrorResponse(http.StatusBadRequest,
			"missing parameter", w)
		return
	}

	v, err := managers.VideoManagerRemove(accessToken, videoId)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusInternalServerError,
			"error deleting video", w)
		return
	}

	api.Api.BuildJsonResponse(true, "video deleted", v, w)
}
