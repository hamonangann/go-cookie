package main

import (
	"net/http"

	cookie "go-cookie/cookie"

	"github.com/gorilla/securecookie"
	"github.com/labstack/echo"
)

func main() {
	const CookieName = "data"

	e := echo.New()

	e.GET("/index", func(c echo.Context) error {
		data, err := cookie.GetCookie(c, CookieName)
		if err != nil && err != http.ErrNoCookie && err != securecookie.ErrMacInvalid {
			return err
		}

		if data == nil {
			data = cookie.M{"message": "Hello", "ID": "asdfghjklzxcvbnm"}

			err = cookie.SetCookie(c, CookieName, data)
			if err != nil {
				return err
			}
		}

		return c.JSON(http.StatusOK, data)
	})

	e.Logger.Fatal(e.Start(":9000"))
}
