-- name: CreatePengguna :one
INSERT INTO pengguna (username, email, password)
VALUES ($1, $2, $3)
RETURNING id;

-- name: GetPasswordPenggunaByEmail :one
SELECT id, password FROM pengguna WHERE email = $1;

-- name: GetPasswordPenggunaByID :one
SELECT password FROM pengguna WHERE id = $1;

-- name: GetPengguna :one
SELECT id, username, email, phone, profile_picture FROM pengguna WHERE id = $1 LIMIT 1;

-- name: UpdatePasswordPengguna :exec
UPDATE pengguna SET password = $2 WHERE id = $1;

-- name: UpdateEmailPengguna :one
UPDATE pengguna SET email = $2 WHERE id = $1 RETURNING *;

-- name: UpdatePhonePengguna :one
UPDATE pengguna SET phone = $2 WHERE id = $1 RETURNING *;

-- name: UpdateProfilePicturePengguna :one
UPDATE pengguna SET profile_picture = $2 WHERE id = $1 RETURNING *;

-- name: CreateKelas :exec
INSERT INTO kelas (nama, tingkat, mata_pelajaran, pembuat, kode_kelas) 
VALUES (  $1, $2, $3, $4, $5)
RETURNING id;

-- name: GetKelasByKode :one
SELECT * FROM kelas WHERE kode_kelas = $1;

-- name: JoinKelas :exec
INSERT INTO anggota_kelas (pengguna_id, kelas_id, role) 
VALUES ($1, $2, $3);

-- name: CekAnggotaKelas :one
SELECT EXISTS(
    SELECT 1 FROM anggota_kelas 
    WHERE pengguna_id = $1 AND kelas_id = $2
);

-- name: LeaveKelas :exec
DELETE FROM anggota_kelas 
WHERE kelas_id = $1 AND pengguna_id = $2;

-- name: ListPostingan :many
SELECT * FROM postingan WHERE tipe = $1 AND pengguna_id = $2 AND kelas_id = $3;

