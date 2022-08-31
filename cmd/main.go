package main

import (
	"github.com/rasul07/books_api_gateway/api"
	_ "github.com/rasul07/books_api_gateway/api/docs"
	"github.com/rasul07/books_api_gateway/config"
	"github.com/rasul07/books_api_gateway/pkg/logger"
	"github.com/rasul07/books_api_gateway/services"
	// "github.com/gomodule/redigo/redis"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "books_api_gateway")

	gprcClients, _ := services.NewGrpcClients(&cfg)

	server := api.New(&api.RouterOptions{
		Log:      log,
		Cfg:      cfg,
		Services: gprcClients,
	})

	server.Run(cfg.HttpPort)

}
