package api

import (
	"github.com/labstack/echo/v4"
	"github.com/praveenmahasena647/go-stickers/cmd/app/api/handlers"
)

type Server struct {
	listenAddr string
}

func New() *Server {
	return &Server{}
}

func (s *Server) Run(port string) error {
	var e = echo.New()

	e.GET("/", handlers.GetAll)
	e.GET("/:id", handlers.GetbyID)
	e.POST("/", handlers.PostOne)
	e.PUT("/:id", handlers.EditbyID)
	e.DELETE("/:id", handlers.DeletebyID)
	return e.Start(port)
}
