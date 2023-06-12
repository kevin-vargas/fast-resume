package main

import (
	"fmt"
	"synthesizer-server/cmd/config"
	"synthesizer-server/internal/server"
	"synthesizer-server/internal/service"
	"synthesizer-server/internal/service/cache/inmemory"
	"synthesizer-server/internal/service/synthetizer/openai"

	"github.com/gofiber/fiber/v2/middleware/etag"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.Make()
	c := inmemory.New()
	st := openai.New(cfg.Token, cfg.BaseOpenAI)
	//st := local.New()
	svc := service.New(c, st)
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	app.Use(etag.New())
	server.Default(app)
	server.Api(app, svc, cfg)
	fmt.Printf("Starting server... %+v\n", cfg)
	app.Listen(cfg.Port)
}
