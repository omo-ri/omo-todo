package database

import (
	"context"
	"github.com/google/uuid"
	"omo-back/src/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, username string, email string, passwordHash string) *uuid.UUID
	GetUserByUuid(ctx context.Context, uuid string) *model.User
	GetUserByUsername(ctx context.Context, username string) *model.User
}
