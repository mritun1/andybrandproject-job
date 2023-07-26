package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secret = os.Getenv("SECRET_KEY")
var SECRET = []byte(secret)

// --------------------------------------------------------------
// CREATE JWT TOKEN - START
// --------------------------------------------------------------
func CreateKey(email string, user_id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["email"] = email
	claims["user_id"] = user_id
	tokenStr, err := token.SignedString(SECRET)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

// --------------------------------------------------------------
// CREATE JWT TOKEN - END
// --------------------------------------------------------------
