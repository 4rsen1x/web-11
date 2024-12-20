package api

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func (srv *Server) Register(c echo.Context) error {
	input := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, "invalid input")
	}

	err := srv.uc.RegisterUser(input.Username, input.Password)
	if err != nil {
		return c.String(http.StatusConflict, err.Error())
	}

	return c.String(http.StatusCreated, "user registered successfully")
}

func (srv *Server) Login(c echo.Context) error {
	input := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, "invalid input")
	}

	isValid, err := srv.uc.AuthenticateUser(input.Username, input.Password)
	if err != nil || !isValid {
		return c.String(http.StatusUnauthorized, "invalid username or password")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = input.Username
	claims["exp"] = time.Now().Add(72 * time.Hour).Unix()

	t, err := token.SignedString([]byte(srv.secretKey))
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to generate token")
	}

	return c.JSON(http.StatusOK, map[string]string{"token": t})
}
