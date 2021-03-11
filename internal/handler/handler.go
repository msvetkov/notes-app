package handler

import "github.com/gin-gonic/gin"

type Handler struct {
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
