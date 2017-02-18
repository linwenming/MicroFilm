package api

import (
	"strconv"

	"github.com/Sirupsen/logrus"
	"MicroFilm/model"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"strings"
	"fmt"
)

func Login() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		username := c.Param("username")
		password := c.QueryParam("password")

		fmt.Println("=========asdasdasd username:" + username)

		user := new(model.User)

		if len(username) == 0 || len(password) == 0 {
			return echo.NewHTTPError(101, "用户名或密码不能为空.")
		}

		tx := c.Get("Tx").(*dbr.Tx)

		if err := user.LoadBy(tx, []string{"username", username}); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError)
		}

		if !strings.EqualFold(user.Password, password) {
			return echo.NewHTTPError(102, "密码错误.")
		}

		return c.JSON(fasthttp.StatusOK, user)
	}
}

func Register() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		m := new(model.User)
		c.Bind(&m)

		tx := c.Get("Tx").(*dbr.Tx)

		user := model.NewUser();

		if err := user.LoadBy(tx, []string{"username", user.Username}); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError)
		}
		if strings.EqualFold(m.Username,user.Username) {
			return echo.NewHTTPError(103, "用户名已经存在.")
		}

		if err := m.Save(tx); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError)
		}
		return c.JSON(fasthttp.StatusOK, user)
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

		active, _ := strconv.Atoi(c.QueryParam("active"))
		users := new(model.Users)
		if err = users.Load(tx, active); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusNotFound, "user does not exists.")
		}

		return c.JSON(fasthttp.StatusOK, users)
	}
}
