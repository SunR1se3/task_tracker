package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"task_tracker/internal/constants"
)

// Protected protect routes
func Aw() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		ErrorHandler: jwtError,
		SigningKey:   jwtware.SigningKey{Key: []byte("secret")},
		TokenLookup:  "cookie:token",
		AuthScheme:   "Bearer",
		ContextKey:   "currentUser",
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT"})

	} else {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT"})
	}
}

func GetTokenClaims(c *fiber.Ctx) jwt.MapClaims {
	u := c.Locals("currentUser")
	if u != nil {
		claims := u.(*jwt.Token).Claims.(jwt.MapClaims)
		return claims
	}
	return jwt.MapClaims{}
}

func GetUserRole(c *fiber.Ctx) int {
	u := c.Locals("currentUser")
	if u != nil {
		claims := u.(*jwt.Token).Claims.(jwt.MapClaims)
		floatRole, ok := claims["role"].(float64)
		if ok {
			return int(floatRole)
		}
	}
	return -1
}

func GetUserId(c *fiber.Ctx) *uuid.UUID {
	u := c.Locals("currentUser")
	if u != nil {
		claims := u.(*jwt.Token).Claims.(jwt.MapClaims)
		idStr, ok := claims["id"].(string)
		if ok {
			id, err := uuid.Parse(idStr)
			if err == nil {
				return &id
			}
		}
	}
	return nil
}

func IsGranted(c *fiber.Ctx) bool {
	role := GetUserRole(c)
	return role == constants.ROLE_ADMIN
}
