package handler

import (
	"net/http"
	"strconv"

	"github.com/gelerum/artpaper/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary Create user
// @Description Create user
// @Tags user
// @ID create-user
// @Param input body model.User true "User account data"
// @Success 201
// @Failure 400 {object} model.ErrorResponse
// @Router /user/create [post]
func (h *Handler) createUser(ctx *gin.Context) {
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}
	err = h.service.CreateUser(user.Username, user.Name, user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.Status(http.StatusCreated)
}

// @Summary Get user
// @Description Get user by username
// @Tags user
// @ID get-user
// @Produce  json
// @Param username path string true "Username"
// @Success 200 {object} model.GetUser
// @Failure 400 {object} model.ErrorResponse
// @Router /user/get/{username} [get]
func (h *Handler) getUser(ctx *gin.Context) {
	loadedUser, exists := ctx.Get("loadedUser")
	if exists {
		ctx.AbortWithStatusJSON(http.StatusOK, loadedUser)
		return
	}
	name, err := h.service.GetUser(ctx.Param("username"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, model.GetUser{Name: name})
	ctx.Set("uploadedUser", model.GetUser{Name: name})
	ctx.Next()
}

// @Summary Update user
// @Security ApiKeyAuth
// @Description Update user by username
// @Tags user
// @ID update-user
// @Accept  json
// @Param username path string true "Username"
// @Param input body model.UpdateUser true "Updated data"
// @Success 200
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Router /user/update/{username} [put]
func (h *Handler) updateUser(ctx *gin.Context) {
	var user model.UpdateUser
	err := ctx.ShouldBindBodyWith(&user, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}
	err = h.service.UpdateUser(ctx.Param("username"), user.Name, user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

// @Summary Delete user
// @Security ApiKeyAuth
// @Description Delete user by username
// @Tags user
// @ID delete-user
// @Param username path string true "Username"
// @Success 200
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Router /user/delete/{username} [delete]
func (h *Handler) deleteUser(ctx *gin.Context) {
	err := h.service.DeleteUser(ctx.Param("username"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

// @Summary Find users
// @Description Find users by username and name
// @Tags user
// @ID find-usera
// @Produce  json
// @Param pattern query string true "Pattern to find"
// @Param quantity query int false "Quantity"
// @Success 200 {array} model.FindUsers
// @Failure 400 {object} model.ErrorResponse
// @Router /user/find [get]
func (h *Handler) findUsers(ctx *gin.Context) {
	pattern := ctx.Query("pattern")
	if pattern == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "empty pattern key"})
	}
	quantityString := ctx.DefaultQuery("quantity", "5")
	quantity, err := strconv.Atoi(quantityString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
	}
	users, err := h.service.FindUsers(pattern, quantity)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
	}
	ctx.JSON(http.StatusOK, users)
}
