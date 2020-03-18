package main

import (
	"github.com/Dadard29/go-api-utils/API"
	"github.com/Dadard29/go-api-utils/database"
	"github.com/Dadard29/go-api-utils/service"
	"github.com/Dadard29/go-subscription-connector/subChecker"
	. "github.com/Dadard29/go-youtube-dl/api"
	"github.com/Dadard29/go-youtube-dl/controllers"
	"github.com/Dadard29/go-youtube-dl/models"
	"net/http"
)

var routes = service.RouteMapping{
	"/video": service.Route{
		Description:   "manage the videos objects",
		MethodMapping: service.MethodMapping{
			http.MethodGet: controllers.VideoGet,
			http.MethodPost: controllers.VideoPost,
			http.MethodPut: controllers.VideoPut,
			http.MethodDelete: controllers.VideoDelete,
		},
	},
	"/video/list": service.Route{
		Description:   "manage the list of videos",
		MethodMapping: service.MethodMapping{
			http.MethodGet: controllers.VideoListGet,
			http.MethodPost: controllers.VideoListCreate,
			http.MethodDelete: controllers.VideoListDelete,
			http.MethodPut: controllers.VideoListUpdate,
		},
	},
	"/download": service.Route{
		Description:   "manage download",
		MethodMapping: service.MethodMapping{
			http.MethodGet: controllers.DownloadGet,
			http.MethodPost: controllers.DownloadPost,
			http.MethodDelete: controllers.DownloadDelete,
		},
	},
	"/download/all": service.Route{
		Description:   "manage download all videos",
		MethodMapping: service.MethodMapping{
			http.MethodPost: controllers.DownloadAllPost,
		},
	},
	"/download/file": service.Route{
		Description:   "manage download files",
		MethodMapping: service.MethodMapping{
			http.MethodGet: controllers.FileGet,
		},
	},
}

func main() {
	Api = API.NewAPI("Youtube-Download", "config/config.json", routes, true)

	dbConfig, err := Api.Config.GetSubcategoryFromFile("api", "db")
	Api.Logger.CheckErr(err)
	Api.Database = database.NewConnector(dbConfig, true, []interface{}{
		models.VideoModel{},
	})

	controllers.Sc = subChecker.NewSubChecker(Api.Config.GetEnv("HOST_SUB"))

	Api.Service.Start()

	Api.Service.Stop()
}
