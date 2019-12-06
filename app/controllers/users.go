package controllers

import (
	"github.com/labstack/echo"
	"github.com/leandropaes/goapi/app/models"
	"net/http"
	"strconv"
)

// UserIndex list all users
func UserIndex(c echo.Context) error {

	var user models.Users

	if err := models.UserModel.Find().All(&user); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Nenhum registro encontrado",
		})
	}

	return c.JSON(http.StatusOK, map[string]models.Users{
		"data": user,
	})
}

// UserShow show user by id
func UserShow(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var user models.User

	result := models.UserModel.Find("id=?", id)

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

	return c.JSON(http.StatusOK, map[string]models.User{
		"data": user,
	})
}

// UserCreate create user
func UserCreate(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	var user models.User
	user.Name = name
	user.Email = email
	user.Password = password

	// TODO: Verificar uma mellhor forma de fazer validação dos campos
	if name != "" && email != "" && password != "" {
		if _, err := models.UserModel.Insert(user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Não foi possível adicionar o registro no banco de dados.",
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "Registro cadastrado com sucesso.",
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"message": "Todos os campos são obrigatórios",
	})
}

// UserUpdate update user
func UserUpdate(c echo.Context) error {
	// TODO: Fazer validação dos campos
	id, _ := strconv.Atoi(c.Param("id"))
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	// TODO: Verificar alguma forma de não ter que sempre passar todos os campos para atualizar
	var user = models.User{
		ID:    id,
		Name:  name,
		Email: email,
		Password: password,
	}

	result := models.UserModel.Find("id=?", id)

	if count, _ := result.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Não foi possível encontrar o registro",
		})
	}

	if err := result.Update(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Não foi possível atualizar o registro",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Registro atualizado com sucesso",
	})
}

// UserDelete delete user
func UserDelete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	result := models.UserModel.Find("id=?", id)

	if count, _ := result.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Não foi possível encontrar o registro",
		})
	}

	if err := result.Delete(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Não foi possível excluir o registro",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Registro excluído com sucesso",
	})
}