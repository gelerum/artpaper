package handler

import (
	"net/http"
	"strconv"

	"github.com/gelerum/artpaper/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *Handler) createArticle(ctx *gin.Context) {
	var article model.Article
	err := ctx.ShouldBindBodyWith(&article, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	articlename, err := h.service.CreateArticle(article.Title, article.Body, article.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"articlename": articlename})

}

func (h *Handler) getArticle(ctx *gin.Context) {
	loadedArticle, exists := ctx.Get("loadedArticle")
	if exists {
		ctx.AbortWithStatusJSON(http.StatusOK, loadedArticle)
		return
	}
	title, body, creationDate, username, err := h.service.GetArticle(ctx.Param("articlename"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	article := map[string]string{
		"title": title, "body": body, "creationDate": creationDate.Format("2006-01-02"), "username": username,
	}
	ctx.JSON(http.StatusOK, article)
	ctx.Set("uploadedArticle", article)
	ctx.Next()
}

func (h *Handler) updateArticle(ctx *gin.Context) {
	var article model.Article
	err := ctx.ShouldBindBodyWith(&article, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	err = h.service.UpdateArticle(ctx.Param("articlename"), article.Title, article.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	ctx.Status(http.StatusOK)
}

func (h *Handler) deleteArticle(ctx *gin.Context) {
	err := h.service.DeleteArticle(ctx.Param("articlename"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	ctx.Status(http.StatusOK)
}

func (h *Handler) getArticles(ctx *gin.Context) {
	title := ctx.Query("title")
	username := ctx.DefaultQuery("username", "")
	from := ctx.DefaultQuery("from", "0001-01-01")
	to := ctx.DefaultQuery("to", "9999-12-31")
	quantityString := ctx.DefaultQuery("quantity", "10")
	quantity, err := strconv.Atoi(quantityString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	articles, err := h.service.GetArticles(title, username, from, to, quantity)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}
	ctx.JSON(http.StatusOK, gin.H{"articles": articles})
}
