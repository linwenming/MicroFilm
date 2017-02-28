package mgr

import (
	"github.com/labstack/echo"
	"github.com/gocraft/dbr"
	"github.com/Sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"MicroFilm/model"
)

func Order_create() echo.HandlerFunc {
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