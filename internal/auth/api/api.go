package api

import (
	"fmt"

	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type Server struct {
	server    *echo.Echo
	address   string
	secretKey string
	uc        Usecase
}

func NewServer(ip string, port int, secretKey string, uc Usecase) *Server {
	srv := &Server{
		server:    echo.New(),
		secretKey: secretKey,
		uc:        uc,
		address:   fmt.Sprintf("%s:%d", ip, port),
	}

	// Роуты
	srv.server.POST("/register", srv.Register)
	srv.server.POST("/login", srv.Login)

	return srv
}

func (srv *Server) Run() {
	srv.server.Logger.Fatal(srv.server.Start(srv.address))
}

func JWTMiddleware(secretKey string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(secretKey),
	})
}