package main

import (
	m "github.com/labstack/echo/v4/middleware"
	"log"
	appConfig "remote-task/config"
	"remote-task/domain/giftCart/repository"
	userRepository "remote-task/domain/user/repository"
	"remote-task/infrastructure/persistence/mysql"
	"remote-task/interfaces/http/handler"
	"remote-task/interfaces/http/middleware"
	"remote-task/interfaces/http/routers"
	migrator "remote-task/migrator"
	"time"
)

func main() {
	appConfig.Init()
	dbCfg := appConfig.Get().Mysql
	appCfg := appConfig.Get().App
	applyMigrations := appCfg.ApplyMigrations
	repo, err := mysql.NewRepositories(dbCfg.Username, dbCfg.Password, dbCfg.Port, dbCfg.Host, dbCfg.Name)
	if err != nil {
		panic(err)
	}
	defer repo.Close()
	if applyMigrations {
		if err := migrator.New(repo).Run("up", 0); err != nil {
			log.Printf("serve migration error: %v\n", err)
		}
	}
	GiftCarRepo := repository.NewGiftCardRepository(repo)
	UserRepo := userRepository.NewUserRepository(repo)
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
