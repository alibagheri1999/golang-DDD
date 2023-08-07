package routers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

func NewServer(router *echo.Echo, port int, gracefulShutdown time.Duration) *Server {
	return &Server{
		addr:             fmt.Sprintf("%s:%v", os.Getenv("HOST_IP"), port),
		router:           router,
		gracefulShutdown: gracefulShutdown,
	}
}

type Server struct {
	addr             string
	router           *echo.Echo
	gracefulShutdown time.Duration
}

func (s *Server) StartListening() {

	go func() {
		if err := s.router.Start(s.addr); err != nil && err != http.ErrServerClosed {
			log.Println("Err server start", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Printf("server shutting down in %s...\n", s.gracefulShutdown)
	c, cancel := context.WithTimeout(context.Background(), s.gracefulShutdown)
	defer cancel()
	if err := s.router.Shutdown(c); err != nil {
		log.Println("Err server shutdown", err)
	}

	<-c.Done()
	log.Println("Good Luck!")
}
