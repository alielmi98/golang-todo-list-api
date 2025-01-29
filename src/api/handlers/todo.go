package handlers

import (
	"net/http"
	"strconv"

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
// @Param Request body dto.CreateToDoRequest true "Create a todo"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.ToDoResponse} "ToDo response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/todo/ [post]
// @Security AuthBearer
func (h *ToDoHandler) CreateToDo(c *gin.Context) {
	var createReqDTO dto.CreateToDoRequest
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

// UpdateToDo godoc
// @Summary Update a todo
// @Description Update a todo job
// @Tags Todo
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateToDoRequest true "Update a todo"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.ToDoResponse} "Todo response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/todo/{id} [put]
// @Security AuthBearer
func (h *ToDoHandler) UpdateToDo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	var updateReqDTO dto.UpdateToDoRequest
	if err := c.ShouldBindJSON(&updateReqDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	response, err := h.todoService.UpdateToDo(c, id, &updateReqDTO)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(response, true, 0))
}

// DeleteTodo godoc
// @Summary Delete a todo
// @Description Delete a todo job by Id
// @Tags Todo
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 201 {object} helper.BaseHttpResponse{result=string} "Todo response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/todo/{id} [delete]
// @Security AuthBearer
func (h *ToDoHandler) DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	err := h.todoService.DeleteToDo(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse("The ToDo Job Deleted successfully", true, 0))
}

// GetToDoById godoc
// @Summary Get a todo
// @Description Get a todo job by Id
// @Tags Todo
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.ToDoResponse} "Todo response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/todo/{id} [get]
// @Security AuthBearer
func (h *ToDoHandler) GetToDoById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	response, err := h.todoService.GetToDoById(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(response, true, 0))
}

// GetToDosByFilter godoc
// @Summary      Get ToDos by filter with pagination
// @Description  Retrieve a list of ToDos based on filter criteria with pagination support
// @Tags         Todo
// @Accept       json
// @Produce      json
// @Param        paginationInput body dto.PaginationInputWithFilter true "Pagination and filter input"
// @Success      200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.ToDoResponse]} "Successful response with list of ToDos"
// @Failure      400 {object} helper.BaseHttpResponse "Invalid input"
// @Failure      404 {object} helper.BaseHttpResponse "No ToDos found"
// @Failure      500 {object} helper.BaseHttpResponse "Internal server error"
// @Router       /v1/todo/filter [post]
// @Security AuthBearer
func (h *ToDoHandler) GetToDosByFilter(c *gin.Context) {
	var paginationInput dto.PaginationInputWithFilter
	if err := c.ShouldBindJSON(&paginationInput); err != nil {
		c.JSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, helper.InvalidInputError, err))
		return
	}

	response, err := h.todoService.GetToDosByFilter(c, &paginationInput)
	if err != nil {
		c.JSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, helper.NotFoundError, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}
