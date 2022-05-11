package utils

import(
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/mtrosin/go-bookstore/pkg/config"
)

type Claims struct {
	Login string `json:"login"`
	jwt.StandardClaims
}

func CheckToken(r *http.Request) int {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			return http.StatusUnauthorized
		}
		return http.StatusBadRequest
	}

	tknStr := c.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return config.GetJwtKey(), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return http.StatusUnauthorized
		}
		return http.StatusBadRequest
	}
	if !tkn.Valid {
		return http.StatusUnauthorized
	}

	return http.StatusOK
}