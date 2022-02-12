package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/gelerum/artpaper/pkg/model"
	"github.com/gelerum/artpaper/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *Handler) AuthValidToken(ctx *gin.Context) {
	accessToken := strings.Split(ctx.Request.Header["Authorization"][0], " ")[1]
	username, valid := service.ValidateToken(accessToken, "access")
	if !valid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Ivalid token"})
		return
	}
	ctx.Set("username", username)
	ctx.Next()
}

func (h *Handler) AuthUsernameParam(ctx *gin.Context) {
	if ctx.Param("username") != ctx.GetString("username") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Doesn't have access"})
		return
	}
	ctx.Next()
}

func (h *Handler) AuthOwnership(ctx *gin.Context) {
	_, _, _, username, err := h.service.GetArticle(ctx.Param("articlename"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	if username != ctx.GetString("username") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Doesn't have access"})
		return
	}
	ctx.Next()
}

func (h *Handler) AuthRefreshTokenExists(ctx *gin.Context) {
	refreshToken := strings.Split(ctx.Request.Header["Authorization"][0], " ")[1]
	username, err := h.service.GetDelRefreshToken(refreshToken)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err})
		return
	}
	ctx.Set("username", username)
	ctx.Next()
}

func (h *Handler) AuthArticleBodyUsername(ctx *gin.Context) {
	var article model.Article
	err := ctx.ShouldBindBodyWith(&article, binding.JSON)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err})
		return
	}
	if article.Username != ctx.GetString("username") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Doesn't have access"})
		return
	}
}

func (h *Handler) LoadArticleCache(ctx *gin.Context) {
	title, body, creationDate, username, err := h.service.LoadArticleCache("/article/get/" + ctx.Param("articlename"))
	if err != nil || title == "" {
		ctx.Next()
		return
	}
	article := map[string]string{
		"title": title, "body": body, "creationDate": creationDate, "username": username,
	}
	ctx.Set("loadedArticle", article)
	ctx.Next()
}

func (h *Handler) UploadArticleCache(ctx *gin.Context) {
	value, exists := ctx.Get("uploadedArticle")
	if !exists {
		log.Println("Handler doesn't upload article to middleware")
		return
	}
	article := value.(map[string]string)
	h.service.UploadArticleCache("/article/get/"+ctx.Param("articlename"), article["title"], article["body"], article["creationDate"], article["username"])
}

func (h *Handler) LoadUserCache(ctx *gin.Context) {
	name, err := h.service.LoadUserCache("/user/get/" + ctx.Param("username"))
	if err != nil || name == "" {
		ctx.Next()
		return
	}
	user := map[string]string{
		"name": name,
	}
	ctx.Set("loadedUser", user)
	ctx.Next()
}

func (h *Handler) UploadUserCache(ctx *gin.Context) {
	value, exists := ctx.Get("uploadedUser")
	if !exists {
		log.Println("Handler doesn't upload article to middleware")
		return
	}
	user := value.(map[string]string)
	h.service.UploadUserCache("/user/get/"+ctx.Param("username"), user["name"])
}
