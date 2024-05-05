package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/souta17/go/dto"
	"github.com/souta17/go/handlerError"
	"github.com/souta17/go/helpers"
	"github.com/souta17/go/services"
)

type postHandler struct {
	service services.PostService
}

func NewPostHandler(s services.PostService) *postHandler {
	return &postHandler{service: s}
}

func (h *postHandler) Create(c *gin.Context) {
	var post dto.PostRequest
	if err := c.ShouldBindJSON(&post); err != nil {
		handlerError.HandlerError(c, &handlerError.BadRequestError{Message: err.Error()})
		return
	}

	if post.Picture != nil {
		if err := os.MkdirAll("/public/images", 0755); err != nil {
			handlerError.HandlerError(c, &handlerError.BadRequestError{Message: err.Error()})
			return
		}
		// ubah name gambarnya
		ext := filepath.Ext(post.Picture.Filename)
		newFileName := uuid.New().String() + ext

		// simpan gambar ke directori

		dst := filepath.Join("public/images/", filepath.Base(newFileName))
		c.SaveUploadedFile(post.Picture, dst)
		post.Picture.Filename = fmt.Sprintf("%s/public/images/%s", c.Request.Host, newFileName)
	}
	userID, _ := c.Get("UserID")
	post.UserId = userID.(int)

	if err := h.service.Create(&post); err != nil {
		handlerError.HandlerError(c, err)
		return
	}
	res := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "succses post your tweet",
	})

	c.JSON(http.StatusCreated, res)
}
