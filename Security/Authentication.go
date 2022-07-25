package security

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go/v4"
)

func JwtAccessToken(infoToken string) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Signature = "BRUH"
	claims := token.Claims.(jwt.MapClaims)
	claims["Token"] = time.Now().Add(time.Hour * 2).Unix()

	//Creates InfoToken
	func() {

	}()
}

func JwtInfoToken() {

}
