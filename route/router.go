package route

import (
	"MicroFilm/api"
	"MicroFilm/db"
	currMw "MicroFilm/middleware"
	echoMw "github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
	"MicroFilm/handler"
	"MicroFilm/api/mgr"
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

	e.HTTPErrorHandler  = handler.JSONHTTPErrorHandler
	//e.SetHTTPErrorHandler(handler.JSONHTTPErrorHandler)

	// Set Custom MiddleWare
	e.Use(currMw.TransactionHandler(db.Init()))

	e.POST("/login", api.Login())
	e.POST("/register", api.Register())

	// Routes
	v1 := e.Group("/api",currMw.AuthroizationHandler())
	{
		v1.GET("/user/:id", api.GetUser())
		v1.GET("/users/:active", api.GetUsers())

		v1.POST("/mgr/uploadfile", mgr.UploadMovieFile())
	}
	v1.Use(echoMw.JWT([]byte("secret")))
	//v1.Use(echoMw.JWTWithConfig(echoMw.JWTConfig{
	//	Skipper:       func(echo.Context) bool { return false },
	//	SigningMethod: echoMw.AlgorithmHS256,
	//	ContextKey:    "user",
	//	TokenLookup:   "header:" + echo.HeaderAuthorization,
	//	SigningKey:    []byte("secret"),
	//	AuthScheme:    "Bearer",
	//}))
	return e
}


