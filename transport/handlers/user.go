package handlers

import (
	"archi/config"
	"archi/model"
	"archi/service"
	"archi/storage"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type UserHandler struct {
	storage     *storage.Storage
	userManager *service.Service
	config      *config.Config
	log         *zap.Logger
}

var Users []model.User

func NewUserHandler(logger *zap.Logger, config *config.Config, storage *storage.Storage, userManager *service.Service) *UserHandler {
	return &UserHandler{log: logger, config: config, storage: storage, userManager: userManager}
}

func (h *UserHandler) SignIn(c echo.Context) error {
	return c.Render(http.StatusOK, "sign_in.html", map[string]interface{}{
		"IsError": false,
	})
}

func (h *UserHandler) Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	if searchByUsername(email, password) {
		cookie := new(http.Cookie)
		cookie.Name = "token"
		cookie.Value = "token"
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)
		return c.Redirect(http.StatusMovedPermanently, "/")
	}
	return c.Redirect(http.StatusMovedPermanently, "/login")
}

func (h *UserHandler) LoginPage(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", nil)
}

func (h *UserHandler) Logout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Expires = time.Now().Add(6 * time.Hour)
	c.SetCookie(cookie)

	return c.Redirect(http.StatusFound, "/")
}

func (h *Handlers) Home(c echo.Context) error {
	fmt.Println(23)
	token, _ := ReadCookie(c)
	fmt.Println(22)
	isAuth := false
	if token == "token" {
		isAuth = true
	}
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"IsAuthenticated": isAuth,
	})
}

func (h *UserHandler) Register(c echo.Context) error {
	fmt.Println(11)
	email := c.FormValue("email")
	username := c.FormValue("username")
	password := c.FormValue("password")

	fmt.Println(email, username, password)
	Users = append(Users, model.User{
		Name:     username,
		Email:    email,
		Password: password,
	})
	return c.Redirect(http.StatusMovedPermanently, "/")
}

func (h *UserHandler) Registration(c echo.Context) error {
	return c.Render(http.StatusOK, "register.html", nil)
}

func searchByUsername(email, password string) bool {
	for _, user := range Users {
		if user.Email == email && user.Password == password {
			return true
		}
	}
	return false
}
