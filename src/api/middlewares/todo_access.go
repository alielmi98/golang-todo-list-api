package middlewares

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/alielmi98/golang-todo-list-api/api/helper"
	"github.com/alielmi98/golang-todo-list-api/constants"
	"github.com/alielmi98/golang-todo-list-api/services"
	"github.com/gin-gonic/gin"
)

func AuthorizeTodoAccess() gin.HandlerFunc {
	service := services.NewToDoService()
	return func(c *gin.Context) {
		userId := int(c.Value(constants.UserIdKey).(float64))
		todoId, _ := strconv.Atoi(c.Params.ByName("id"))

		todo, err := service.GetToDoById(c.Request.Context(), todoId)
		if err != nil {
			c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
				helper.GenerateBaseResponseWithError(nil, false, helper.NotFoundError, err))
			return
		}

		if todo.UserId != userId {
			err := errors.New("access denied")
			c.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateBaseResponseWithError(
				nil, false, helper.AuthError, err,
			))
			return
		}

		c.Next()
	}
}
