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
	// claims := MyCustomClaims{
	// 	"bar",
	// 	jwt.StandardClaims{
	// 		ExpiresAt: 15000,
	// 		Issuer:    "test",
	// 	},
	// }
	claims := MyCustomClaims{}
	claims.Foo = "fred"
	claims.StandardClaims.ExpiresAt = time.Now().Unix() + 3600

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Printf("Sign ERROR: %v\n", err)
	}

	return ss

}

func decodeToken(ss string) {
	token, err := jwt.ParseWithClaims(ss, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		fmt.Printf("%v %v\n", claims.Foo, claims.StandardClaims.ExpiresAt - time.Now().Unix())
	} else {
		fmt.Printf("Decode ERROR: %v\n", err)
	}
}

func main() {
	log.Println("...main...")

	ss := createToken()
	log.Println(ss)

	decodeToken(ss)
}
