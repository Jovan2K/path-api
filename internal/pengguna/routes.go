package pengguna

import (
	// "path-api/internal/utils"

	// "github.com/golang-jwt/jwt/v5"
	// echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
)

func (h *Handler) Routes(e *echo.Group) {
	pengguna := e.Group("/pengguna")

	pengguna.GET("", h.Get)
	pengguna.PATCH("/email", h.UpdateEmail)
	pengguna.PATCH("/phone", h.UpdatePhone)
	pengguna.PATCH("/profile-picture", h.UpdateProfilePicture)
	pengguna.POST("/reset-password", h.ResetPassword)
}
