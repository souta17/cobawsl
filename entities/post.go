package entities

import "time"

type Post struct {
	ID         int
	UserId     int
	Tweet      string
	PictureUrl *string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
