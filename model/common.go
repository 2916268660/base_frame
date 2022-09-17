package model

import (
	"github.com/golang-jwt/jwt"
)

type PageCommon struct {
	PageSize int
	PageNum  int
}

type MyClaims struct {
	UserId     string `json:"userId"`
	UserName   string `json:"userName"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Department string `json:"department"`
	Role       string `json:"role"`
	jwt.StandardClaims
}
