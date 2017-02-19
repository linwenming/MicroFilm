package route

import (
	"MicroFilm/api"
	"MicroFilm/db"
	currMw "MicroFilm/middleware"
	echoMw "github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
	"MicroFilm/handler"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"fmt"
)

func Init() *echo.Echo {

	e := echo.New()

	e.Debug = true

	// Set Bundle MiddleWare
	e.Use(echoMw.Logger())
	e.Use(echoMw.Gzip())
	e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))
	//e.Use(echoMw.JWT([]byte("secret")))

	e.HTTPErrorHandler  = handler.JSONHTTPErrorHandler
	//e.SetHTTPErrorHandler(handler.JSONHTTPErrorHandler)

	// Set Custom MiddleWare
	e.Use(currMw.TransactionHandler(db.Init()))

	e.POST("/login", api.Login())
	e.POST("/register", api.Register())

	// Routes
	v1 := e.Group("/api")
	{
		v1.GET("/", restricted)
		v1.GET("/user/:id", api.GetUser())
		v1.GET("/users/:active", api.GetUsers())
	}
	return e
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	fmt.Println("user name:" + name)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
