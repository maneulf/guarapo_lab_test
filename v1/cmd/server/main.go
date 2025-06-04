package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	apserver "github.com/maneulf/guarapo_lab_test/internal/server"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// server up
	s := apserver.New()
	if s == nil {
		log.Print("Unable to start server")
		return

	}
	s.Run()

	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	s.Shutdown(ctx)
}
