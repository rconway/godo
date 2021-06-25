package main

import (
	"log"
	"time"

	"github.com/rconway/godo/jwt"
)

func main() {
	log.Println("...main...")

	type CustomClaims struct {
		Age int `json:"age"`
		jwt.UserClaims
	}

	// Outgoing claims
	claimsTx := CustomClaims{
		Age: 123,
		UserClaims: jwt.UserClaims{
			Username: "fred",
			StandardClaims: jwt.StandardClaims{
				Issuer:    "richard",
				ExpiresAt: time.Now().Unix() + 60,
			},
		},
	}

	// Create token as string
	ss, err := jwt.ClaimsToSignedString(&claimsTx)
	log.Println(ss)

	if err != nil {
		log.Fatal(err)
	}

	// Receive token as string
	claims, err := jwt.ClaimsFromSignedString(ss, &CustomClaims{})
	claimsRx := claims.(*CustomClaims)

	if err == nil {
		log.Printf("Age: %v - Username: %v - Issuer: %v - Expires: %v\n", claimsRx.Age, claimsRx.Username, claimsRx.Issuer, claimsRx.ExpiresAt-time.Now().Unix())
	} else {
		log.Fatal("ERROR decoding token")
	}

}
