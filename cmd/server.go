package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"remote-task/infrastructure/persistence/mysql"
	"remote-task/interfaces/http/handler"
	"remote-task/interfaces/http/middleware"
	"remote-task/interfaces/http/routers"
	"strconv"
	"time"
)

func main() {
	err := godotenv.Load("/Users/alibagheri/GolandProjects/remote-task/.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")
	port := os.Getenv("APP_PORT")
	log.Println(port)
	iPort, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}
	repo, err := mysql.NewRepositories(user, password, dbport, host, dbname)
	if err != nil {
		panic(err)
	}
	defer repo.Close()
	GiftCarRepo := mysql.NewGiftCardRepository(repo)
	UserRepo := mysql.NewUserRepository(repo)
	service := handler.NewGiftCartHandler(GiftCarRepo, UserRepo)
	handlers := handler.New(*service)
	router := routers.NewRouter()
	router.Use(middleware.HealthCheck(repo))
	router.Use(middleware.CORS())
	routers.RegisterRoutes(router, handlers)
	routers.NewServer(router, iPort, time.Duration(1)).StartListening()
}
