package mgr

import (
	"github.com/labstack/echo"
	"github.com/gocraft/dbr"
	"github.com/Sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"MicroFilm/model"
	"time"
	"fmt"
	"strconv"
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

// 支付回调
func Order_paymentCallback() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		orderSn := c.FormValue("orderSn")
		m := &model.OrderDetail{}

		if(c.FormValue("status") != nil) {
			m.Status == 2
		} else {
			m.Status == 3;
		}

		tx := c.Get("Tx").(*dbr.Tx)

		if err := m.LoadBySn(tx, orderSn); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError, "回调订单号不存在")
		}
		// 0待生成 1生成订单（待支付） 2:支付成功 3支付失败
		values := map[string]interface{}{
			"status":m.Status,
		}
		if err := m.UpdateBy(tx, values); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError, "订单号状态修改失败")
		}

		orderLog := &model.OrderLog{}
		orderLog.OrderSn = m.OrderSn
		orderLog.CreateTime = time.Now().Unix()
		orderLog.Content = fmt.Sprintf("订单状态回调：%d", m.Status)
		orderLog.Save(tx);

		return c.JSON(fasthttp.StatusOK, map[string]interface{}{
			"code":0,
			"msg":"successful.",
		})
	}
}

func Order_getBySn() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		orderSn := c.QueryParam("orderSn")
		m := &model.OrderDetail{}

		tx := c.Get("Tx").(*dbr.Tx)

		if err := m.LoadBySn(tx, orderSn); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError, "回调订单号不存在")
		}

		return c.JSON(fasthttp.StatusOK, map[string]interface{}{
			"code":0,
			"msg":"successful.",
			"data":m,
		})
	}
}


func Order_list() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		pageSize,NotFound := strconv.ParseInt(c.QueryParam("pageSize"),10,64)
		pageNumber,_ := strconv.ParseInt(c.QueryParam("pageNumber"),10,64)
		if NotFound != nil {
			pageSize = 10
		}

		orders := new(model.OrderDetails)
		tx := c.Get("Tx").(*dbr.Tx)

		if err := orders.Load(tx, pageSize,pageNumber); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError, err.Error())
		}

		return c.JSON(fasthttp.StatusOK, map[string]interface{}{
			"code":0,
			"msg":"successful.",
			"data": orders,
		})
	}
}


