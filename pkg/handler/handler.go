package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gtikhomiroff/todo-app/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	routesAuth := router.Group("/auth")
	{
		routesAuth.POST("/sign-in", h.singIn)
		routesAuth.POST("/sign-up", h.singUp)
	}

	routesApi := router.Group("/api")
	{
		todoLists := routesApi.Group("/lists")
		{
			todoLists.POST("/", h.createTodoList)
			todoLists.GET("/", h.getTodoLists)
			todoLists.GET("/:id", h.getTodoList)
			todoLists.PUT("/:id", h.updateTodoList)
			todoLists.DELETE("/:id", h.deleteTodoList)

			todoItems := todoLists.Group("/:id/items")
			{
				todoItems.POST("/", h.createTodoItem)
				todoItems.GET("/", h.getTodoItems)
				todoItems.GET("/:item_id", h.getTodoItem)
				todoItems.PUT("/:item_id", h.updateTodoItem)
				todoItems.DELETE("/:item_id", h.deleteTodoItem)
			}
		}
	}

	return router
}
