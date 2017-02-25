package app

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"fmt"
)

// step 1: 提交订单生产支付信息页面
func CommitOrder() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id := c.QueryParam("mid");
		fmt.Printf("id:%d", id)

		return c.JSON(fasthttp.StatusOK, interface{}("test"))
	}
}

func Payment() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id := c.QueryParam("mid");
		fmt.Printf("id:%d", id)

		return c.JSON(fasthttp.StatusOK, interface{}("test"))
	}
}

func GetOrderStatus() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id := c.QueryParam("mid");
		fmt.Printf("id:%d", id)

		return c.JSON(fasthttp.StatusOK, interface{}("test"))
	}
}

func GetOrderDetail() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id := c.QueryParam("mid");
		fmt.Printf("id:%d", id)

		return c.JSON(fasthttp.StatusOK, interface{}("test"))
	}
}

// 第三方支付平台回调付款消息
func OnPaymentResult() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id := c.QueryParam("mid");
		fmt.Printf("id:%d", id)

		return c.JSON(fasthttp.StatusOK, interface{}("test"))
	}
}