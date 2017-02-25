package app

import (
	"github.com/labstack/echo"
	//"github.com/gocraft/dbr"
	//"github.com/Sirupsen/logrus"
	"github.com/valyala/fasthttp"
	//"strconv"
	"fmt"
)

func GetMoviesBySection() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		section := c.QueryParam("section");
		fmt.Printf("section:%s", section)

		return c.JSON(fasthttp.StatusOK, interface{}("test"))
	}
}

func GetMoviesByCategory() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		category_id := c.QueryParam("category_id");
		fmt.Printf("category_id:%s", category_id)

		return c.JSON(fasthttp.StatusOK, interface{}("test"))
	}
}

func GetMoviesByTags() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		tags := c.QueryParam("tags");
		fmt.Printf("tags:%s", tags)

		return c.JSON(fasthttp.StatusOK, interface{}("test"))
	}
}

func GetMovieDetail() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id := c.QueryParam("mid");
		fmt.Printf("id:%d", id)

		return c.JSON(fasthttp.StatusOK, interface{}("test"))
	}
}

func ZanMovie() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id := c.QueryParam("mid");
		fmt.Printf("id:%d", id)

		return c.JSON(fasthttp.StatusOK, interface{}("test"))
	}
}

func PlayAuthorized() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		uid := c.QueryParam("uid");
		mid := c.QueryParam("mid");
		fmt.Printf("uid:%d  mid:%d", uid,mid)

		return c.JSON(fasthttp.StatusOK, interface{}("test"))
	}
}