package app

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"fmt"
)

func GetCommentList() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		mid := c.QueryParam("mid");
		fmt.Printf("mid:%s", mid)

		return c.JSON(fasthttp.StatusOK, interface{}("test"))
	}
}

func Reply() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		mid := c.QueryParam("mid");
		fmt.Printf("mid:%s", mid)

		return c.JSON(fasthttp.StatusOK,interface{}("test"))
	}
}

func ZanComment() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id := c.QueryParam("mid");
		fmt.Printf("id:%d", id)

		return c.JSON(fasthttp.StatusOK, interface{}("test"))
	}
}