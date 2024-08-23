package handler

import (
	"auth-service/pkg/models"
	"auth-service/service"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type AuthHandler interface {
	Register(c *gin.Context)
	LoginEmail(c *gin.Context)
	LoginUsername(c *gin.Context)
}

type authHandler struct {
	srv service.AuthService
	log *slog.Logger
}

func NewAuthHandler(log *slog.Logger, sr service.AuthService) AuthHandler {
	return &authHandler{log: log, srv: sr}
}

// @Summary Register Users
// @Description create users
// @Tags Auth
// @Accept json
// @Produce json
// @Param Register body models.RegisterRequest true "register user"
// @Success 200 {object} models.RegisterResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /register [post]
func (h *authHandler) Register(c *gin.Context) {
	var auth models.RegisterRequest

	if err := c.ShouldBindJSON(&auth); err != nil {
		h.log.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.srv.Register(auth)
	if err != nil {
		h.log.Error("Error occurred while register", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary LoginEmail Users
// @Description sign in user
// @Tags Auth
// @Accept json
// @Produce json
// @Param LoginEmail body models.LoginEmailRequest true "register user"
// @Success 200 {object} models.LoginEmailResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /login/email [post]
func (h *authHandler) LoginEmail(c *gin.Context) {
	var auth models.LoginEmailRequest

	if err := c.ShouldBindJSON(&auth); err != nil {
		h.log.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.srv.LoginEmail(auth)
	if err != nil {
		h.log.Error("Error occurred while login", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("access_token", res.AccessToken, 3600, "", "", false, true)
	c.SetCookie("refresh_token", res.RefreshToken, 3600, "", "", false, true)

	c.JSON(http.StatusOK, res)
}

// @Summary LoginUsername Users
// @Description sign in user
// @Tags Auth
// @Accept json
// @Produce json
// @Param LoginUsername body models.LoginUsernameRequest true "register user"
// @Success 200 {object} models.LoginUsernameResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /login/username [post]
func (h *authHandler) LoginUsername(c *gin.Context) {
	var auth models.LoginUsernameRequest
	if err := c.ShouldBindJSON(&auth); err != nil {
		h.log.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.srv.LoginUsername(auth)
	if err != nil {
		h.log.Error("Error occurred while login", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("access_token", res.AccessToken, 3600, "", "", false, true)
	c.SetCookie("refresh_token", res.RefreshToken, 3600, "", "", false, true)

	c.JSON(http.StatusOK, res)
}
