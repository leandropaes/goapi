package lib

import (
	"github.com/dgrijalva/jwt-go"
	_ "github.com/joho/godotenv/autoload"
	"github.com/leandropaes/goapi/app/models"
	"os"
	"strconv"
	"time"
)

type JwtClaims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// CreateJwToken generate jwt token
func CreateJwToken(user models.User) (string, error) {

	jwtExpired, _ := strconv.ParseInt(os.Getenv("JWT_EXPIRED_MINUTES"), 10, 64)

	claims := JwtClaims{
		Name:  user.Name,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			Id:        strconv.Itoa(user.ID),
			ExpiresAt: time.Now().Add(time.Duration(jwtExpired) * time.Minute).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rawToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return token, nil
}
