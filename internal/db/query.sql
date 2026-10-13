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

<<<<<<< HEAD
-- name: CreateKelas :one
INSERT INTO kelas (nama, tingkat, mata_pelajaran, pembuat, kode_kelas)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;
=======
-- name: ListPostingan :many
SELECT * FROM postingan WHERE tipe = $1 AND pengguna_id = $2 AND kelas_id = $3;

>>>>>>> 44dbf50dc6e3bce119655133a0df12bfc3d6ab19
