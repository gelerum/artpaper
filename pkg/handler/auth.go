package handler

import (
	"net/http"

	"github.com/gelerum/artpaper/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary Login
// @Description Login with username and password returning access and refresh tokens
// @ID login
// @Accept  json
// @Produce  json
// @Param input body model.Login true "username and password"
// @Success 200 {object} model.TokensResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /login [post]
func (h *Handler) login(ctx *gin.Context) {
	var user model.Login
	err := ctx.ShouldBindBodyWith(&user, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}
	accessToken, refreshToken, err := h.service.Login(user.Username, user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, model.TokensResponse{AccessToken: accessToken, RefreshToken: refreshToken})
}

// @Summary Refresh
// @Security ApiKeyAuth
// @Description Refresh tokens
// @ID refresh
// @Produce json
// @Success 200 {object} model.TokensResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Router /refresh [post]
func (h *Handler) refresh(ctx *gin.Context) {
	username := ctx.GetString("username")
	accessToken, refreshToken, err := h.service.RefreshTokens(username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, model.TokensResponse{AccessToken: accessToken, RefreshToken: refreshToken})
}
