package impl

import (
	"context"
	"fmt"
	"omo-back/src/database/impl"
)

type UserServiceImpl struct {
	userRepository *impl.UserRepositoryImpl
}

func NewUserServiceImpl(userRepository *impl.UserRepositoryImpl) *UserServiceImpl {
	return &UserServiceImpl{userRepository: userRepository}
}

func (u *UserServiceImpl) Register(ctx context.Context, username string, email string, password string) error {
	existingUser := u.userRepository.GetUserByUsername(ctx, username)
	if existingUser != nil {
		return fmt.Errorf("user with username %s already exists", username)
	}
	userId := u.userRepository.Create(ctx, username, email, password)
	if userId == nil {
		return fmt.Errorf("failed to create user")
	}
	return nil
}
