package pages

import (
	"net/http"

	"github.com/labstack/echo"
)

func auth(c echo.Context) error {
	return c.String(http.StatusOK, "/public/auth.html")
}
