package kelas

import (
	"net/http"
	"path-api/internal/db"
	"github.com/labstack/echo/v5"
)

type Handler struct {
	db *db.Queries
}

func New(db *db.Queries) *Handler {
	return &Handler{db}
}

type CreateKelasRequest struct {
	Nama          string `json:"nama"`
	Tingkat       string `json:"tingkat"`
	MataPelajaran string `json:"mata_pelajaran"`
	Pembuat       int32  `json:"pembuat"`
	KodeKelas     string `json:"kode_kelas"`
}

func (h *Handler) Create(c *echo.Context) error {
	var req CreateKelasRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Format input JSON tidak valid")
	}

	// Validasi kolom wajib
	if req.Nama == "" || req.KodeKelas == "" || req.Pembuat == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Nama kelas, kode kelas, dan ID pembuat wajib diisi")
	}

	// Panggil logic layer dengan mengirimkan h.db
	err := BuatKelasLogic(c.Request().Context(), h.db, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Gagal membuat kelas: "+err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Kelas berhasil dibuat",
	})
}