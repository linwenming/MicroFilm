package app

import (
	"github.com/labstack/echo"
	//"github.com/gocraft/dbr"
	//"github.com/Sirupsen/logrus"
	"github.com/valyala/fasthttp"
	//"strconv"
	"fmt"
	"strconv"
	"github.com/gocraft/dbr"
	"github.com/Sirupsen/logrus"
	"MicroFilm/model"
)

func Movie_listByCate() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		cateId, _ := strconv.ParseInt(c.QueryParam("cateId"), 10, 64)
		logrus.Debug("根据分类查询电影: ", cateId)

		tx := c.Get("Tx").(*dbr.Tx)

		var movies model.Movies
		if err := movies.LoadByCate(tx, cateId); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError, err.Error())
		}

		return c.JSON(fasthttp.StatusOK, map[string]interface{}{
			"code":0,
			"msg":"successful.",
			"data":movies,
		})
	}
}

func Movie_getDetail() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		mid, _ := strconv.ParseInt(c.QueryParam("mid"), 10, 64)
		logrus.Debug("查询电影详情: ", mid)

		tx := c.Get("Tx").(*dbr.Tx)

		var movie model.Movie
		if err := movie.Load(tx, mid); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError, err.Error())
		}

		return c.JSON(fasthttp.StatusOK, map[string]interface{}{
			"code":0,
			"msg":"successful.",
			"data":movie,
		})
	}
}

// 点赞
func Movie_zan() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id := c.QueryParam("mid");
		fmt.Printf("id:%d", id)

		return c.JSON(fasthttp.StatusOK, interface{}("test"))
	}
}

func PlayAuthorized() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		uid := c.QueryParam("uid");
		mid := c.QueryParam("mid");
		fmt.Printf("uid:%d  mid:%d", uid, mid)

		return c.JSON(fasthttp.StatusOK, interface{}("test"))
	}
}