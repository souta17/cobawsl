package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/souta17/go/dto"
	"github.com/souta17/go/handlerError"
	"github.com/souta17/go/helpers"
	"github.com/souta17/go/services"
)

type authHandler struct {
	service services.AuthService
}

func NewAuthHandler(s services.AuthService) *authHandler {
	return &authHandler{s}
}

func (h *authHandler) Register(c *gin.Context) {
	var register dto.RegisterRequest
	if err := c.ShouldBindJSON(&register); err != nil {
		handlerError.HandlerError(c, &handlerError.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.Register(&register); err != nil {
		handlerError.HandlerError(c, err)
		return
	}
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "register has been succesful",
	})

	c.JSON(http.StatusCreated, response)
}
