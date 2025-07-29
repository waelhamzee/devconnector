package comment

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

func (c *Controller) CreateComment(ctx *gin.Context) {
	var req struct {
		PostID  uuid.UUID `json:"post_id"`
		UserID  uuid.UUID `json:"user_id"`
		Content string    `json:"content"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		httphelper.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	comment, err := c.Service.CreateComment(ctx, req.PostID, req.UserID, req.Content)
	if err != nil {
		httphelper.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, comment)
}

func (c *Controller) ListCommentsByPost(ctx *gin.Context) {
	postID, err := uuid.Parse(ctx.Param("post_id"))
	if err != nil {
		httphelper.ErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid post_id"))
		return
	}
	comments, err := c.Service.ListCommentsByPost(ctx, postID)
	if err != nil {
		httphelper.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, comments)
}

func (c *Controller) DeleteComment(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		httphelper.ErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid id"))
		return
	}
	if err := c.Service.DeleteComment(ctx, id); err != nil {
		httphelper.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
