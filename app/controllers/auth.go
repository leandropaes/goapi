package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/leandropaes/goapi/app/models"
	"net/http"
	"strconv"
	"time"
)

type JwtClaims struct {
	Name string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func createJwToken(user models.User) (string, error) {
	claims := JwtClaims{
		Name:           user.Name,
		Email:          user.Email,
		StandardClaims: jwt.StandardClaims{
			Id:        strconv.Itoa(user.ID),
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rawToken.SignedString([]byte("mySecret"))

	if err != nil {
		return "", err
	}

	return token, nil
}

func AuthLogin(c echo.Context) error {
	login := c.FormValue("login")
	password := c.FormValue("password")

	var user models.User

	result := models.UserModel.Find("email=? AND password=?", login, password)

	if count, _ := result.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Não foi possível encontrar o registro",
		})
	}

	if err := result.One(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Erro ao tentar carregar o registro",
		})
	}

	token, err := createJwToken(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Erro ao tentar gerar o token",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": user,
		"token": token,
	})
}
