package kelas

import (
	"net/http"
	"path-api/internal/db"
	"context"
	"github.com/labstack/echo/v5"
)

func (h *Handler) List(c *echo.Context) error {
	return echo.NewHTTPError(http.StatusBadRequest, "tidak ada.")
}
func BuatKelasLogic(ctx context.Context, queries *db.Queries, req CreateKelasRequest) error {
	
	// Catatan: Jika Tingkat dan MataPelajaran bersifat opsional di PostgreSQL, 
	// sqlc biasanya men-generate tipenya menjadi pgtype.Text atau sql.NullString.
	// Sesuaikan tipe datanya dengan file query.sql.go milikmu.
	
	/*
	_, err := queries.CreateKelas(ctx, db.CreateKelasParams{
		Nama:          req.Nama,
		Tingkat:       req.Tingkat,       // Ubah formatnya jika sqlc menggunakan pgtype.Text
		MataPelajaran: req.MataPelajaran, // Ubah formatnya jika sqlc menggunakan pgtype.Text
		Pembuat:       req.Pembuat,
		KodeKelas:     req.KodeKelas,
	})
	
	return err
	*/

	return nil // Hapus baris ini setelah blok komentar di atas diaktifkan dan disesuaikan
}
