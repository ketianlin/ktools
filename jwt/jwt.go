package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func (e Enter) GenerateToken(id int64, username string, second int, jwtKey string) (string, error) {
	uc := UserClaim{
		Id:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(second)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (e Enter) AnalyzeToken(token, jwtKey string) (*UserClaim, error) {
	uc := new(UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}
	return uc, err
}
