package postingan

import "github.com/labstack/echo/v5"

func (h *Handler) Routes(e *echo.Group) {
	postingan := e.Group("/postingan")

	postingan.GET("", h.List)
}
