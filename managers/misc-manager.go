package managers

import (
	"github.com/Dadard29/go-api-utils/log"
	"github.com/Dadard29/go-api-utils/log/logLevel"
	"github.com/Dadard29/go-subscription-connector/subChecker"
	"github.com/Dadard29/go-youtube-dl/api"
)

var logger = log.NewLogger("MANAGER", logLevel.DEBUG)

var scConfig, _ = api.Api.Config.GetSubcategoryFromFile("api", "subscription")
var sc = subChecker.NewSubCheckerFromConfig(scConfig)

const errorNotSubscribed = "invalid subscription token"
