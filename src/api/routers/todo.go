package routers

import (
	"github.com/alielmi98/golang-todo-list-api/api/handlers"
	"github.com/alielmi98/golang-todo-list-api/api/middlewares"
	"github.com/alielmi98/golang-todo-list-api/config"
	"github.com/gin-gonic/gin"
)

func Todo(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewToDoHandler()
	router.POST("/", h.CreateToDo)
	router.POST("/filter", h.GetToDosByFilter)
	router.PUT("/:id", middlewares.AuthorizeTodoAccess(), h.UpdateToDo)
	router.GET("/:id", middlewares.AuthorizeTodoAccess(), h.GetToDoById)
	router.DELETE("/:id", middlewares.AuthorizeTodoAccess(), h.DeleteTodo)

}
