package controllers

import (
	"github.com/Dadard29/go-api-utils/auth"
	"github.com/Dadard29/go-youtube-dl/api"
	"github.com/Dadard29/go-youtube-dl/managers"
	"net/http"
)

// GET
// Authorization: 	token
// Params: 			None
// Body: 			None
func DownloadGet(w http.ResponseWriter, r *http.Request) {

	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, w) {
		return
	}

	s, err := managers.GetDownloadStatus(accessToken)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusInternalServerError, "error getting last status", w)
		return
	}
	api.Api.BuildJsonResponse(true, "status retrieved", s, w)
}

// GET
// Authorization: 	token
// Params: 			None
// Body: 			None
func DownloadDelete(w http.ResponseWriter, r *http.Request) {

	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, w) {
		return
	}

	managers.CancelDownload(accessToken)
	api.Api.BuildJsonResponse(true, "download cancelled", nil, w)
}

// POST
// Authorization: 	token
// Params: 			videoId
// Body: 			None
func DownloadPost(w http.ResponseWriter, r *http.Request) {

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

	err := managers.Download(accessToken, videoId)
	if err != nil {
		api.Api.BuildErrorResponse(http.StatusInternalServerError,
			"error launching download", w)
		return
	}

	api.Api.BuildJsonResponse(true, "download launch", nil, w)
}

// GET
// Authorization: 	token
// Params: 			None
// Body: 			None
func FileGet(w http.ResponseWriter, r *http.Request) {

	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, w) {
		return
	}

	p, err := managers.GetDownloadFile(accessToken)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusInternalServerError, err.Error(), w)
		return
	}

	w.Header().Add("Access-Control-Allow-Origin", api.Api.Service.CorsOrigin())
	http.ServeFile(w, r, p)
}

func DownloadAllPost(w http.ResponseWriter, r *http.Request) {

	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, w) {
		return
	}

	err := managers.DownloadAll(accessToken)
	if err != nil {
		api.Api.BuildErrorResponse(http.StatusInternalServerError,
			"error launching download", w)
		return
	}

	api.Api.BuildJsonResponse(true, "download launch", nil, w)
}