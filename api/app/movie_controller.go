package app

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"fmt"
	"strconv"
	"github.com/gocraft/dbr"
	"github.com/Sirupsen/logrus"
	"MicroFilm/model"
	"time"
	"github.com/dgrijalva/jwt-go"
)

func Movie_listByCate() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		cateId, _ := strconv.ParseInt(c.QueryParam("cid"), 10, 64)
		logrus.Debug("根据分类查询电影: ", cateId)

		tx := c.Get("Tx").(*dbr.Tx)

		movies :=  new(model.Movies)
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

func Movie_listByFuzzy() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		keyword := c.QueryParam("keyword")
		logrus.Debug("根据关键字查询电影: ", keyword)

		tx := c.Get("Tx").(*dbr.Tx)

		var movies model.Movies
		if err := movies.LoadByFuzzy(tx, keyword); err != nil {
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

func Movie_zan() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		mid,_ := strconv.ParseInt(c.QueryParam("mid"),10,64)

		claims := c.Get("User").(jwt.MapClaims)
		uid := int64(claims["uid"].(float64))

		fmt.Printf("mid:%d  uid:%d", mid,uid)

		tx := c.Get("Tx").(*dbr.Tx)

		var zan model.Zan
		if err := zan.LoadBy(tx, uid,mid); err != nil {
			logrus.Debug(err)
			return c.JSON(fasthttp.StatusOK, map[string]interface{}{
				"code":1,
				"msg":"你已经点过赞.",
			})
		}
		zan.Uid = uid;
		zan.Mid = mid;
		zan.CreateTime = time.Now().Unix()
		zan.Save(tx)

		var movie model.Movie
		if err := movie.Load(tx, mid); err != nil {
			logrus.Debug(err)
			return c.JSON(fasthttp.StatusOK, map[string]interface{}{
				"code":2,
				"msg":"movie is not found.",
			})
		}

		movie.UpdateBy(tx, map[string]interface{}{"zanCount":movie.ZanCount + 1,})

		return c.JSON(fasthttp.StatusOK, map[string]interface{}{
			"code":0,
			"msg":"successfully.",
		})
	}
}

func Movie_Authorized() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		mid := c.FormValue("mid");
		fmt.Printf("mid:%d", mid)

		return c.JSON(fasthttp.StatusOK, interface{}("test"))
	}
}