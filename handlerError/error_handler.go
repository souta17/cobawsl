package handlerError

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/souta17/go/dto"
	"github.com/souta17/go/helpers"
)

func HandlerError(c *gin.Context, err error) {
	var statusCode int

	switch err.(type) {
	case *NotfoundError:
		statusCode = http.StatusNotFound
	case *InternalServerError:
		statusCode = http.StatusInternalServerError
	case *UnauthorizedError:
		statusCode = http.StatusUnauthorized
	case *BadRequestError:
		statusCode = http.StatusBadRequest
	}
	response := helpers.Response(dto.ResponseParams{
		StatusCode: statusCode,
		Message:    err.Error(),
	})
	c.JSON(statusCode, response)
}
