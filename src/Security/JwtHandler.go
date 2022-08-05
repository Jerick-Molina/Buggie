package security

import (
	"fmt"
	"net/http"
	"time"

	properties "github.com/Jerick-Molina/Buggie/Properties"

	"github.com/dgrijalva/jwt-go/v4"
)

var secretKey = []byte("p8cafxzquew4juy1rk9f")

func JwtAccessToken(u properties.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	claims["iss"] = "localhost"
	claims["userId"] = u.Id
	claims["role"] = u.Role
	claims["companyId"] = u.CompanyId
	//Creates InfoToken

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		//
		return "", nil
	}

	return tokenString, nil
}

func IsAuthorized(roles []string, token string) (jwt.MapClaims, int, string) {

	if token != "" {
		token = token[7:]
		token, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, nil
			}
			return secretKey, nil
		})
		if err != nil {
			fmt.Println(err.Error())
			return nil, http.StatusUnauthorized, "Token signature is invalid"
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			for i := 0; i < len(roles); i++ {

				if roles[i] == claims["role"] {
					return claims, http.StatusOK, "Valid"
				}
			}

		}
		return nil, http.StatusUnauthorized, "Role not valid"
	}

	return nil, http.StatusUnauthorized, "JWT doesn't exist"

}
