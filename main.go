package main

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func main() {
	log.Println("...main...")

	type CustomToken struct {
		Age int
		UserToken
	}

	// Outgoing token
	tokenTx := &CustomToken{
		Age: 123,
		UserToken: UserToken{
			Username: "fred",
			StandardClaims: jwt.StandardClaims{
				Issuer:    "richard",
				ExpiresAt: time.Now().Unix() + 60,
			},
		},
	}

	// Create token
	ss, err := tokenTx.ToSignedString()
	log.Println(ss)

	if err != nil {
		log.Fatal(err)
	}

	// Receive token
	tokenRx := &UserToken{}
	ok := tokenRx.FromSignedString(ss)

	if ok {
		log.Printf("Username: %v - Issuer: %v - Expires: %v\n", tokenRx.Username, tokenRx.Issuer, tokenRx.ExpiresAt-time.Now().Unix())
	} else {
		log.Fatal("ERROR decoding token")
	}

}
