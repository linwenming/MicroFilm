package middleware

import (
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
	"github.com/Sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"fmt"
)

//
// https://github.com/kyokomi/echo-jwt-sample/blob/master/hs256.go
// https://github.com/go-demo/jwtsample/blob/master/main.go

func AuthroizationHandler() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			fmt.Println("=========authroized start ===========")
			tokenString := c.Request().Header.Get("Bearer")
			if tokenString == "" {
				return echo.NewHTTPError(fasthttp.StatusUnauthorized,"token is null")
			}
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

				//if token.Claims["user"].(string) != "admin" {
				//	return errors.New("User isvalid"),nil
				//}
				//return []byte("secret"), nil
				return []byte("secret"), nil
			})
			if err != nil  || !token.Valid{
				logrus.Debug(err)
				return echo.NewHTTPError(fasthttp.StatusUnauthorized,"token isvalid")
				//return errors.New("token isvalid")
			}

			//claims := token.Claims.(jwt.MapClaims)
			//name := claims["name"].(string)
			//admin := claims["admin"].(bool)
			//exp  := claims["exp"].(float64)
			//fmt.Printf("Welcome %s admin:%t  exp:%d \n",name,admin,exp)

			fmt.Println("=========authroized end   ===========")
			//c.Set("User", token.Claims["user"])
			//uid := claims["uid"].(float64)
			//name := claims["name"].(string)
			//admin := claims["admin"].(bool)
			//exp  := claims["exp"].(float64)
			//fmt.Printf("Welcome %s uid:%d admin:%t  exp:%d \n",uid,name,admin,exp)
			claims := token.Claims.(jwt.MapClaims)
			c.Set("User", claims)
			return next(c)
		})
	}
}

//func restricted(c echo.Context) error {
//	user := c.Get("user").(*jwt.Token)
//	claims := user.Claims.(jwt.MapClaims)
//	name := claims["name"].(string)
//	fmt.Println("user name:" + name)
//	return c.String(http.StatusOK, "Welcome "+name+"!")
//}