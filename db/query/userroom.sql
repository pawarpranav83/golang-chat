-- name: AddUsertoRoom :one
INSERT INTO userroom (
    room_id,
    user_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetRoomusers :many
SELECT * FROM userroom
WHERE room_id = $1;

-- name: GetUserRooms :many
SELECT * FROM userroom
WHERE user_id = $1;

-- name: GetRoomuser :one
SELECT * FROM userroom
WHERE room_id = $1 AND user_id = $2;

-- name: DeleteUserfromRoom :exec
DELETE FROM userroom WHERE room_id = $1 AND user_id = $2;