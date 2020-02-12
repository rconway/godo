package main

import (
	"fmt"
	"log"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

// UserTokenClaims zzz
type UserTokenClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// UserToken zzz
type UserToken struct {
	UserTokenClaims
}

func signingKey() []byte {
	key, ok := os.LookupEnv("JWT_SIGN_KEY")
	if ! ok || len(key) == 0 {
		key = "AllYourBase"
	}
	log.Printf("Using SIGN KEY = %v\n", key)
	return []byte(key)
}

// NewUserToken zzz
func NewUserToken(claims *UserTokenClaims) *UserToken {
	ut := &UserToken{}
	if claims != nil {
		ut.UserTokenClaims = *claims
	} else {
		ut.UserTokenClaims = UserTokenClaims{}
	}
	return ut
}

// EncodeJWT zzz
func (ut *UserToken) EncodeJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, ut.UserTokenClaims)
	ss, err := token.SignedString(signingKey())
	return ss, err
}

// DecodeJWT zzz
func (ut *UserToken) DecodeJWT(ss string) bool {
	token, err := jwt.ParseWithClaims(ss, &UserTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey(), nil
	})

	claims, ok := token.Claims.(*UserTokenClaims)
	ok = ok && token.Valid

	if ok {
		ut.UserTokenClaims = *claims
	} else {
		fmt.Printf("Decode ERROR: %v\n", err)
	}

	return ok
}
