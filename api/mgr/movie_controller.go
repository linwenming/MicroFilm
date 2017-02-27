package mgr

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"fmt"
	"os"
	"io"
	"MicroFilm/conf"
	"MicroFilm/model"
	"github.com/gocraft/dbr"
	"github.com/Sirupsen/logrus"
	"strconv"
)

func UploadMovieFile() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		//videoname := c.FormValue("name")

		// Source
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create(conf.MOVIES_DIR + file.Filename)
		if err != nil {
			if os.IsNotExist(err) {
				os.MkdirAll(conf.MOVIES_DIR, os.ModeDir)
				dst, err = os.Create(conf.MOVIES_DIR + file.Filename)
			} else {
				return err
			}
		}

		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		return c.JSON(fasthttp.StatusOK, map[string]interface{}{
			"code":0,
			"msg":"successful.",
			"data":map[string]interface{}{"filename":conf.SERVER_URL + file.Filename},
		})
	}
}

func AddMovie() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		m := model.NewMovieForm()
		c.Bind(m)

		tx := c.Get("Tx").(*dbr.Tx)

		if err := m.Save(tx); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError)
		}

		return c.JSON(fasthttp.StatusOK, map[string]interface{}{
			"code":0,
			"msg":"successful.",
		})
	}
}

func GetMovieById() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		tx := c.Get("Tx").(*dbr.Tx)
		m := model.NewMovieForm()

		if err := m.Load(tx, id); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError, err.Error())
		}

		return c.JSON(fasthttp.StatusOK, m)
	}
}

func EditBaseInfoOfMovie() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		m := model.NewMovieForm()
		c.Bind(m)

		logrus.Debug(m)

		tx := c.Get("Tx").(*dbr.Tx)

		if err := m.Update(tx); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError)
		}

		return c.JSON(fasthttp.StatusOK, map[string]interface{}{
			"code":0,
			"msg":"successful.",
		})
	}
}

// 特殊属性
func EditSpecialOfMovie() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id,_ := strconv.ParseInt(c.FormValue("id"),10,64);
		score,_ := strconv.ParseInt(c.FormValue("score"),10,64);
		playCount,_ := strconv.ParseInt(c.FormValue("playCount"),10,64);
		replyCount,_ := strconv.ParseInt(c.FormValue("replyCount"),10,64);
		zanCount,_ := strconv.ParseInt(c.FormValue("zanCount"),10,64);

		if id <= 0 {
			return c.JSON(fasthttp.StatusOK, map[string]interface{}{
				"code":1,
				"msg":"id 值不正确",
			})
		}

		m := new(model.Movie)
		m.Id = id;

		tx := c.Get("Tx").(*dbr.Tx)

		value := map[string]interface{}{
			"score":score,
			"play_count":playCount,
			"reply_count":replyCount,
			"zan_count":zanCount,
		}

		if err := m.UpdateBy(tx, value); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError)
		}

		return c.JSON(fasthttp.StatusOK, map[string]interface{}{
			"code":0,
			"msg":"successful.",
		})
	}
}

func UpOrDownMovie() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		section := c.QueryParam("section");
		fmt.Printf("section:%s", section)

		return c.JSON(fasthttp.StatusOK, interface{}("test"))
	}
}

func GetMoviesBy() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		section := c.QueryParam("section");
		fmt.Printf("section:%s", section)

		return c.JSON(fasthttp.StatusOK, interface{}("test"))
	}
}
