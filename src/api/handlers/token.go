package handlers

import (
	"net/http"

	"github.com/alielmi98/golang-todo-list-api/api/dto"
	"github.com/alielmi98/golang-todo-list-api/api/helper"
	"github.com/alielmi98/golang-todo-list-api/config"
	"github.com/alielmi98/golang-todo-list-api/services"

	"github.com/gin-gonic/gin"
)

type TokensHandler struct {
	service *services.TokenService
}

func NewTokensHandler(cfg *config.Config) *TokensHandler {
	service := services.NewTokenService(cfg)
	return &TokensHandler{service: service}
}

// RefreshToken godoc
// @Summary RefreshToken
// @Description RefreshToken
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/tokens/refresh-token [post]
func (h *TokensHandler) RefreshToken(c *gin.Context) {
	tokenDetail, err := h.service.RefreshToken(c)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.ValidationError, err))
		return
	}
	response := new(dto.TokenResponse)
	response.AccessToken = tokenDetail.AccessToken
	response.AccessTokenExpireTime = tokenDetail.AccessTokenExpireTime

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(response, true, helper.Success))
}
