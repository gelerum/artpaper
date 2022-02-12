package handler

import (
	"github.com/gelerum/artpaper/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) (handler *Handler) {
	handler = &Handler{
		service: service,
	}
	return
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	router.POST("/login", h.login)
	router.POST("/refresh/:username", h.AuthRefreshTokenExists, h.AuthUsernameParam, h.refresh)

	article := router.Group("/article")
	{
		article.POST("/create", h.AuthValidToken, h.AuthArticleBodyUsername, h.createArticle)
		article.GET("/get/:articlename", h.LoadArticleCache, h.getArticle, h.UploadArticleCache)
		article.PUT("/update/:articlename", h.AuthValidToken, h.AuthOwnership, h.updateArticle)
		article.DELETE("/delete/:articlename", h.AuthValidToken, h.AuthOwnership, h.deleteArticle)
		article.GET("/find", h.getArticles)
	}
	user := router.Group("/user")
	{
		user.POST("/create", h.createUser)
		user.GET("/get/:username", h.LoadUserCache, h.getUser, h.UploadUserCache)
		user.PUT("/update/:username", h.AuthValidToken, h.AuthUsernameParam, h.updateUser)
		user.DELETE("/delete/:username", h.AuthValidToken, h.AuthUsernameParam, h.deleteUser)
		user.GET("/find", h.findUsers)
	}
	return router
}
