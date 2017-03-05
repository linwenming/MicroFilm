package route

import (
	"MicroFilm/db"
	currMw "MicroFilm/middleware"
	echoMw "github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
	"MicroFilm/handler"
	"MicroFilm/api/mgr"
	"MicroFilm/api/app"
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

	e.HTTPErrorHandler = handler.JSONHTTPErrorHandler
	//e.SetHTTPErrorHandler(handler.JSONHTTPErrorHandler)

	// Set customer MiddleWare
	session := db.Init();
	e.Use(currMw.TransactionHandler(session))
	e.Use(currMw.CacheHandler(session))

	e.POST("/login", safe.Login())
	e.POST("/register", safe.Register())

	// Routes
	_app := e.Group("/app", currMw.AuthroizationHandler())
	{
		_app.GET("/movie/list", app.Movie_listByCate())
		_app.GET("/movie/:mid", app.Movie_listByCate())
		_app.GET("/movie/zan", app.Movie_zan())

		_app.GET("/comment/list", app.Comment_list())
		_app.POST("/comment/reply", app.Comment_reply())
		_app.GET("/comment/zan", app.Comment_zan())
	}
	_app.Use(echoMw.JWT([]byte("secret")))

	_mgr := e.Group("/mgr", currMw.AuthroizationHandler())
	{
		_mgr.POST("/movie/uploadvideo", mgr.Movie_upload())
		_mgr.POST("/movie", mgr.Movie_add())
		_mgr.DELETE("/movie/:id", mgr.Movie_del())
		_mgr.PUT("/movie", mgr.Movie_editBaseInfo())
		_mgr.PATCH("/movie/statprop", mgr.Movie_editStatProperty())
		_mgr.GET("/movie/:id", mgr.Movie_loadById())
		_mgr.GET("/movie/status", mgr.Movie_updateStatus())

		_mgr.POST("/cate", mgr.Cate_add())
		_mgr.DELETE("/cate/:id", mgr.Cate_del())
		_mgr.PUT("/cate", mgr.Cate_edit())
		_mgr.GET("/cate/:id", mgr.Cate_loadById())
		_mgr.GET("/cate/list", mgr.Cate_list())

		_mgr.POST("/order", mgr.Order_create())
		_mgr.GET("/order/callback", mgr.Order_paymentCallback())
		_mgr.GET("/order/sn/:orderSn", mgr.Order_getBySn())
		_mgr.GET("/order/list", mgr.Order_list())
		//manage.GET("/order/:id", mgr.Cate_loadById())
		//manage.GET("/order/list", mgr.Cate_list())

	}
	_mgr.Use(echoMw.JWT([]byte("secret")))

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


