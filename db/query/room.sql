-- name: CreateRoom :one
INSERT INTO rooms (
    roomname
) VALUES (
    $1
) RETURNING *;

-- name: GetRoom :one
SELECT * FROM rooms
WHERE id = $1 LIMIT 1;

-- name: ListRooms :many
SELECT * FROM rooms
ORDER BY id;

-- name: Deleteroom :exec
DELETE FROM rooms WHERE id = $1;

-- name: DeleteRoombyRoomname :exec
DELETE FROM rooms WHERE roomname = $1;