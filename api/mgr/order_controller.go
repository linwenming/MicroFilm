package mgr

import (
	"github.com/labstack/echo"
	"github.com/gocraft/dbr"
	"github.com/Sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"MicroFilm/model"
	"time"
	"fmt"
)

func Order_create() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		m := &model.OrderDetail{}
		c.Bind(m)

		tx := c.Get("Tx").(*dbr.Tx)

		product := &model.Product{}
		if err := product.Load(tx, m.ProductId); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError, "该商品不存在")
		}
		// init attr
		m.Status = 0
		m.CreateTime = time.Now().Unix()
		m.ProductName = product.ProductName
		m.ProductPrice = product.Price
		// Quantity must gt 0
		if m.Quantity <= 0 {
			m.Quantity = 1
		}
		m.TotalPrice = product.Price * m.Quantity

		if err := m.Save(tx); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError, err.Error())
		}

		// out log
		orderLog := &model.OrderLog{}
		orderLog.OrderSn = m.OrderSn
		orderLog.CreateTime = time.Now().Unix()
		orderLog.Content = fmt.Sprintf("创建订单：创建者id为%d,消费总价为%d", m.Uid, m.TotalPrice)
		orderLog.Save(tx);

		return c.JSON(fasthttp.StatusOK, map[string]interface{}{
			"code":0,
			"msg":"successful.",
			"data":&map[string]interface{}{
				"product":m.ProductName,
				"orderSn":m.OrderSn,
				"totalPrice":m.TotalPrice,
			},
		})
	}
}
