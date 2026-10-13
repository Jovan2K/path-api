package kelas

import (
	"context"
	"errors"
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

		Tingkat:       pgtype.Text{String: req.Tingkat, Valid: req.Tingkat != ""},
		MataPelajaran: pgtype.Text{String: req.MataPelajaran, Valid: req.MataPelajaran != ""},

		Pembuat:   req.Pembuat,
		KodeKelas: req.KodeKelas,
	})

	return err
}

func JoinKelasLogic(ctx context.Context, queries *db.Queries, userID int32, kodeKelas string) error {
	kelas, err := queries.GetKelasByKode(ctx, kodeKelas)
	if err != nil {
		return errors.New("kode kelas tidak ditemukan")
	}

	err = queries.JoinKelas(ctx, db.JoinKelasParams{
		PenggunaID: userID,
		KelasID:    kelas.ID,
		Role:       "murid",
	})

	return err
}
