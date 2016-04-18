package pages

import (
	"net/http"

	"github.com/labstack/echo"
)

func Reg(c echo.Context) error {
	return c.Render(http.StatusOK, "reg.html", map[string]interface{}{"test": "string"})
}
