package controllers

import (
	"github.com/Dadard29/go-api-utils/log"
	"github.com/Dadard29/go-api-utils/log/logLevel"
	"github.com/Dadard29/go-subscription-connector/subChecker"
	"github.com/Dadard29/go-youtube-dl/api"
	"net/http"
)

var Sc *subChecker.SubChecker

var logger = log.NewLogger("CONTROLLER", logLevel.DEBUG)

const errorNotSubscribed = "invalid subscription token"
const authorizationKey = "Authorization"
const accessTokenKey = "X-Access-Token"
const apiName = "youtube-download"

func checkToken(token string, w http.ResponseWriter) bool {
	msg, err := Sc.CheckToken(token, apiName)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusInternalServerError, msg, w)
		return false
	}

	return true
}
