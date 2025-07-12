package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/siti-nabila/backend-siti-nabila/internal/domain"
	"github.com/spf13/viper"
)

func RoleAuthorization(userService domain.UserService, allowedRoles ...int) fiber.Handler {
	return func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization")
		jwtSecret := viper.GetString("SECRET_KEY")

		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing token"})
		}

		tokenStr := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token claims"})
		}

		userId := claims["user_id"].(int)

		user, err := userService.GetUserByUserId(userId)
		if err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "failed to get role"})
		}

		// Check apakah role ID user termasuk allowedRoles
		for _, allowed := range allowedRoles {
			if user.RoleId == allowed {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "access denied"})

	}
}
