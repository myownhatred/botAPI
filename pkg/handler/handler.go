package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/myownhatred/botAPI/pkg/service"
)

type Handler struct {
	service       *service.Service
	googleService *service.GoogleService
}

func NewHandler(ser *service.Service, gser *service.GoogleService) *Handler {
	return &Handler{
		service:       ser,
		googleService: gser,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.singUp)
		auth.POST("/sign-in", h.singIn)
	}

	todo := router.Group("/todo")
	{
		lists := todo.Group("/lists")
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

	google := router.Group("/google")
	{
		google.GET("/pic", h.getPicture)
		google.POST("/pic", h.getPicture)
		google.GET("/vid", h.getVideo)
		google.POST("/vid", h.getVideo)
	}

	return router
}
