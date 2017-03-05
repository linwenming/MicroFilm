package mgr

import (
	"github.com/labstack/echo"
	"github.com/gocraft/dbr"
	"github.com/Sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"MicroFilm/model"
	"strconv"
	"MicroFilm/middleware"
)

func Cate_add() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		m := &model.Category{}
		c.Bind(m)

		tx := c.Get("Tx").(*dbr.Tx)

		if err := m.Save(tx); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError,err.Error())
		}

		return c.JSON(fasthttp.StatusOK, map[string]interface{}{
			"code":0,
			"msg":"successful.",
		})
	}
}

func Cate_del() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		id,_ := strconv.ParseInt(c.Param("id"),10,64);
		m := &model.Category{}
		tx := c.Get("Tx").(*dbr.Tx)

		if err := m.Delete(tx,id); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError,err.Error())
		}

		return c.JSON(fasthttp.StatusOK, map[string]interface{}{
			"code":0,
			"msg":"successful.",
		})
	}
}

// 修改基本属性
func Cate_edit() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		m := &model.Category{}
		c.Bind(m)

		logrus.Debug(m)

		tx := c.Get("Tx").(*dbr.Tx)

		if err := m.Update(tx); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError,err.Error())
		}

		return c.JSON(fasthttp.StatusOK, map[string]interface{}{
			"code":0,
			"msg":"successful.",
		})
	}
}

func Cate_loadById() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		tx := c.Get("Tx").(*dbr.Tx)
		m := &model.Category{}

		if err := m.Load(tx, id); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError, err.Error())
		}

		return c.JSON(fasthttp.StatusOK, m)
	}
}

func Cate_list() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		//tx := c.Get("Tx").(*dbr.Tx)
		//
		//var list model.CategoryList
		//
		//if err := list.Load(tx); err != nil {
		//	logrus.Debug(err)
		//	return echo.NewHTTPError(fasthttp.StatusInternalServerError, err.Error())
		//}
		list := middleware.CacheCateList(c).(model.CategoryList)

		return c.JSON(fasthttp.StatusOK, list)
	}
}