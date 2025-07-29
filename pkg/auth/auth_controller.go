package auth

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	httphelper "github.com/waelhamzee/devconnector/internal/http"
)

type AuthController struct {
	Service *AuthService
}

func NewAuthController(service *AuthService) *AuthController {
	return &AuthController{Service: service}
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (ctl *AuthController) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httphelper.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	user, err := ctl.Service.Register(req.Name, req.Email, req.Password)
	if err != nil {
		httphelper.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	token, err := generateJWT(user.ID)
	if err != nil {
		httphelper.ErrorResponse(c, http.StatusInternalServerError, errors.New("could not generate token"))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"token": token})
}

func (ctl *AuthController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httphelper.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	user, err := ctl.Service.Authenticate(req.Email, req.Password)
	if err != nil {
		httphelper.ErrorResponse(c, http.StatusUnauthorized, errors.New("invalid credentials"))
		return
	}
	token, err := generateJWT(user.ID)
	if err != nil {
		httphelper.ErrorResponse(c, http.StatusInternalServerError, errors.New("could not generate token"))
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func generateJWT(userID uuid.UUID) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"sub": userID.String(),
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
