
PROJECT NOTES - 

In db/query, the comments on top instructs sqlc to generate go code for that query.

For testing we use testify package instead of doing if else statements.

NOTE - Each unit test function should be independent, like for get user, we should create user first, in the get user test func, instead of using create user test func.

DBTX Interface in db.go - Doubt.

Rewatch #6 lecture.
Rewatch #19 for Paseto working.

Changes in DB without taking down the whole db - #15

Doubt - Background Ctx.

Gin Docs Ref - https://github.com/gin-gonic/gin/blob/v1.9.1/docs/doc.md
Validator Docs Ref - https://pkg.go.dev/github.com/go-playground/validator
Viper Docs Ref - https://github.com/spf13/viper
Mapstructure Docs Ref - https://github.com/mitchellh/mapstructure
Pasteo Docs Ref - https://github.com/o1egl/paseto

To change the env vars (in app.env) in golang, just specify the env var and its value before the go run command - SERVER_ADDRESS=0.0.0.0:8081 make server


App Working - 
User should delete/leave all rooms that he is in, before account deletion.
Room is deleted if it has no users left, that is, when last user is leaving the room is deleted.
Note - Cannot delete rooms / users until it has entries in userroom table.


Remaining - 
    
    Make roomid as int64 is ws.go
    Before a user connects to a room with ws, we check whether that user is a member of that room.
    Use IN operation to find the usernames of users in a particular room - use db middleware.
