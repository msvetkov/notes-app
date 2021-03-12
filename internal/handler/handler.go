package handler

import (
	"github.com/gin-gonic/gin"
	"notes-app/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		notes := api.Group("/notes")
		{
			notes.GET("/", h.getNotes)
			notes.GET("/:id", h.getNote)
			notes.POST("/", h.createNote)
			notes.PUT("/:id", h.updateNote)
			notes.DELETE("/:id", h.deleteNote)
		}
	}

	return router
}
