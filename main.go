package main

import (
	"fmt"
	"log"
	"net/http"

	cookie "go-cookie/cookie"

	"github.com/gorilla/securecookie"
	"github.com/labstack/echo"
)

const SESSION_ID = "id"

var store = cookie.NewCookieStore()

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

	e.GET("/session/set", func(c echo.Context) error {
		session, _ := store.Get(c.Request(), SESSION_ID)
		log.Printf(session.ID)
		session.Values["message1"] = "hello"
		session.Values["message2"] = "world"
		session.Save(c.Request(), c.Response())

		return c.Redirect(http.StatusTemporaryRedirect, "/session/get")
	})

	e.GET("/session/get", func(c echo.Context) error {
		session, _ := store.Get(c.Request(), SESSION_ID)

		if len(session.Values) == 0 {
			return c.String(http.StatusOK, "empty result")
		}

		return c.String(http.StatusOK, fmt.Sprintf(
			"%s %s",
			session.Values["message1"],
			session.Values["message2"],
		))
	})

	e.GET("/session/delete", func(c echo.Context) error {
		session, _ := store.Get(c.Request(), SESSION_ID)
		session.Options.MaxAge = -1
		session.Save(c.Request(), c.Response())

		return c.Redirect(http.StatusTemporaryRedirect, "/session/get")
	})

	e.Logger.Fatal(e.Start(":9000"))
}
