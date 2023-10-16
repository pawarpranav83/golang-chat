package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/pawarpranav83/golang-chat/db/sqlc"
)

// Docs for package validator - https://pkg.go.dev/github.com/go-playground/validator
type createUserRequest struct {
	// binding is used to use the default validator provided by gin to validate teh data
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// The ctx provides various methods to read input params, and to write out responses
func (server *Server) createUser(ctx *gin.Context) {
	// gin uses validator package internally to perform data validation, docs link above
	var req createUserRequest

	// ShouldBindWith binds the passed struct pointer using the specified binding engine.
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// Used to send a JSON response, first arg is the status code
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type getUserRequest struct {
	// uri: id, is validation for the url param, https://github.com/gin-gonic/gin/blob/v1.9.1/docs/doc.md#model-binding-and-validation
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.ID)
	if err != nil {
		// when input id doesn't exist
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (server *Server) listUsers(ctx *gin.Context) {
	users, err := server.store.ListUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, users)
}

type deleteUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// NOT WORKING
func (server *Server) deleteUser(ctx *gin.Context) {
	var req deleteUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteUser(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "success")
}

type getUpdateUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}
type updateUserRequest struct {
	Username string `json:"username" binding:"required,min=1"`
}

func (server *Server) updateUser(ctx *gin.Context) {
	var reqGet getUpdateUserRequest
	var reqUpdate updateUserRequest

	if err := ctx.ShouldBindUri(&reqGet); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&reqUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateUserParams{
		ID:       reqGet.ID,
		Username: reqUpdate.Username,
	}

	user, err := server.store.UpdateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
