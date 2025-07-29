package post

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	httphelper "github.com/waelhamzee/devconnector/internal/http"
)

type Controller struct {
	Service *Service
}

func NewController(s *Service) *Controller {
	return &Controller{Service: s}
}

func (c *Controller) CreatePost(ctx *gin.Context) {
	var req struct {
		UserID  uuid.UUID `json:"user_id"`
		Content string    `json:"content"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		httphelper.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	post, err := c.Service.CreatePost(ctx, req.UserID, req.Content)
	if err != nil {
		httphelper.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, post)
}

func (c *Controller) GetPost(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		httphelper.ErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid id"))
		return
	}
	post, err := c.Service.GetPostByID(ctx, id)
	if err != nil {
		httphelper.ErrorResponse(ctx, http.StatusNotFound, errors.New("not found"))
		return
	}
	ctx.JSON(http.StatusOK, post)
}

func (c *Controller) ListPosts(ctx *gin.Context) {
	posts, err := c.Service.ListPosts(ctx)
	if err != nil {
		httphelper.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, posts)
}

func (c *Controller) DeletePost(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		httphelper.ErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid id"))
		return
	}
	if err := c.Service.DeletePost(ctx, id); err != nil {
		httphelper.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
