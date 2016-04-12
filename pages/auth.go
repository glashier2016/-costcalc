package pages

import (
	"net/http"

	"github.com/labstack/echo"
)

func Auth(c echo.Context) error {
	return c.Render(http.StatusOK, "auth.html", map[string]interface{}{"test": "string"})
}
