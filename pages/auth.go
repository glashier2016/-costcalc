package pages

import (
	"net/http"

	"github.com/labstack/echo"
)

func Auth(c echo.Context) error {
	return c.String(http.StatusOK, "/public/auth.html")
}
