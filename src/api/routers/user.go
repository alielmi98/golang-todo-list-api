package routers

import (
	"github.com/alielmi98/golang-todo-list-api/api/handlers"
	"github.com/alielmi98/golang-todo-list-api/config"

	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUsersHandler(cfg)
	router.POST("/login-by-username", h.LoginByUsername)
	router.POST("/register-by-username", h.RegisterByUsername)

}
