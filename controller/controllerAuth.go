package controller

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/tubes/Art-Auction-Tubes/model"
)

var token = os.Getenv("TOKEN")
var jwtKey = []byte(token)
var tokenName = "token"

type Claims struct {
	ID       int    `json:id`
	Username string `json:Username`
	UserType int    `json:user_type`
	jwt.StandardClaims
}

func generateToken(w http.ResponseWriter, id int, username string, userType int) {
	tokenExpiryTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		ID:       id,
		Username: username,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpiryTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)

	if err != nil {
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    signedToken,
		Expires:  tokenExpiryTime,
		Secure:   false,
		HttpOnly: true,
	})
}

func resetUserToken(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    "",
		Expires:  time.Now(),
		Secure:   false,
		HttpOnly: true,
	})
}

func Authenticate(next http.HandlerFunc, accessType int) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isValidToken := validateUserToken(r, accessType)
		if !isValidToken {
			sendUnAuthorizedResponse(w)
		} else {
			next.ServeHTTP(w, r)
		}

	})
}

func validateUserToken(r *http.Request, accessType int) bool {
	isAccessTokenValid, userType := validateTokenFromCookies(r)
	if isAccessTokenValid {
		isUserValid := userType == accessType
		if accessType == 2 {
			isUserValid = true
		}
		if isUserValid {
			return true
		}
	}
	return false
}

func validateTokenFromCookies(r *http.Request) (bool, int) {
	if cookie, err := r.Cookie(tokenName); err == nil {
		accessToken := cookie.Value
		accessClaims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err == nil && parsedToken.Valid {
			return true, accessClaims.UserType
		}
	}
	return false, -1
}

func sendUnAuthorizedResponse(w http.ResponseWriter) {
	var response model.GeneralResponse
	response.Status = 401
	response.Message = "Unauthorized Status"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
