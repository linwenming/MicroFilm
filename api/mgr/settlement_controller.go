package mgr

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"fmt"
	"strconv"
	"github.com/gocraft/dbr"
	"github.com/Sirupsen/logrus"
	"MicroFilm/model"
)

func Stmt_listByPage() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		pageSize,NotFound := strconv.ParseUint(c.QueryParam("pageSize"),10,64)
		pageNumber,_ := strconv.ParseUint(c.QueryParam("pageNumber"),10,64)
		if NotFound != nil {
			pageSize = 10
		}

		data := new(model.Settlements)
		tx := c.Get("Tx").(*dbr.Tx)

		if err := data.LoadByPage(tx, pageSize,pageNumber); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError, err.Error())
		}

		return c.JSON(fasthttp.StatusOK, map[string]interface{}{
			"code":0,
			"msg":"successful.",
			"data": data,
		})
	}
}

func Stmt_listByDate() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		stime,_ := strconv.ParseInt(c.QueryParam("startTime"),10,64);
		etime,_ := strconv.ParseInt(c.QueryParam("endTime"),10,64);

		fmt.Printf("stime:%d  etime:%d", stime,etime)

		data := new(model.Settlements)
		tx := c.Get("Tx").(*dbr.Tx)

		if err := data.LoadByTime(tx, stime,etime); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError, err.Error())
		}

		return c.JSON(fasthttp.StatusOK, map[string]interface{}{
			"code":0,
			"msg":"successful.",
			"data": data,
		})
	}
}

func Stmt_getSettlement() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id,_ := strconv.ParseInt(c.QueryParam("id"),10,64);
		fmt.Printf("id:%s", id)

		data := new(model.Settlement)
		tx := c.Get("Tx").(*dbr.Tx)

		if err := data.Load(tx, id); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError, err.Error())
		}

		return c.JSON(fasthttp.StatusOK, map[string]interface{}{
			"code":0,
			"msg":"successful.",
			"data": data,
		})
	}
}

func Stmt_closed() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id,_ := strconv.ParseInt(c.QueryParam("id"),10,64);
		status,_ := strconv.ParseInt(c.QueryParam("status"),10,64);
		fmt.Printf("id:%d status:%d", id,status)

		stmt := new(model.Settlement)
		tx := c.Get("Tx").(*dbr.Tx)

		if err := stmt.Load(tx, id); err != nil {
			logrus.Debug(err)
			return c.JSON(fasthttp.StatusOK, map[string]interface{}{
				"code":1,
				"msg":"记录不存在.",
			})
		}

		stmt.Status = status;
		stmt.Update(tx);

		return c.JSON(fasthttp.StatusOK, map[string]interface{}{
			"code":0,
			"msg":"successful.",
		})
	}
}

