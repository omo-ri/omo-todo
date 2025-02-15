package model

import (
	"github.com/google/uuid"
	"time"
)

type ParsedJWT struct {
	UsedID         uuid.UUID
	ExpirationTime time.Time
}
