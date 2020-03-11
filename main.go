package main

import (
	"github.com/Dadard29/go-api-utils/API"
	"github.com/Dadard29/go-api-utils/database"
	"github.com/Dadard29/go-api-utils/service"
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
}

func main() {
	Api = API.NewAPI("Youtube-Download", "config/config.json", routes, true)

	dbConfig, err := Api.Config.GetSubcategoryFromFile("api", "db")
	Api.Logger.CheckErr(err)
	Api.Database = database.NewConnector(dbConfig, true, []interface{}{
		models.VideoModel{},
	})

	Api.Service.Start()

	Api.Service.Stop()
}
