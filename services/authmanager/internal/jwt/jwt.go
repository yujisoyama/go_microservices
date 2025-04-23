package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yujisoyama/go_microservices/pkg/utils"
)

type JWTService struct {
	SecretKey []byte
	ExpTime   int64
	IatTime   int64
}

func NewJWTConfigs() *JWTService {
	return &JWTService{
		SecretKey: []byte(utils.GetEnv("JWT_SECRET_KEY")),
		ExpTime:   time.Now().Add(5 * time.Hour).Unix(),
		IatTime:   time.Now().Unix(),
	}
}

type TokenInfo struct {
	jwt.RegisteredClaims
	UserId  string `json:"user_id"`
	OAuthId string `json:"oauth_id"`
}

func (jwtService *JWTService) GenerateToken(info TokenInfo) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  info.UserId,
		"oauth_id": info.OAuthId,
		"exp":      jwtService.ExpTime,
		"iat":      jwtService.IatTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtService.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (jwtService *JWTService) VerifyToken(tokenString string) (*TokenInfo, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenInfo{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return jwtService.SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Token is invalid")
	}

	if tokenInfo, ok := token.Claims.(*TokenInfo); ok && token.Valid {
		if tokenInfo.ExpiresAt.Unix() < time.Now().Unix() {
			return nil, fmt.Errorf("Token expired")
		} else {
			return tokenInfo, nil
		}
	}

	return nil, fmt.Errorf("Token is invalid")
}
