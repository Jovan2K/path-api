package postingan

import "github.com/labstack/echo/v5"

func (h *Handler) Routes(e *echo.Group) {
	postingan := e.Group("/kelas/:kelasID")

	postingan.GET("/tugas", h.ListTugas)
	postingan.GET("/materi", h.ListMateri)
}


