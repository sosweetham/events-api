package utils

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"kodski.com/events-api/env"
)

type JWTAuth struct {
	Email string `binding:"required" json:"email"`
	UserId int64 `binding:"required" json:"userId"`
	Exp int64 `json:"exp"`
}

func (jwtAuth *JWTAuth) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": jwtAuth.Email,
		"userId": jwtAuth.UserId,
		"exp": jwtAuth.Exp,
	})

	return token.SignedString([]byte(env.AppEnv.JWTSecret))
}

func NewJWTAuth(email string, userId int64, exp int64) *JWTAuth {
	return &JWTAuth{
		Email: email,
		UserId: userId,
		Exp: exp,
	}
}

func VerifyToken(token string) (*JWTAuth, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(env.AppEnv.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("invalid claims")
	}

	email, _ := claims["email"].(string)
	userId, _ := (claims["userId"].(float64))
	exp, _ := claims["exp"].(float64)

	jwtAuth := NewJWTAuth(email, int64(userId), int64(exp))

	return jwtAuth, nil
}