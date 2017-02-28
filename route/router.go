package route

import (
	"MicroFilm/db"
	currMw "MicroFilm/middleware"
	echoMw "github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
	"MicroFilm/handler"
	"MicroFilm/api/mgr"
	"MicroFilm/api/safe"
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

	e.POST("/login", safe.Login())
	e.POST("/register", safe.Register())

	// Routes
	app := e.Group("/app",currMw.AuthroizationHandler())
	{
		//api.GET("/user/:id", sa.GetUser())
	}
	app.Use(echoMw.JWT([]byte("secret")))

	manage := e.Group("/mgr",currMw.AuthroizationHandler())
	{
		manage.POST("/movie/uploadvideo", mgr.Movie_upload())
		manage.POST("/movie", mgr.Movie_add())
		manage.DELETE("/movie/:id", mgr.Movie_del())
		manage.PUT("/movie", mgr.Movie_editBaseInfo())
		manage.PATCH("/movie/statprop", mgr.Movie_editStatProperty())
		manage.GET("/movie/:id", mgr.Movie_loadById())
		manage.GET("/movie/status", mgr.Movie_updateStatus())

		manage.POST("/cate", mgr.Cate_add())
		manage.DELETE("/cate/:id", mgr.Cate_del())
		manage.PUT("/cate", mgr.Cate_edit())
		manage.GET("/cate/:id", mgr.Cate_loadById())
		manage.GET("/cate/list", mgr.Cate_list())

		manage.POST("/order", mgr.Order_create())
		//manage.DELETE("/order/:id", mgr.Cate_del())
		//manage.PUT("/order", mgr.Cate_edit())
		//manage.GET("/order/:id", mgr.Cate_loadById())
		//manage.GET("/order/list", mgr.Cate_list())

	}
	manage.Use(echoMw.JWT([]byte("secret")))

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


