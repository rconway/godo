package main

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// MyCustomClaims zzz
type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}

func createToken() string {

	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(mySigningKey)

	return ss

}

func decodeToken(ss string) {

	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	token, err := jwt.ParseWithClaims(ss, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.Foo, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err)
	}

	jwt.TimeFunc = time.Now

}

func main() {
	log.Println("...main...")

	ss := createToken()
	log.Println(ss)

	decodeToken(ss)
}
