package routes

import (
	"github.com/jean27garbi/echo-app1/handlers"
	"github.com/labstack/echo/v4"
)

func UserRouter() *echo.Echo {
	e := echo.New()

	// Group
	user := e.Group("/users")
	// user.Use(echoMiddleWare.CORS())

	// Routing
	user.GET("/", handlers.ListUsers)
	user.POST("/", handlers.SaveUser)
	user.POST("/avatar/", handlers.SaveUserImage)
	user.GET("/:id", handlers.GetUser)
	user.PUT("/:id", handlers.UpdateUser)
	user.DELETE("/:id", handlers.DeleteUser)

	// Static
	user.Static("/static", "static")

	return e
}
