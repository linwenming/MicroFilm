package api

import (
	"strconv"

	"github.com/Sirupsen/logrus"
	"MicroFilm/model"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

func Login() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		user := new(model.User)
		c.Bind(&user)

		tx := c.Get("Tx").(*dbr.Tx)

		if err := user.Save(tx); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError)
		}
		return c.JSON(fasthttp.StatusCreated, user)
	}
}

func GetUser() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		id, _ := strconv.ParseInt(c.Param("id"), 0, 64)

		tx := c.Get("Tx").(*dbr.Tx)

		user := new(model.User)
		if err := user.Load(tx, id); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusNotFound, "user does not exists.")
		}
		return c.JSON(fasthttp.StatusOK, user)
	}
}

func GetUsers() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		tx := c.Get("Tx").(*dbr.Tx)

		active,_ := strconv.Atoi(c.QueryParam("active"))
		users := new(model.Users)
		if err = users.Load(tx, active); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusNotFound, "user does not exists.")
		}

		return c.JSON(fasthttp.StatusOK, users)
	}
}
