package app

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"fmt"
	"strconv"
	"github.com/Sirupsen/logrus"
	"github.com/gocraft/dbr"
	"MicroFilm/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func Comment_list() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		mid, _ := strconv.ParseInt(c.QueryParam("mid"), 10, 64)

		claims := c.Get("User").(jwt.MapClaims)
		uid := int64(claims["uid"].(float64))

		logrus.Debug("根据电影编号查询评论列表: ", mid ,"  uid:",uid)

		tx := c.Get("Tx").(*dbr.Tx)

		var comments model.Comments
		if err := comments.Load(tx, uid, mid); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError, err.Error())
		}

		return c.JSON(fasthttp.StatusOK, map[string]interface{}{
			"code":0,
			"msg":"successful.",
			"data":comments,
		})
	}
}

func Comment_reply() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		comment := &model.Comment{
			CreateTime:time.Now().Unix(),
		}
		c.Bind(comment)

		claims := c.Get("User").(jwt.MapClaims)
		uid := claims["uid"].(float64)
		comment.Uid = int64(uid)

		tx := c.Get("Tx").(*dbr.Tx)

		if err := comment.Save(tx); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError, err.Error())
		}

		return c.JSON(fasthttp.StatusOK, map[string]interface{}{
			"code":0,
			"msg":"successful.",
		})
	}
}

func Comment_zan() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id := c.QueryParam("mid");
		fmt.Printf("id:%d", id)

		return c.JSON(fasthttp.StatusOK, interface{}("test"))
	}
}