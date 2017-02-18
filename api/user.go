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

		username := c.FormValue("username")
		password := c.FormValue("password")
		fmt.Printf("username:%s  password:%s",username,password + "\n")

		if len(username) == 0 || len(password) == 0 {
			return c.JSON(fasthttp.StatusOK, map[string]interface{}{
				"code":101,
				"msg":"用户名或密码不能为空.",
			})
		}

		tx := c.Get("Tx").(*dbr.Tx)

		user := new(model.User)

		if err := user.LoadByUsername(tx, username); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError,"服务器内部错误")
		}

		if !strings.EqualFold(user.Password, password) {
			return c.JSON(fasthttp.StatusOK, map[string]interface{}{
				"code":103,
				"msg":"密码错误.",
			})
		}

		return c.JSON(fasthttp.StatusOK, map[string]interface{}{
			"code":0,
			"msg":"登录成功.",
		})
	}
}

func Register() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		m := model.NewUser()
		c.Bind(&m)

		tx := c.Get("Tx").(*dbr.Tx)

		user := new(model.User)
		user.LoadByUsername(tx,m.Username)

		fmt.Printf("username:%s  db username:%s",m.Username,user.Username + "\n")

		if strings.EqualFold(m.Username,user.Username) {
			return c.JSON(fasthttp.StatusOK, map[string]interface{}{
				"code":104,
				"msg":"用户名已经存在.",
			})
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
