package service

import "context"

type UserService interface {
	Register(ctx context.Context, username string, email string, password string) error
}
