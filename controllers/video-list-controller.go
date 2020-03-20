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
// Authorization: 	JWT + token
// Params: 			None
// Body: 			None
func VideoListGet(w http.ResponseWriter, r *http.Request) {
	jwt := auth.ParseApiKey(r, authorizationKey, true)
	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, jwt, w) {
		return
	}

	l, err := managers.VideoListManagerGet(accessToken)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusInternalServerError,
			"error getting video list", w)
		return
	}

	api.Api.BuildJsonResponse(true, "video list retrieved", l, w)
}


// POST
// Authorization: 	JWT + token
// Params: 			playlistId
// Body: 			None
func VideoListCreate(w http.ResponseWriter, r *http.Request) {
	jwt := auth.ParseApiKey(r, authorizationKey, true)
	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, jwt, w) {
		return
	}

	playlistId := r.URL.Query().Get("playlistId")
	if playlistId == "" {
		api.Api.BuildErrorResponse(http.StatusBadRequest,
			"missing parameter", w)
		return
	}

	l, err := managers.VideoListManagerCreateFromPlaylist(accessToken, playlistId)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusInternalServerError,
			"error creating video list", w)
		return
	}

	api.Api.BuildJsonResponse(true, "video list updated from playlist", l, w)
}

// DELETE
// Authorization: 	JWT + token
// Params: 			None
// Body: 			None
func VideoListDelete(w http.ResponseWriter, r *http.Request) {
	jwt := auth.ParseApiKey(r, authorizationKey, true)
	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, jwt, w) {
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusBadRequest, "invalid body", w)
		return
	}

	var j models.VideoDeleteAllJson
	err = json.Unmarshal(body, &j)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusBadRequest, "invalid json body", w)
		return
	}

	l, err := managers.VideoListManagerDelete(accessToken, j)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusInternalServerError,
			"error deleting video list", w)
		return
	}

	api.Api.BuildJsonResponse(true, "video list deleted", l, w)
}


// PUT
// Authorization: 	JWT + token
// Params: 			None
// Body: 			models.VideoUpdateAllJson
func VideoListUpdate(w http.ResponseWriter, r *http.Request) {
	jwt := auth.ParseApiKey(r, authorizationKey, true)
	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, jwt, w) {
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusBadRequest, "invalid body", w)
		return
	}

	var j models.VideoUpdateAllJson
	err = json.Unmarshal(body, &j)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusBadRequest, "invalid json body", w)
		return
	}

	l, err := managers.VideoListManagerUpdate(accessToken, j)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusInternalServerError,
			"error updating video list", w)
		return
	}

	api.Api.BuildJsonResponse(true, "video list updated", l, w)

}
