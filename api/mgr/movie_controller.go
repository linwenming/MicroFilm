package mgr

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"fmt"
	"os"
	"io"
	"MicroFilm/conf"
	"echo-sample/model"
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

		m := model.NewMember()

		section := c.QueryParam("section");
		fmt.Printf("section:%s", section)

		return c.JSON(fasthttp.StatusOK, interface{}("test"))
	}
}

func EditMovieOfBaseInfo() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		section := c.QueryParam("section");
		fmt.Printf("section:%s", section)

		return c.JSON(fasthttp.StatusOK, interface{}("test"))
	}
}

func GetMovies() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		section := c.QueryParam("section");
		fmt.Printf("section:%s", section)

		return c.JSON(fasthttp.StatusOK, interface{}("test"))
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

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}