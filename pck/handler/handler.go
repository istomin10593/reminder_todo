package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/istomin10593/reminder_todo/pck/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.sighUp)
		auth.POST("/sigh-in", h.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)

			}
		}
	}
	return router
}
