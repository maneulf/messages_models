package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	apserver "github.com/maneulf/base_ms/pkg/server"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	s := apserver.New()

	apiv1 := s.Group("/api/v1")

	apiv1.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	s.Run()

	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	s.Shutdown(ctx)
}
