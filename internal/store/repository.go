package store

import "Forum/internal/model"

type UserRepository interface {
	Login(email,password string) error
	Register(user *model.User) error
}

type PostRepository interface {
	Create()
	Delete()
}
