package dto

import "mime/multipart"

type PostResponse struct {
	ID         int    `json:"id"`
	UserID     string `json:"_"`
	User       User   `gorm:"foreignKey:UserID" json:"user"`
	Tweet      string `json:"tweet"`
	PictureUrl string `json:"picture_url"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type PostRequest struct {
	UserId  int                   `form:"user_id"`
	Tweet   string                `form:"tweet"`
	Picture *multipart.FileHeader `form:"picture"`
}

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
