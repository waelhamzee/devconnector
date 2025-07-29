package connection

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

func (c *Controller) CreateConnection(ctx *gin.Context) {
	var req struct {
		UserID   uuid.UUID `json:"user_id"`
		TargetID uuid.UUID `json:"target_id"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		httphelper.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	if err := c.Service.CreateConnection(ctx, req.UserID, req.TargetID); err != nil {
		httphelper.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusCreated)
}

func (c *Controller) ListConnections(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("user_id"))
	if err != nil {
		httphelper.ErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid user_id"))
		return
	}
	conns, err := c.Service.ListConnections(ctx, userID)
	if err != nil {
		httphelper.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, conns)
}

func (c *Controller) DeleteConnection(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("user_id"))
	if err != nil {
		httphelper.ErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid user_id"))
		return
	}
	targetID, err := uuid.Parse(ctx.Param("target_id"))
	if err != nil {
		httphelper.ErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid target_id"))
		return
	}
	if err := c.Service.DeleteConnection(ctx, userID, targetID); err != nil {
		httphelper.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
