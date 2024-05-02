package services

import (
	"github.com/souta17/go/dto"
	"github.com/souta17/go/entities"
	"github.com/souta17/go/handlerError"
	"github.com/souta17/go/helpers"
	"github.com/souta17/go/repositories"
)

type AuthService interface {
	Register(req *dto.RegisterRequest) error
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	repository repositories.AuthRepository
}

func NewAuthService(r repositories.AuthRepository) *authService {
	return &authService{
		repository: r,
	}
}

func (s *authService) Register(req *dto.RegisterRequest) error {
	if emailExist := s.repository.EmailExist(req.Email); emailExist {
		return &handlerError.BadRequestError{
			Message: "email already in use",
		}
	}

	if req.Password != req.RePassword {
		return &handlerError.BadRequestError{
			Message: "password do not match",
		}
	}

	hashPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		return &handlerError.InternalServerError{
			Message: err.Error(),
		}
	}
	user := entities.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashPassword,
		Gender:   req.Gender,
	}

	if err := s.repository.Register(&user); err != nil {
		return &handlerError.InternalServerError{Message: err.Error()}
	}
	return err
}

func (s *authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	var data dto.LoginResponse

	user, err := s.repository.GetUserByEmail(req.Email)
	if err != nil {
		return nil, &handlerError.NotfoundError{
			Message: "wrong email or password",
		}
	}
	if err := helpers.ComparePassword(user.Password, req.Password); err != nil {
		return nil, &handlerError.NotfoundError{
			Message: "wrong email or password",
		}
	}
	token, err := helpers.GenerateToken(user)
	if err != nil {
		return nil, &handlerError.InternalServerError{Message: err.Error()}
	}
	data = dto.LoginResponse{
		ID:    user.ID,
		Name:  user.Email,
		Token: token,
	}
	return &data, nil
}
