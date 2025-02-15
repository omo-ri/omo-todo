package impl

import (
	"context"
	"github.com/google/uuid"
	"log"
	"omo-back/src/database"
	"omo-back/src/internal/model"
	"time"
)

var _ database.UserRepository = (*UserRepositoryImpl)(nil)

type UserRepositoryImpl struct {
	postgresService *PostgresServiceImpl
}

func NewUserRepositoryImpl(postgresService *PostgresServiceImpl) *UserRepositoryImpl {
	return &UserRepositoryImpl{postgresService: postgresService}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, username string, email string, passwordHash string) *uuid.UUID {
	userId := uuid.New()
	_, err := r.postgresService.Exec(ctx,
		"INSERT INTO users (id, username, password_hash, email, avatar, last_login) VALUES ($1, $2, $3, $4, '', $5)",
		userId, username, passwordHash, email, time.Now())
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return nil
	}
	return &userId
}

func (r *UserRepositoryImpl) GetUserByUuid(ctx context.Context, uuid string) *model.User {
	var user model.User
	var row = r.postgresService.QueryRow(ctx, "SELECT id, username, email, avatar, last_login FROM users WHERE id = $1", uuid)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Avatar)
	if err != nil {
		log.Printf("Failed to get user by uuid: %v", err)
		return nil
	}
	return &user
}

func (r *UserRepositoryImpl) GetUserByUsername(ctx context.Context, username string) *model.User {
	var user model.User
	var row = r.postgresService.QueryRow(ctx, "SELECT id, username, email, avatar, last_login FROM users WHERE username = $1", username)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Avatar)
	if err != nil {
		log.Printf("Failed to get user by username: %v", err)
		return nil
	}
	return &user
}
