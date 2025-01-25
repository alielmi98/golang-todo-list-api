package handlers

import (
	"net/http"

	"github.com/alielmi98/golang-todo-list-api/api/dto"
	"github.com/alielmi98/golang-todo-list-api/api/helper"
	"github.com/alielmi98/golang-todo-list-api/services"
	"github.com/gin-gonic/gin"
)

type ToDoHandler struct {
	todoService services.ToDoService
}

func NewToDoHandler() *ToDoHandler {
	return &ToDoHandler{todoService: services.NewToDoService()}
}

// CreateTodo godoc
// @Summary Create a todo
// @Description Create a new todo job
// @Tags Todo
// @Accept json
// @produces json
// @Param Request body dto.CreateTodoRequest true "Create a todo"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.TodoResponse} "Todo response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/todo/ [post]
// @Security AuthBearer
func (h *ToDoHandler) CreateToDo(c *gin.Context) {
	var createReqDTO dto.CreateTodoRequest
	if err := c.ShouldBindJSON(&createReqDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	response, err := h.todoService.CreateToDo(c, &createReqDTO)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(response, true, 0))
}
