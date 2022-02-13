package handler

import (
	"net/http"
	"strconv"

	"github.com/gelerum/artpaper/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary Create article
// @Security ApiKeyAuth
// @Description Create article
// @Tags article
// @ID create-article
// @Accept  json
// @Produce  json
// @Param input body model.CreateArticle true "article body, title, owner's username"
// @Success 200 {object} model.Articlename
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Router /article/create [post]
func (h *Handler) createArticle(ctx *gin.Context) {
	var article model.CreateArticle
	err := ctx.ShouldBindBodyWith(&article, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}
	articlename, err := h.service.CreateArticle(article.Title, article.Body, article.Username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, model.Articlename{Articlename: articlename})

}

// @Summary Get article
// @Description Get article by articlename
// @Tags article
// @ID get-article
// @Produce  json
// @Param articlename path string true "Articlename"
// @Success 200 {object} model.GetArticle
// @Failure 400 {object} model.ErrorResponse
// @Router /article/get/{articlename} [get]
func (h *Handler) getArticle(ctx *gin.Context) {
	loadedArticle, exists := ctx.Get("loadedArticle")
	if exists {
		ctx.AbortWithStatusJSON(http.StatusOK, loadedArticle)
		return
	}
	title, body, creationDate, username, err := h.service.GetArticle(ctx.Param("articlename"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}
	article := model.GetArticle{
		Title: title, Body: body, CreationDate: creationDate.Format("2006-01-02"), Username: username,
	}
	ctx.JSON(http.StatusOK, article)
	ctx.Set("uploadedArticle", article)
	ctx.Next()
}

// @Summary Update article
// @Security ApiKeyAuth
// @Description Update article by articlename
// @Tags article
// @ID update-article
// @Accept  json
// @Param articlename path string true "Articlename"
// @Param input body model.UpdateArticle true "article body, title, owner's username"
// @Success 200
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Router /article/update/{articlename} [put]
func (h *Handler) updateArticle(ctx *gin.Context) {
	var article model.UpdateArticle
	err := ctx.ShouldBindBodyWith(&article, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}
	err = h.service.UpdateArticle(ctx.Param("articlename"), article.Title, article.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

// @Summary Delete article
// @Security ApiKeyAuth
// @Description Delete article by articlename
// @Tags article
// @ID delete-article
// @Accept  json
// @Param articlename path string true "Articlename"
// @Success 200
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Router /article/delete/{articlename} [delete]
func (h *Handler) deleteArticle(ctx *gin.Context) {
	err := h.service.DeleteArticle(ctx.Param("articlename"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

// @Summary Find articles
// @Description Find articles by title, author, date from and till
// @Tags article
// @ID find-articles
// @Produce  json
// @Param title query string false "Title to find"
// @Param username query string false "author"
// @Param from query string false "from date"
// @Param to query string false "till date"
// @Param quantity query int false "Quantity"
// @Success 200 {array} model.Article
// @Failure 400 {object} model.ErrorResponse
// @Router /article/find [get]
func (h *Handler) findArticles(ctx *gin.Context) {
	title := ctx.DefaultQuery("title", "")
	username := ctx.DefaultQuery("username", "")
	from := ctx.DefaultQuery("from", "0001-01-01")
	to := ctx.DefaultQuery("to", "9999-12-31")
	quantityString := ctx.DefaultQuery("quantity", "10")
	quantity, err := strconv.Atoi(quantityString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
	}
	articles, err := h.service.GetArticles(title, username, from, to, quantity)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
	}
	ctx.JSON(http.StatusOK, articles)
}
