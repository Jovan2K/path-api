package kelas

import (
	"encoding/json"
	"net/http"
	"path-api/internal/db"
	"github.com/golang-jwt/jwt/v5"
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

	// --- MULAI KODE BARU: Ambil ID dari JWT ---
	// "user" adalah key default dari Echo JWT Middleware
	userToken, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Token tidak ditemukan")
	}

	// Ini mencegah error akibat bentrok tipe data atau versi JWT
	claimsJSON, _ := json.Marshal(userToken.Claims)
	var claims map[string]interface{}
	json.Unmarshal(claimsJSON, &claims)

	// Ambil ID menggunakan key "user_id" (sesuai format token aslimu)
	idClaim, ok := claims["user_id"].(float64)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Key 'user_id' tidak ditemukan di token")
	}

	// Timpa req.Pembuat dengan ID asli dari token
	req.Pembuat = int32(idClaim)
	// --- SELESAI KODE BARU ---
	// Validasi kolom wajib (req.Pembuat sudah pasti terisi dari token)
	if req.Nama == "" || req.KodeKelas == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Nama kelas dan kode kelas wajib diisi")
	}

	err := BuatKelasLogic(c.Request().Context(), h.db, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Gagal membuat kelas: "+err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Kelas berhasil dibuat",
	})
}
