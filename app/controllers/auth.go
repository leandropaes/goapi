package controllers

import (
	"github.com/labstack/echo"
	"github.com/leandropaes/goapi/app/models"
	"github.com/leandropaes/goapi/lib"
	"net/http"
)

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

	token, err := lib.CreateJwToken(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Erro ao tentar gerar o token",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  user,
		"token": token,
	})
}
