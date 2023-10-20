package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/pawarpranav83/golang-chat/db/sqlc"
	"github.com/pawarpranav83/golang-chat/token"
)

type roomcreationRequest struct {
	Roomname string `json:"roomname" binding:"required"`
}

// type roomcreationUserRequest struct {
// 	UserID int64 `uri:"user_id" binding:"required"`
// }

func (server *Server) createRoomUser(ctx *gin.Context) {
	var reqRoomCreation roomcreationRequest
	// var reqUserId roomcreationUserRequest

	if err := ctx.ShouldBindJSON(&reqRoomCreation); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Doubt - Why *token.Payload works and not token.Payload
	payload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	// if err := ctx.ShouldBindUri(&reqUserId); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	// 	return
	// }

	arg := db.RoomCreationTxParams{
		Roomname: reqRoomCreation.Roomname,
		UserID:   payload.ID,
	}

	result, err := server.store.RoomCreationTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// Get all users in the room
type getRoomRequest struct {
	RoomId int64 `uri:"roomId" binding:"required,min=1"`
}

func (server *Server) getRoomUser(ctx *gin.Context) {
	var req getRoomRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	rooms, err := server.store.GetRoomusers(ctx, req.RoomId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println(rooms)
	ctx.JSON(http.StatusOK, rooms)
}

type addUserRequest struct {
	RoomId int64 `uri:"roomId" binding:"required,min=1"`
	UserId int64 `uri:"userId" binding:"required,min=1"`
}

func (server *Server) addUsertoRoom(ctx *gin.Context) {
	var req addUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	currId := ctx.MustGet(authorizationPayloadKey).(*token.Payload).ID
	// If current user is not in the room, then, he is not allowed to add others.
	_, err := server.store.GetRoomuser(ctx, db.GetRoomuserParams{RoomID: req.RoomId, UserID: currId})
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	arg := db.AddUsertoRoomParams{
		RoomID: req.RoomId,
		UserID: req.UserId,
	}

	res, err := server.store.AddUsertoRoom(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

type userLeaveRequest struct {
	RoomId int64 `uri:"roomId" binding:"required,min=1"`
}

func (server *Server) userLeaveRoom(ctx *gin.Context) {
	var req userLeaveRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	currId := ctx.MustGet(authorizationPayloadKey).(*token.Payload).ID

	err := server.store.DeleteUserfromRoom(ctx, db.DeleteUserfromRoomParams{RoomID: req.RoomId, UserID: currId})
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
	}

	// Room is deleted if the last user leaves.
	if usersInRoom, _ := server.store.GetRoomusers(ctx, req.RoomId); usersInRoom == nil {
		err := server.store.Deleteroom(ctx, req.RoomId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
	}

	ctx.JSON(http.StatusOK, "success")
}

func (server *Server) getRoomsofUser(ctx *gin.Context) {
	userId := ctx.MustGet(authorizationPayloadKey).(*token.Payload).ID
	rooms, err := server.store.GetUserRooms(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, rooms)
}
