package server

import (
	"net/http"
	"synthesizer-server/cmd/config"
	"synthesizer-server/internal/entity"

	"github.com/gofiber/fiber/v2"
)

type Service interface {
	GetSynthesize(str string) (*entity.Data[string], error)
}

func Api(app fiber.Router, s Service, cfg config.Config) {

	app.Post("/synthesize", func(c *fiber.Ctx) error {
		payload := entity.Data[string]{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		synthesize, err := s.GetSynthesize(payload.Data)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}
		c.JSON(synthesize)
		return nil
	})
}
func Default(app fiber.Router) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})
}
