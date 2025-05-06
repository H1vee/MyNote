package repository

import "github.com/golang-jwt/jwt/v5"

type IJWTRepository interface {
	GenerateAccessToken(userID uint) (string, error)
	GenerateRefreshToken(userID uint) (string, error)
	ValidateToken(tokenStr string) (*jwt.Token, error)
	ParseClaims(tokenStr string) (jwt.MapClaims, error)
}
