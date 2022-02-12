package handler

import (
	"net/http"
	"strconv"

	"github.com/gelerum/artpaper/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *Handler) createUser(ctx *gin.Context) {
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	err = h.service.CreateUser(user.Username, user.Name, user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	ctx.Status(http.StatusCreated)
}

func (h *Handler) getUser(ctx *gin.Context) {
	loadedUser, exists := ctx.Get("loadedUser")
	if exists {
		ctx.AbortWithStatusJSON(http.StatusOK, loadedUser)
		return
	}
	name, err := h.service.GetUser(ctx.Param("username"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	user := map[string]string{
		"name": name,
	}
	ctx.JSON(http.StatusOK, user)
	ctx.Set("uploadedUser", user)
	ctx.Next()
}

func (h *Handler) updateUser(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBindBodyWith(&user, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	err = h.service.UpdateUser(ctx.Param("username"), user.Name, user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	ctx.Status(http.StatusOK)
}

func (h *Handler) deleteUser(ctx *gin.Context) {
	err := h.service.DeleteUser(ctx.Param("username"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	ctx.Status(http.StatusOK)
}

func (h *Handler) findUsers(ctx *gin.Context) {
	pattern := ctx.Query("pattern")
	if pattern == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "empty pattern key"})
	}
	quantityString := ctx.DefaultQuery("quantity", "5")
	quantity, err := strconv.Atoi(quantityString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	users, err := h.service.FindUsers(pattern, quantity)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}
