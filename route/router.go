package route

import (
	"MicroFilm/api"
	"MicroFilm/db"
	currMw "MicroFilm/middleware"
	echoMw "github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
	"MicroFilm/handler"
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
	e.Use(echoMw.JWT([]byte("secret")))
	// 自定义
	//e.Use(echoMw.JWTWithConfig(echoMw.JWTConfig{
	//	SigningKey: []byte("secret"),
	//	TokenLookup: "query:token",
	//}))

	e.HTTPErrorHandler  = handler.JSONHTTPErrorHandler
	//e.SetHTTPErrorHandler(handler.JSONHTTPErrorHandler)

	// Set Custom MiddleWare
	e.Use(currMw.TransactionHandler(db.Init()))

	// Routes
	v1 := e.Group("/api")
	{
		v1.POST("/login", api.Login())
		v1.GET("/user/:id", api.GetUser())
		v1.GET("/users/:active", api.GetUsers())
	}
	return e
}
