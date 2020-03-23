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
