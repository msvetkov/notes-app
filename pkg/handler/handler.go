package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/msvetkov/notes-app/docs"
	"github.com/msvetkov/notes-app/pkg/service"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"time"
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

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "PUT", "DELETE", "POST", "Content-Length"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "Origin"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	return router
}
