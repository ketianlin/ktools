package jwt

import "github.com/dgrijalva/jwt-go"

type Enter struct{}

type UserClaim struct {
	Id       int64
	Username string
	jwt.StandardClaims
}
