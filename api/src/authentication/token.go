// uses an external package json web token
// go get github.com/dgrijalva/jwt-go
package authentication

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken  generates a token with permitions
func CrateToken(userID uint64) (string, error) {
	permitions := jwt.MapClaims{}
	permitions["authorized"] = true

	//expires the token in 6 hours (timenow + 6hours, and converts in Unix format)
	permitions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permitions["user.ID"] = userID

	//generate a secret / token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permitions)
	return token.SignedString([]byte(config.SecretKey))

}

// Validate Token validates
func ValidateToken(r *http.Request) error {

	tokenstring := extractToken(r)

	token, err := jwt.Parse(tokenstring, returnverificationkey)

	if err != nil {

		return errors.New("Invalid Token")
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return nil
	}

	return errors.New("Invalid Token")
}

func ExtractUserId(r *http.Request) (uint64, error) {
	tokenstring := extractToken(r)
	token, err := jwt.Parse(tokenstring, returnverificationkey)
	if err != nil {
		return 0, errors.New("Invalid Token")
	}
	if permitions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", permitions["user.ID"]), 10, 64)
		if err != nil {
			return 0, err
		}

		return userID, nil
	}

	return 0, errors.New("Invalid Token")

}

// Receives the Bearer and o token   - Bearer is the person that requst the token
func extractToken(r *http.Request) string {

	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 { // validate if the request comes with two words (beares and token)

		return strings.Split(token, " ")[1] // returns the second word  - the token

	}

	return ""
}

// verify is the method used to generate the token is correct
func returnverificationkey(token *jwt.Token) (interface{}, error) {

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

		return nil, fmt.Errorf("Unespected Sign Method! %v", token.Header["alg"])

	}
	return config.SecretKey, nil

}
