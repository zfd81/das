package http

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
)

type PortalClaims struct {
	Uid string `json:"uid"`
	jwt.StandardClaims
}

type Security struct {
	secret []byte
	second int
}

func (s *Security) CreateToken(claims *PortalClaims, second int) (string, error) {
	now := time.Now()
	claims.IssuedAt = now.Unix()
	claims.ExpiresAt = now.Add(time.Second * time.Duration(second)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.secret)
	return tokenString, err
}

func (s *Security) ParseToken(tokenString string) (*PortalClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &PortalClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secret), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*PortalClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

func NewSecurity(secret string) *Security {
	return &Security{secret: []byte(secret)}
}
