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

	sudahGabung, err := queries.CekAnggotaKelas(ctx, db.CekAnggotaKelasParams{
		PenggunaID: userID,
		KelasID:    kelas.ID,
	})
	if err == nil && sudahGabung {
		return errors.New("anda sudah bergabung di kelas ini")
	}

	err = queries.JoinKelas(ctx, db.JoinKelasParams{
		PenggunaID: userID,
		KelasID:    kelas.ID,
		Role:       "murid",
	})

	return err
}

func LeaveKelasLogic(ctx context.Context, queries *db.Queries, userID int32, kodeKelas string) error {
	// 1. Cari dulu ID kelas berdasarkan kode_kelas
	kelas, err := queries.GetKelasByKode(ctx, kodeKelas)
	if err != nil {
		return errors.New("kelas tidak ditemukan")
	}

	// 2. Hapus data keanggotaan dari tabel anggota_kelas
	err = queries.LeaveKelas(ctx, db.LeaveKelasParams{
		KelasID:    kelas.ID,
		PenggunaID: userID,
	})

	return err
}
