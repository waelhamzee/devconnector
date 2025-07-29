package user

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	httphelper "github.com/waelhamzee/devconnector/internal/http"
)

type UserController struct {
	service *UserService
}

func NewUserController(service *UserService) *UserController {
	return &UserController{service: service}
}

func (ctl *UserController) GetUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		httphelper.ErrorResponse(c, http.StatusBadRequest, errors.New("invalid id"))
		return
	}
	user, err := ctl.service.GetUserByID(id)
	if err != nil {
		httphelper.ErrorResponse(c, http.StatusNotFound, errors.New("user not found"))
		return
	}
	c.JSON(http.StatusOK, user)
}
