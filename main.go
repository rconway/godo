package main

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func main() {
	log.Println("...main...")

	// ss := createToken()
	// log.Println(ss)
	// decodeToken(ss)

	claims := &UserTokenClaims{}
	claims.Username = "rconway"
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + 3600,
		Issuer:    "godo",
	}

	// Create token
	token := NewUserToken(claims)
	ss, err := token.EncodeJWT()
	log.Println(ss)

	if err != nil {
		log.Fatal(err)
	}

	// Receive token
	token2 := NewUserToken(nil)
	ok := token2.DecodeJWT(ss)

	if ok {
		log.Printf("Username: %v - Issuer: %v - Expires: %v\n", token2.Username, token2.Issuer, token.ExpiresAt - time.Now().Unix())
	} else {
		log.Fatal("ERROR decoding token")
	}

}
