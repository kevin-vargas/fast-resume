package main

import (
	"bff-server/cmd/config"
	"bff-server/internal/repository/slack"
	"bff-server/internal/server"
	"bff-server/internal/service"
	"fmt"

	"github.com/gofiber/fiber/v2/middleware/etag"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.Make()

	//repo := rest.New(cfg.URIComponentService, cfg.TTLService, cfg.Timeout)
	//repo := local.New()
	repo := slack.New(cfg.SlackAPIUri)
	svc := service.New(repo, cfg.BaseSynthesizerServer)
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	app.Use(etag.New())
	server.Default(app)
	api := app.Group(cfg.APIPrefix)
	server.Api(api, svc, cfg)
	fmt.Printf("Starting server... %+v\n", cfg)
	app.Listen(cfg.Port)
}
