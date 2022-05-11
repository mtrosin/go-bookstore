package controllers

import(
	"time"
	"encoding/json"
	"net/http"
	"github.com/mtrosin/go-bookstore/pkg/config"
	"github.com/mtrosin/go-bookstore/pkg/models"
	"github.com/mtrosin/go-bookstore/pkg/utils"
	"github.com/dgrijalva/jwt-go"
)

type Credentials struct {
	Password string `json:"password"`
	Login string `json:"login"`
}

type Claims = utils.Claims

func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the expected password from our in memory map
	user, _ := models.GetUser(creds.Login, creds.Password)
	if user.Login == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Login: creds.Login,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(config.GetJwtKey())
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}