package handler

import (
	"net/http"

	"github.com/gelerum/artpaper/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *Handler) login(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBindBodyWith(&user, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	accessToken, refreshToken, err := h.service.Login(user.Username, user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"accessToken": accessToken, "refreshToken": refreshToken})
}

func (h *Handler) refresh(ctx *gin.Context) {
	accessToken, refreshToken, err := h.service.RefreshTokens(ctx.Param("username"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"accessToken": accessToken, "refreshToken": refreshToken})
}
