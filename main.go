package main

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func signingKey() []byte {
	return []byte("AllYourBase")
}

// MyCustomClaims zzz
type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func createToken() string {

	// Create the Claims
	claims := MyCustomClaims{}
	// 	"rconway",
	// 	jwt.StandardClaims{
	// 		ExpiresAt: time.Now().Unix() + 3600,
	// 		Issuer:    "godo",
	// 	},
	// }
	claims.Username = "rconway"
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + 3600,
		Issuer:    "godo",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey())

	if err != nil {
		fmt.Printf("Sign ERROR: %v\n", err)
	}

	return ss

}

func decodeToken(ss string) {
	token, err := jwt.ParseWithClaims(ss, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey(), nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		fmt.Printf("%v %v\n", claims.Username, claims.StandardClaims.ExpiresAt - time.Now().Unix())
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
