package controllers

import (
	"github.com/labstack/echo"
	"github.com/leandropaes/goapi/app/models"
	"net/http"
	"strconv"
)

// SellerIndex list all sellers
func SellerIndex(c echo.Context) error {

	var seller models.Sellers

	if err := models.SellerModel.Find().All(&seller); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Nenhum registro encontrado",
		})
	}

	return c.JSON(http.StatusOK, map[string]models.Sellers{
		"data": seller,
	})
}

// SellerShow show seller by id
func SellerShow(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var seller models.Seller

	result := models.SellerModel.Find("id=?", id)

	if count, _ := result.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Não foi possível encontrar o registro",
		})
	}

	if err := result.One(&seller); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Erro ao tentar carregar o registro",
		})
	}

	return c.JSON(http.StatusOK, map[string]models.Seller{
		"data": seller,
	})
}

// SellerCreate create seller
func SellerCreate(c echo.Context) error {
	name := c.FormValue("name")

	var seller models.Seller
	seller.Name = name

	// TODO: Verificar uma mellhor forma de fazer validação dos campos
	if name != "" {
		if _, err := models.SellerModel.Insert(seller); err != nil {
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

// SellerUpdate update seller
func SellerUpdate(c echo.Context) error {
	// TODO: Fazer validação dos campos
	id, _ := strconv.Atoi(c.Param("id"))
	name := c.FormValue("name")

	// TODO: Verificar alguma forma de não ter que sempre passar todos os campos para atualizar
	var seller = models.Seller{
		ID:    id,
		Name:  name,
	}

	result := models.SellerModel.Find("id=?", id)

	if count, _ := result.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Não foi possível encontrar o registro",
		})
	}

	if err := result.Update(seller); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Não foi possível atualizar o registro",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Registro atualizado com sucesso",
	})
}

// SellerDelete delete seller
func SellerDelete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	result := models.SellerModel.Find("id=?", id)

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