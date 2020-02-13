package main

import (
	"fmt"
	"log"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

// DefaultSecretKey zzz
const DefaultSecretKey = "MySecretKey"

// UserToken zzz
type UserToken struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// signingKey zzz
func signingKey() []byte {
	key, ok := os.LookupEnv("JWT_SIGN_KEY")
	if !ok || len(key) == 0 {
		key = DefaultSecretKey
	}
	log.Printf("Using SIGN KEY = %v\n", key)
	return []byte(key)
}

// ToSignedString zzz
func (ut *UserToken) ToSignedString() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, ut)
	ss, err := token.SignedString(signingKey())
	return ss, err
}

// FromSignedString zzz
func (ut *UserToken) FromSignedString(ss string) bool {
	token, err := jwt.ParseWithClaims(ss, &UserToken{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey(), nil
	})

	// TODO: handle err

	claims, ok := token.Claims.(*UserToken)
	ok = ok && token.Valid

	if ok {
		*ut = *claims
	} else {
		fmt.Printf("Decode ERROR: %v\n", err)
	}

	return ok
}
