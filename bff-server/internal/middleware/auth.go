package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var Token = struct{}{}

func GetToken(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(Token).(string)
	return val, ok
}

type tokenClaim struct {
	Token string `json:"token,omitempty"`
	jwt.RegisteredClaims
}

func NewAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		t := c.Cookies("token")
		if t == "" {
			return c.Status(http.StatusUnauthorized).SendString("missing token")
		}
		// TODO: validate sign jwks
		claim := new(tokenClaim)
		jwtParsed, _, err := new(jwt.Parser).ParseUnverified(
			t,
			claim,
		)
		if err != nil {
			fmt.Println(err.Error())
			return c.Status(http.StatusUnauthorized).SendString("invalid token")
		}
		claims, ok := jwtParsed.Claims.(*tokenClaim)
		if !ok {
			fmt.Println(err.Error())
			return c.Status(http.StatusUnauthorized).SendString("invalid token")
		}
		c.Locals(Token, claims.Token)
		c.Next()
		return nil
	}
}
