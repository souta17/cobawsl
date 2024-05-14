package services

import (
	"github.com/souta17/go/dto"
	"github.com/souta17/go/entities"
	"github.com/souta17/go/handlerError"
	"github.com/souta17/go/repositories"
)

type PostService interface {
	Create(req *dto.PostRequest) error
}

type postService struct {
	repository repositories.PostRepository
}

func NewPostService(r repositories.PostRepository) *postService {
	return &postService{
		repository: r,
	}
}

func (s *postService) Create(req *dto.PostRequest) error {
	post := entities.Post{
		UserId: req.UserId,
		Tweet:  req.Tweet,
	}
	if req.Picture != nil {
		post.PictureUrl = &req.Picture.Filename
	}
	if err := s.repository.Create(&post); err != nil {
		return &handlerError.InternalServerError{Message: err.Error()}
	}

	return nil
}
