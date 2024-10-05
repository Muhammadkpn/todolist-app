package model

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtMapClaims struct {
	DataJwt
	//	RoleTypeName string `json:"RoleTypeName"`
	jwt.RegisteredClaims
}
type JwtMapClaimsReset struct {
	DataJwtReset
	//	RoleTypeName string `json:"RoleTypeName"`
	jwt.RegisteredClaims
}
type DataJwt struct {
	Id       string
	Username string
	Email    string
}

type DataJwtReset struct {
	ID        string
	ResetCode string
}
