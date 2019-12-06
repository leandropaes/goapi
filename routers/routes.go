package routers

import (
	"github.com/labstack/echo"
	controllers "github.com/leandropaes/goapi/app/controllers"
	"net/http"
)

// App é uma instancia de Echo
var App *echo.Echo

func init() {
	App = echo.New()

	// Página inicial da aplicação
	App.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "GO API!!",
		})
	})

	// Grupo de rotas
	api := App.Group("/v1")

	// Users
	api.GET("/users", controllers.UserIndex)
	api.GET("/users/:id", controllers.UserShow)
	api.POST("/users", controllers.UserCreate)
	api.PUT("/users/:id", controllers.UserUpdate)
	api.DELETE("/users/:id", controllers.UserDelete)

	// Sellers
	api.GET("/sellers", controllers.SellerIndex)
	api.GET("/sellers/:id", controllers.SellerShow)
	api.POST("/sellers", controllers.SellerCreate)
	api.PUT("/sellers/:id", controllers.SellerUpdate)
	api.DELETE("/sellers/:id", controllers.SellerDelete)
}
