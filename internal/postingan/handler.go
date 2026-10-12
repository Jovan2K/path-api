package postingan

import (
	"path-api/internal/db"
)

type Handler struct {
	db *db.Queries
}

func New(db *db.Queries) *Handler {
	return &Handler{db}
}

type ListPostinganRequest struct {
	Tipe      db.TipePost `query:"tipe" validate:"required,oneof=tugas materi"`
	KelasID   int32  `param:"kelas_id" validate:"required"`
}

