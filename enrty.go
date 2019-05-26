package main

import (
	"fmt"
	"log"
	"net/http"
	"persey/dao"
	"persey/routes"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var (
	entryRoute            *routes.EntryRoute
	statisticDao          *dao.StatisticDao
	databaseConfiguration *configuration.Database
	initConfiguration     *configuration.InitPackage
	mapConfigs            []string
	router                *mux.Router
	database              *dao.Database
)

const (
	nameService = "[PERSEY]: "
)

func parsedParams() {
	databaseConfiguration, _ = databaseConfig.Parse(mapConfigs[0])
	initConfig, _ = initConfig.Parse(mapConfigs[1])
}

func configureDaos() {
	statisticDao := &dao.StatisticDao{}
	// statisticDao.New(database)

}

func databasePreparing() {
	err := database.ConnectionUser(databaseConfig)
	if err != nil {
		log.Println(nameService+"Error connected to database: ", err.Error())
		return
	}

	err2 := database.ConnectToBD()
	if err2 != nil {
		log.Println(nameService+"Error connect to database: ", err2.Error())
		return
	}
}

func init() {
	databaseConfig = &configuration.DatabaseConfiguration{}
	initConfig = &configuration.InitPackage{}
	entryRoute = &routes.EntryPointRoute{}
	database = &dao.Database{}

	entryRoute := routes.EntryPointRoute{}
	mapConfigs = []string{
		"./settings/database.json",
		"./settings/init.json",
	}

	router = entryRoute.SettingRouter(tokenDao, userDao, jwtMiddle)
}

func applyCorsPolicy() http.Handler {
	handler := cors.New(
		cors.Options{
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders: []string{"Content-Type", "Accept", "Authorization", "Access-Control-Allow-Origin"},
			AllowedOrigins: []string{"*"},
			Debug:          true,
		},
	).Handler(router)

	return handler
}

func main() {
	if err := http.ListenAndServe(initConfig.ServerAddress+":"+strconv.Itoa(initConfig.ServerPort), routerForListen()); err != nil {
		log.Fatalln(nameService+"Error starting api! err code: ", err.Error())
	}
	fmt.Println(nameService + " was started")
}
