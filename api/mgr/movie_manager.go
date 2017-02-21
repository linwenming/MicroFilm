package mgr

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"fmt"
)

func UploadMovieFile() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		section := c.QueryParam("section");
		fmt.Printf("section:%s", section)

		return c.JSON(fasthttp.StatusOK, interface{}("test"))
	}
}

func AddMovie() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		section := c.QueryParam("section");
		fmt.Printf("section:%s", section)

		return c.JSON(fasthttp.StatusOK, interface{}("test"))
	}
}

func EditMovieOfBaseInfo() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		section := c.QueryParam("section");
		fmt.Printf("section:%s", section)

		return c.JSON(fasthttp.StatusOK, interface{}("test"))
	}
}

func GetUnactivedMovies() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		section := c.QueryParam("section");
		fmt.Printf("section:%s", section)

		return c.JSON(fasthttp.StatusOK, interface{}("test"))
	}
}
