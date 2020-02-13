package main

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func main() {
	log.Println("...main...")

	// Outgoing token
	tokenTx := &UserToken{}
	tokenTx.Username = "rconway"
	tokenTx.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + 3600,
		Issuer:    "godo",
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
