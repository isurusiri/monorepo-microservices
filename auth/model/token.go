package model

import "github.com/dgrijalva/jwt-go"

// TokenClaim model
type TokenClaim struct {
	jwt.StandardClaims
	Csrf string `json:"csrf"`
}

// Token model
type Token struct {
	AuthToken    string
	RefreshToken string
	CSRFKey      string
}
