package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/msvetkov/notes-app/pkg/service"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "github.com/msvetkov/notes-app/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.DELETE("/delete-account", h.deleteUser)
	}

	api := router.Group("/api", h.userIdentity)
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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
