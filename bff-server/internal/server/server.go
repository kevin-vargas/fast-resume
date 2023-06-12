package server

import (
	"bff-server/cmd/config"
	"bff-server/internal/entity"
	"bff-server/internal/middleware"
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Default(app fiber.Router) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})
}

type Service interface {
	GetChannels(context.Context) (*entity.Channels, error)
	SummarizeChannel(context.Context, string) (*entity.Data[string], error)
	SummarizeIAChannel(context.Context, string) (*entity.Data[string], error)
}

func Api(app fiber.Router, s Service, cfg config.Config) {

	// TODO: from config
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     cfg.AllowOrigins,
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))

	resources := app.Group("/", middleware.NewAuth())
	resources.Get("/channels", func(c *fiber.Ctx) error {
		apis, err := s.GetChannels(c.Context())
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}
		c.JSON(apis)
		return nil
	})
	resources.Get("/summarize/:id", func(c *fiber.Ctx) error {
		summarize, err := s.SummarizeChannel(c.Context(), c.Params("id"))
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}
		c.JSON(summarize)
		return nil
	})
	resources.Get("/summarize-ia/:id", func(c *fiber.Ctx) error {
		summarize, err := s.SummarizeIAChannel(c.Context(), c.Params("id"))
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}
		c.JSON(summarize)
		return nil
	})
}
