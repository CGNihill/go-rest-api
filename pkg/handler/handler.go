package handler

import "github.com/gin-gonic/gin"

type Handler struct{}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-in")
		auth.POST("/sign-up")
	}

	api := router.Group("/api")
	{
		list := api.Group("/lists")
		{
			list.POST("/")
			list.GET("/")
			list.GET("/:id")
			list.POST("/:id")
			list.DELETE("/:id")

			items := list.Group(":id/items")
			{
				items.POST("/")
				items.GET("/")
				items.GET("/:item_id")
				items.POST("/:item_id")
				items.DELETE("/:item_id")
			}
		}
	}
	return router
}
