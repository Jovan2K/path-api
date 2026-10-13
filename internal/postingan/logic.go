package postingan

import (
	"net/http"
	"path-api/internal/db"
	"path-api/internal/utils"

	"github.com/labstack/echo/v5"
)

func (h *Handler) ListPostingan(c *echo.Context) error {
	req := new(ListPostinganRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return utils.ValidationErrorHandler(c, err)
	}
	claims, err := utils.GetClaims(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Tidak dapat mengakses data klaim")
	}
	tugas, err := h.db.ListPostingan(c.Request().Context(), db.ListPostinganParams{
		PenggunaID: int32(claims.UserID),
		KelasID:    req.KelasID,
		Tipe:       req.Tipe,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Gagal mengambil data "+string(req.Tipe))
	}
	return c.JSON(http.StatusOK, tugas)
}
