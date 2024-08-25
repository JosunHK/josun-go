package dummy

import (
	dummyTemplates "github.com/JosunHK/josun-go.git/web/templates/contents/dummy"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Dummy(c echo.Context) templ.Component {
	return dummyTemplates.Dummy()
}
