package mgr

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
