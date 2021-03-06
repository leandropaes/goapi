package routers

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	controllers "github.com/leandropaes/goapi/app/controllers"
	"net/http"
	"os"
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

	// Auth
	App.POST("/auth/login", controllers.AuthLogin)

	// Grupo de rotas
	api := App.Group("/v1")

	// Middleware JWT para grupo de rota "api"
	api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(os.Getenv("JWT_SECRET")),
		SigningMethod: "HS512",
	}))

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
