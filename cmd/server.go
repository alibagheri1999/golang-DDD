package main

import (
	m "github.com/labstack/echo/v4/middleware"
	appConfig "remote-task/config"
	"remote-task/domain/giftCart/repository"
	repository2 "remote-task/domain/user/repository"
	"remote-task/infrastructure/persistence/mysql"
	"remote-task/interfaces/http/handler"
	"remote-task/interfaces/http/middleware"
	"remote-task/interfaces/http/routers"
	"time"
)

func main() {
	appConfig.Init()
	dbCfg := appConfig.Get().Mysql
	appCfg := appConfig.Get().App
	repo, err := mysql.NewRepositories(dbCfg.Username, dbCfg.Password, dbCfg.Port, dbCfg.Host, dbCfg.Name)
	if err != nil {
		panic(err)
	}
	defer repo.Close()
	GiftCarRepo := repository.NewGiftCardRepository(repo)
	UserRepo := repository2.NewUserRepository(repo)
	giftCartService := handler.NewGiftCartHandler(GiftCarRepo, UserRepo)
	generalService := handler.NewGeneralHandler()
	handlers := handler.New(*giftCartService, *generalService)
	router := routers.NewRouter()
	router.Use(middleware.HealthCheck(repo))
	router.Use(m.CORS())
	router.Use(middleware.CORS())

	routers.RegisterRoutes(router, handlers)
	routers.NewServer(router, appCfg.Port, time.Duration(1)).StartListening()
}
