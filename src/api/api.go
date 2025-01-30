package api

import (
	"fmt"
	"log"

	"github.com/alielmi98/golang-todo-list-api/api/middlewares"
	"github.com/alielmi98/golang-todo-list-api/api/routers"
	"github.com/alielmi98/golang-todo-list-api/config"
	"github.com/alielmi98/golang-todo-list-api/constants"
	"github.com/alielmi98/golang-todo-list-api/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	r := gin.New()

	r.Use(middlewares.Cors(cfg))
	RegisterRoutes(r, cfg)
	RegisterSwagger(r, cfg)
	log.Printf("Caller:%s Level:%s Msg:%s", constants.General, constants.Startup, "Started")
	r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		//User
		users := v1.Group("/users")
		tokens := v1.Group("/tokens")
		todo := v1.Group("/todo", middlewares.Authentication(cfg))
		routers.User(users, cfg)
		routers.Token(tokens, cfg)
		routers.Todo(todo, cfg)

	}

}
func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "golang todo list api documentation"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.InternalPort)
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
