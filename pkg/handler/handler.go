package handler

import (
	"github.com/CGNihill/go-rest-api/pkg/service"
	"github.com/gin-gonic/gin"
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
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	api := router.Group("/api", h.userIdentity)
	{
		list := api.Group("/lists")
		{
			list.POST("/", h.createList)
			list.GET("/", h.getAllLists)
			list.GET("/:id", h.getListById)
			list.POST("/:id", h.updateList)
			list.DELETE("/:id", h.deleteList)

			items := list.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
			}
		}

		items := api.Group("items")
		{
			items.GET("/:id", h.getItemById)
			items.POST("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}
	}
	return router
}
