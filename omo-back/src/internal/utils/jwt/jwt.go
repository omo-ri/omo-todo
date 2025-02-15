package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"log"
	omoconfig "omo-back/src/config"
	"omo-back/src/internal/model"
	"time"
)

func GenerateSignedToken(userID *uuid.UUID) (*string, error) {
	expirationTime := time.Now().Add(omoconfig.JwtExpireTime)
	claims := jwt.MapClaims{
		"userID": userID.String(),
		"exp":    expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(omoconfig.JwtSecret))
	if err != nil {
		log.Printf("Failed to generate token: %v", err)
		return nil, err
	}
	return &tokenString, nil
}

func ParseJwtToken(token *jwt.Token) *model.ParsedJWT {
	if token == nil {
		return nil
	}
	if token.Method != jwt.SigningMethodHS256 {
		log.Printf("Invalid signing method: %v", token.Method)
		return nil
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		UserIDStr, ok := claims["userID"].(string)
		if !ok {
			log.Printf("Failed to parse userID from token")
			return nil
		}
		UserID, err := uuid.Parse(UserIDStr)
		if err != nil {
			log.Printf("Failed to parse userID from token: %v", err)
			return nil
		}
		expirationTime, ok := claims["exp"].(float64)
		if !ok {
			log.Printf("Failed to parse expiration time from token")
			return nil
		}
		return &model.ParsedJWT{
			UsedID:         UserID,
			ExpirationTime: time.Unix(int64(expirationTime), 0),
		}
	}
	return nil
}
