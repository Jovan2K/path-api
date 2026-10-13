package kelas

import (
	"context"
	"net/http"
	"path-api/internal/db"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v5"
)

func (h *Handler) List(c *echo.Context) error {
	return echo.NewHTTPError(http.StatusBadRequest, "tidak ada.")
}
func BuatKelasLogic(ctx context.Context, queries *db.Queries, req CreateKelasRequest) error {
	err := queries.CreateKelas(ctx, db.CreateKelasParams{
		Nama: req.Nama,

		// Ubah dua baris di bawah ini:
		Tingkat:       pgtype.Text{String: req.Tingkat, Valid: req.Tingkat != ""},
		MataPelajaran: pgtype.Text{String: req.MataPelajaran, Valid: req.MataPelajaran != ""},

		Pembuat:   req.Pembuat,
		KodeKelas: req.KodeKelas,
	})

	return err
}
