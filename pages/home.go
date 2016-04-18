package pages

import (
	"net/http"

	"github.com/labstack/echo"
)

func Home(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{"test": "string"})
}
