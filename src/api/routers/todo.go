package routers

import (
	"github.com/alielmi98/golang-todo-list-api/api/handlers"
	"github.com/alielmi98/golang-todo-list-api/config"
	"github.com/gin-gonic/gin"
)

func Todo(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewToDoHandler()
	router.POST("/", h.CreateToDo)

}
